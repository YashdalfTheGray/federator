//go:build linux || darwin
// +build linux darwin

package helpers

import "os"

func GetCurrentUsername() string {
	return os.Getenv("USER")
}
