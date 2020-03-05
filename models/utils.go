package models

import (
	"fmt"
	"time"
)

// FormatEnvVar prints out a key value pair as an export environment
// variable command
func FormatEnvVar(key, value string) string {
	return fmt.Sprintf("export %s=%s", key, value)
}

// GetExpirationString prints out a user readable string for when
// the given time will arrive
func GetExpirationString(t *time.Time) string {
	return fmt.Sprintf("This session will expire at %s", t.Local().String())
}
