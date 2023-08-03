//go:build linux || darwin
// +build linux darwin

package utils

import "os"

func GetCurrentUsername() string {
	return os.Getenv("USER")
}
