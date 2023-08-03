//go:build windows
// +build windows

package helpers

import "os"

func GetCurrentUsername() string {
	return os.Getenv("USERNAME")
}
