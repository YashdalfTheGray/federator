package utils_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/YashdalfTheGray/federator/utils"
)

func TestGetAWSSessionName(t *testing.T) {
	testCases := []struct {
		desc, inRoleArn, outRoleSessionName string
		expectError                         bool
	}{
		{
			desc:               "returns a role session name given valid role arn",
			inRoleArn:          "arn:aws:iam::123456789012:role/testRole",
			outRoleSessionName: fmt.Sprintf("federator-%s-testRole", os.Getenv("USER")),
			expectError:        false,
		},
		{
			desc:               "returns an error if the role arn is malformed",
			inRoleArn:          "arn:aws:iam::1234567890:role/testRole",
			outRoleSessionName: "",
			expectError:        true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			roleSessionName, err := utils.GetSessionName(tC.inRoleArn)

			if tC.expectError && err == nil {
				t.Error("Expected error not to be nil")
			}

			if !tC.expectError && roleSessionName != tC.outRoleSessionName {
				t.Errorf("Expected role session name to be %s but got %s", tC.outRoleSessionName, roleSessionName)
			}
		})
	}
}
