package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// GetAWSSession returns a session that uses the currently configured
// AWS CLI credentials
func GetAWSSession(options session.Options) *session.Session {
	return session.Must(session.NewSessionWithOptions(options))
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
func AuthWithSTS(roleArn string) (*sts.AssumeRoleOutput, error) {
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

	return stsClient.AssumeRole(serviceAssumeRoleInput)
}
