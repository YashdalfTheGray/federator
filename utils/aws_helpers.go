package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// GetAWSConfig pulls the default config from the AWS CLI
func GetAWSConfig() aws.Config {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("error loading default config")
	}

	return config
}

// GetAWSConfigForRegion pulls the credentials from the AWS
// CLI configuration but allows a region override
func GetAWSConfigForRegion(region string) aws.Config {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("error loading default config")
	}

	config.Region = region

	return config
}

// GetSessionName returns a session name appropriate for use as the
// AssumeRole role session name parameter
func GetSessionName(roleArn string) (string, error) {
	user, _ := os.LookupEnv("USER")
	roleRegex := regexp.MustCompile("arn:aws:iam::[0-9]{12}:role/([a-zA-Z0-9-]+)")
	match := roleRegex.FindAllStringSubmatch(roleArn, -1)

	if !roleRegex.MatchString(roleArn) {
		return "", errors.New("Invalid Role ARN")
	}

	return fmt.Sprintf("federator-%s-%s", user, match[0][1]), nil
}

// AuthWithSTS uses a role ARN and the session with the default creds
// to assume a role.
func AuthWithSTS(roleArn, externalID string, config aws.Config) (*sts.AssumeRoleOutput, error) {
	roleSessionName, roleSessionNameErr := GetSessionName(roleArn)
	if roleSessionNameErr != nil {
		log.Fatalln(roleSessionNameErr.Error())
	}

	svc := sts.NewFromConfig(config)

	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: &roleSessionName,
	}

	if externalID != "" {
		input.ExternalId = &externalID
	}

	return svc.AssumeRole(context.TODO(), input)
}
