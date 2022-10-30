package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/YashdalfTheGray/federator/constants"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// CredsDetails holds all of the details we need to print
// credentials once we have federated
type CredsDetails struct {
	ExpiresAfter    *time.Time `json:"expiresAfter"`
	AccessKeyID     string     `json:"AccessKeyId"`
	SecretAccessKey string     `json:"SecretAccessKey"`
	SessionToken    string     `json:"SessionToken"`
}

// NewCredsDetails returns a new CredsDetails object with the
// right things
func NewCredsDetails(out *sts.AssumeRoleOutput) *CredsDetails {
	return &CredsDetails{
		ExpiresAfter:    out.Credentials.Expiration,
		AccessKeyID:     *out.Credentials.AccessKeyId,
		SecretAccessKey: *out.Credentials.SecretAccessKey,
		SessionToken:    *out.Credentials.SessionToken,
	}
}

// ToJSONString converts the struct with the creds to a JSON object
func (c CredsDetails) ToJSONString() (string, error) {
	prettyJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}

	return string(prettyJSON), nil
}

// ToString converts the struct with the creds to a human readable string
func (c CredsDetails) ToString() string {
	result := (fmt.Sprintf("This session will expire at %s", c.ExpiresAfter.Local().String()) + "\n")
	result += (FormatEnvVar(constants.EnvAWSAccessKeyID, c.AccessKeyID) + "\n")
	result += (FormatEnvVar(constants.EnvAWSSecretAccessKey, c.SecretAccessKey) + "\n")
	result += (FormatEnvVar(constants.EnvAWSSessionToken, c.SessionToken) + "\n")
	return result
}

// AwsCliCreds mirrors what we must do for use as a credential_process.
// see: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
// and: https://github.com/aws/aws-sdk-go-v2/blob/b7d8e15425d2f86a0596e8d7db2e33bf382a21dd/credentials/processcreds/provider.go#L152
type AwsCliCreds struct {
	Version         int
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string
	SessionToken    string
	Expiration      *time.Time
}

func NewAwsCliCreds(output *sts.AssumeRoleOutput) *AwsCliCreds {
	return &AwsCliCreds{
		Version:         1,
		AccessKeyID:     *output.Credentials.AccessKeyId,
		SecretAccessKey: *output.Credentials.SecretAccessKey,
		SessionToken:    *output.Credentials.SessionToken,
		Expiration:      output.Credentials.Expiration,
	}
}

func (creds AwsCliCreds) ToJSONString() (string, error) {
	uglyJson, err := json.Marshal(&creds)
	if err != nil {
		return "", err
	}
	return string(uglyJson), nil
}
