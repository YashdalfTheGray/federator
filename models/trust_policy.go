package models

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
