package models

// TrustPolicy models a trust policy that can be added
// to an IAM role, although somewhat primitively
type TrustPolicy struct {
	Version   string
	Statement []struct {
		Effect    string
		Resource  string
		Action    []string
		Principal struct {
			AWS []string
		}
		Condition []struct{}
	}
}
