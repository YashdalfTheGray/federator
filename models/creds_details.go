package models

import (
	"encoding/json"
	"fmt"
	"time"

	stsv2 "github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/service/sts"

	"github.com/YashdalfTheGray/federator/constants"
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

// NewCredsDetailsv2 returns a new CredsDetails object with the
// right things
func NewCredsDetailsv2(out *stsv2.AssumeRoleResponse) *CredsDetails {
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
		return "", nil
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
