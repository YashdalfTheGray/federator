package models

import "time"

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
