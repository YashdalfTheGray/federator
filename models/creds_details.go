package models

import (
	"time"

	"github.com/aws/aws-sdk-go/service/sts"
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
