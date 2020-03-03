package models

import "time"

// LinkDetails is a struct that contains the right things to
// help print out the login link once we have federated
type LinkDetails struct {
	ExpiresAfter *time.Time `json:"expiresAfter"`
	LoginURL     string     `json:"loginUrl"`
}
