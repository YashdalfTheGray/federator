package utils

// AvailableRegions is an array of regions that are supported
var AvailableRegions = [...]string{
	"af-south-1",
	"ap-northeast-1",
	"ap-southeast-2",
	"aws-global",
	"ca-central-1",
	"eu-south-1",
	"eu-west-3",
	"me-south-1",
	"ap-east-1",
	"ap-south-1",
	"sa-east-1",
	"us-east-2-fips",
	"eu-north-1",
	"eu-west-1",
	"eu-west-2",
	"us-east-1-fips",
	"us-east-2",
	"us-west-1-fips",
	"ap-northeast-2",
	"ap-southeast-1",
	"us-west-1",
	"us-west-2",
	"us-west-2-fips",
	"eu-central-1",
	"us-east-1",
}

// ValidateRegion validates the region passed in as a
// command line argument
func ValidateRegion(region string) bool {
	for _, regionOption := range AvailableRegions {
		if regionOption == region {
			return true
		}
	}
	return false
}
