package main

import "time"

type PrivacyConfiguration struct {
	// represent a privacy configuration
	// for user with visibility attributes
	// and block/message control
	phoneVisibility   int
	emailVisibility   int
	wantToGetMails    int
	whoCanSendMessage []User
	blocked           []User
	createdAt         time.Time
	updatedAt         time.Time
	meta              Meta
}
