package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	stsv2 "github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/aws/session"
	sts "github.com/aws/aws-sdk-go/service/sts"
)

// GetAWSSession returns a session that uses the currently configured
// AWS CLI credentials
func GetAWSSession(options session.Options) *session.Session {
	return session.Must(session.NewSessionWithOptions(options))
}

// GetAWSConfig pulls the default config from the AWS CLI
func GetAWSConfig() aws.Config {
	config, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("error loading default config")
	}

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
func AuthWithSTS(roleArn, externalID string) (*sts.AssumeRoleOutput, error) {
	sesh := GetAWSSession(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})

	roleSessionName, roleSessionNameErr := GetSessionName(roleArn)
	if roleSessionNameErr != nil {
		log.Fatalln(roleSessionNameErr.Error())
	}

	stsClient := sts.New(sesh)

	serviceAssumeRoleInput := &sts.AssumeRoleInput{
		RoleArn:         &roleArn,
		RoleSessionName: &roleSessionName,
	}

	if externalID != "" {
		serviceAssumeRoleInput.ExternalId = &externalID
	}

	return stsClient.AssumeRole(serviceAssumeRoleInput)
}
