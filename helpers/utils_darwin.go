//go:build darwin
// +build darwin

package helpers

import "os"

func GetCurrentUsername() string {
	return os.Getenv("USER")
}
