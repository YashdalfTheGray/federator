//go:build linux || darwin
// +build linux darwin

package models

import "fmt"

// FormatEnvVar prints out a key value pair as an export environment
// variable command
func FormatEnvVar(key, value string) string {
	return fmt.Sprintf("export %s=%s", key, value)
}
