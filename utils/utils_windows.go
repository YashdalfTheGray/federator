//go:build windows
// +build windows

package utils

import "os"

func GetCurrentUsername() string {
	return os.Getenv("USERNAME")
}
