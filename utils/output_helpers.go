package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/sts"

	"github.com/YashdalfTheGray/federator/constants"
)

// PrintCredsFromSTSOutput prints out the credentials we got from the
// STS output in a way that the user can export them in to the shell
// as well as the expiration information about the session
func PrintCredsFromSTSOutput(out *sts.AssumeRoleOutput) {
	fmt.Println("Successfully authenticated with STS. Commands to use below.")
	fmt.Println(fmt.Sprintf("This session will expire at %s", out.Credentials.Expiration.Local().String()))
	fmt.Println(fmt.Sprintf("export %s=%s", constants.EnvAWSAccessKeyID, *out.Credentials.AccessKeyId))
	fmt.Println(fmt.Sprintf("export %s=%s", constants.EnvAWSSecretAccessKey, *out.Credentials.SecretAccessKey))
	fmt.Println(fmt.Sprintf("export %s=%s", constants.EnvAWSSessionToken, *out.Credentials.SessionToken))
}

// PrintLoginURLDetails prints out the login URL as well as the expiration date
// of the session
func PrintLoginURLDetails(out *sts.AssumeRoleOutput, loginURL string) {
	fmt.Println("Successfully authenticated with STS. Login URL below.")
	fmt.Println(fmt.Sprintf("This session will expire at %s", out.Credentials.Expiration.Local().String()))
	fmt.Println(loginURL)
}
