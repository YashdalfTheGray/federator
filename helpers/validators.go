package helpers

// AvailableRegions is an array of regions that are supported
// Sadly, since the endpoints package doesn't exist in v2, this
// list is maintained by looking up the govcloud/china/iso
// partitions manually, and then running an ec2 describe-regions
// aws ec2 describe-regions --all-regions | jq '.Regions | .[] | .RegionName'
var AvailableRegions = [...]string{
	// aws partition
	"ap-south-2",
	"ap-south-1",
	"eu-south-1",
	"eu-south-2",
	"me-central-1",
	"il-central-1",
	"ca-central-1",
	"eu-central-1",
	"eu-central-2",
	"us-west-1",
	"us-west-2",
	"af-south-1",
	"eu-north-1",
	"eu-west-3",
	"eu-west-2",
	"eu-west-1",
	"ap-northeast-3",
	"ap-northeast-2",
	"me-south-1",
	"ap-northeast-1",
	"sa-east-1",
	"ap-east-1",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-southeast-3",
	"ap-southeast-4",
	"us-east-1",
	"us-east-2",

	// govcloud partition
	"us-gov-east-1",
	"us-gov-west-1",

	// china partition
	"cn-north-1",
	"cn-northwest-1",

	// iso partitions
	"us-iso-east-1",
	"us-iso-west-1",
	"us-isob-east-1",
}

// IsRegionValid validates the region passed in as a
// command line argument
func IsRegionValid(region string) bool {
	for _, regionOption := range AvailableRegions {
		if regionOption == region {
			return true
		}
	}
	return false
}
