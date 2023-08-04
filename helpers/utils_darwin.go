package helpers

import "os"

func GetCurrentUsername() string {
	return os.Getenv("USER")
}
