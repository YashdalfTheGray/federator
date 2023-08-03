package models

import (
	"fmt"
	"time"
)

// GetExpirationString prints out a user readable string for when
// the given time will arrive
func GetExpirationString(t *time.Time) string {
	return fmt.Sprintf("This session will expire at %s", t.Local().String())
}
