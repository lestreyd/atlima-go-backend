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

type OrganizerInterface interface {
	// AvailableBy contains all need information
	// about organizer availability
	// With Id, Slug and Site you can identify
	// organizer strictly
	AvailableBy()
	Id(id int)
	Slug(slug string)
	Site(site string)
	Image(image []byte)
	// Located store full information about
	// Country, Region and City of Organizer
	Located()
	// InCountry , InRegion and InCity
	// helps to set location of Organizer
	InCountry(country Country)
	InRegion(region Region)
	InCity(city City)
	// WithContacts provide phone and email
	// fields for storing data about contacts
	WithContacts(contacts Contacts)
	// WithContent :Multilingual text can be storing here
	WithContent(content Content)
	// AdministratedBy usually point at array
	// of users with administrator rights for
	// editing organizer and events for this
	// organizer
	AdministratedBy(administrators []User)
	// WithMeta fields: showAs,
	WithMeta(meta Meta)
}

type OrganizerBuilder struct {
	organizer *Organizer
}
