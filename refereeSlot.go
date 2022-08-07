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

type RefereeSlotInterface interface {
	AvailableBy(id int)
	For(user User)
	GrantedRole(role int)
	WithMeta(meta Meta)
}
