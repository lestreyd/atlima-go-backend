package main

type RefereeSlot struct {
	// represents referee function for
	// user in specific event with
	// specific role
	id   int
	user User
	role int
	meta Meta
}
