package utils_test

import (
	"testing"

	"github.com/YashdalfTheGray/federator/utils"
)

func TestValidateRegions(t *testing.T) {
	testCases := []struct {
		desc, input string
		expected    bool
	}{
		{
			desc:     "returns true when a valid region is passed in",
			input:    "us-west-2",
			expected: true,
		},
		{
			desc:     "returns false when an invalid region is passed in",
			input:    "test-region",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := utils.ValidateRegion(tC.input); result != tC.expected {
				t.Errorf("Expected the result to be %t but wasn't for input %s", tC.expected, tC.input)
			}
		})
	}
}
