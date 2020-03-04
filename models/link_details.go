package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// LinkDetails is a struct that contains the right things to
// help print out the login link once we have federated
type LinkDetails struct {
	ExpiresAfter *time.Time `json:"expiresAfter"`
	LoginURL     string     `json:"loginUrl"`
}

// NewLinkDetails returns a new LinkDetails with the right
// things stored
func NewLinkDetails(expiresAfter *time.Time, url string) *LinkDetails {
	return &LinkDetails{
		ExpiresAfter: expiresAfter,
		LoginURL:     url,
	}
}

// ToJSONString converts the struct with the creds to a JSON object
func (l LinkDetails) ToJSONString() (string, error) {
	prettJSON, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "", nil
	}

	return string(prettJSON), nil
}

// ToString converts the struct with the creds to a human readable string
func (l LinkDetails) ToString() string {
	result := fmt.Sprintf("This session will expire at %s", l.ExpiresAfter.Local().String())
	result += "\n"
	result += l.LoginURL
	return result
}
