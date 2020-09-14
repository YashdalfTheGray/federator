package models

import "encoding/json"

type statementCondition struct {
	StringEquals struct {
		ExternalID string `json:"sts:ExternalId"`
	}
}

type policyStatement struct {
	Effect    string
	Resource  string
	Action    []string
	Principal struct {
		AWS []string
	}
	Condition []statementCondition
}

// TrustPolicy models a trust policy that can be added
// to an IAM role, although somewhat primitively
type TrustPolicy struct {
	Version   string
	Statement []policyStatement
}

// NewTrustPolicy creates a new trust policy that trusts a specified resource
// with an optional external ID. The resource could be an account ID or an
// IAM user or role ARN.
//
// Examples of valid resources are
//  "123456789012"
//  "arn:aws:iam::123456789012:root"
//  "arn:aws:iam::AWS-account-ID:user/user-name"
//  "arn:aws:iam::AWS-account-ID:role/role-name"
func NewTrustPolicy(resourceToTrust, externalID string) TrustPolicy {
	result := TrustPolicy{
		Version: "2012-10-17",
		Statement: []policyStatement{
			{
				Effect: "Allow",
				Action: []string{"sts:AssumeRole"},
				Principal: struct{ AWS []string }{
					AWS: []string{resourceToTrust},
				},
			},
		},
	}

	if externalID != "" {
		result.Statement[0].Condition[0].StringEquals.ExternalID = externalID
	}

	return result
}

// ToJSONString converts the struct with the creds to a JSON object
func (tp TrustPolicy) ToJSONString() (string, error) {
	prettyJSON, err := json.MarshalIndent(tp, "", "  ")
	if err != nil {
		return "", nil
	}

	return string(prettyJSON), nil
}
