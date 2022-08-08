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
	WithContent(content []Content)
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

func (b *OrganizerBuilder) AvailableBy() *OrganizerAvailabilityBuilder {
	// builder for meta
	return &OrganizerAvailabilityBuilder{*b}
}

type OrganizerAvailabilityBuilder struct {
	// builder meta
	OrganizerBuilder
}

func (b *OrganizerBuilder) Id(id int) *OrganizerAvailabilityBuilder {
	// builder for meta
	b.organizer.id = id
	return &OrganizerAvailabilityBuilder{*b}
}

func (b *OrganizerBuilder) Slug(slug string) *OrganizerAvailabilityBuilder {
	// builder for meta
	b.organizer.slug = slug
	return &OrganizerAvailabilityBuilder{*b}
}

func (b *OrganizerBuilder) Site(site string) *OrganizerAvailabilityBuilder {
	// builder for meta
	b.organizer.site = site
	return &OrganizerAvailabilityBuilder{*b}
}

func (b *OrganizerBuilder) Located() *OrganizerLocationBuilder {
	return &OrganizerLocationBuilder{*b}
}

type OrganizerLocationBuilder struct {
	// builder meta
	OrganizerBuilder
}

func (b *OrganizerBuilder) InCountry(country Country) *OrganizerLocationBuilder {
	b.organizer.country = country
	return &OrganizerLocationBuilder{*b}
}

func (b *OrganizerBuilder) InRegion(region Region) *OrganizerLocationBuilder {
	b.organizer.region = region
	return &OrganizerLocationBuilder{*b}
}

func (b *OrganizerBuilder) InCity(city City) *OrganizerLocationBuilder {
	b.organizer.city = city
	return &OrganizerLocationBuilder{*b}
}

func (b *OrganizerBuilder) WithContacts(contacts Contacts) *OrganizerContactsBuilder {
	// add Meta object to SlotResult
	b.organizer.contacts = contacts
	return &OrganizerContactsBuilder{*b}
}

type OrganizerContactsBuilder struct {
	// builder for meta
	OrganizerBuilder
}

func (b *OrganizerBuilder) WithContent(contents []Content) *OrganizerContentBuilder {
	// add Meta object to SlotResult
	for i := range contents {
		b.organizer.content = append(b.organizer.content, contents[i])
	}

	return &OrganizerContentBuilder{*b}
}

type OrganizerContentBuilder struct {
	// builder for meta
	OrganizerBuilder
}

func (b *OrganizerBuilder) AdministratedBy(users []User) *OrganizerAdministratorsBuilder {
	// add Meta object to SlotResult
	for i := range users {
		b.organizer.administrators = append(b.organizer.administrators, users[i])
	}
	return &OrganizerAdministratorsBuilder{*b}
}

type OrganizerAdministratorsBuilder struct {
	// builder for meta
	OrganizerBuilder
}

func (b *OrganizerBuilder) WithMeta(meta Meta) *OrganizerMetaBuilder {
	// add Meta object to SlotResult
	b.organizer.meta = meta
	return &OrganizerMetaBuilder{*b}
}

type OrganizerMetaBuilder struct {
	// builder for meta
	OrganizerBuilder
}

func (b *OrganizerBuilder) Build() *Organizer {
	// Privacy object
	return b.organizer
}
