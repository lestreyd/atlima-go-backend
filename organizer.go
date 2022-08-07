package main

type Organizer struct {
	// represents organizer
	// which can create new events
	// and manage it in platform
	id             int
	slug           string
	site           string
	image          []byte
	country        Country
	region         Region
	city           City
	contacts       Contacts
	content        []Content
	administrators []User
	meta           Meta
}
