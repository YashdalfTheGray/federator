//go:build windows
// +build windows

package models

import "fmt"

// FormatEnvVar prints out a key value pair as an export environment
// variable command, this is specific to powershell
func FormatEnvVar(key, value string) string {
	return fmt.Sprintf("$Env:%s=\"%s\"", key, value)
}
