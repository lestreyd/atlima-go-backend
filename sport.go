package main

// sport type with builder pattern support
// each instance can be created with SportBuilder
// by call SportBuilder.
//				AvailableBy().
//						Id(1).
//						Site('www.atlima.com').
//						Slug('ipsd').
//				WithImage(img).
//				Moderated(false).
//				WithContent(content).
//				AdministratedBy(sportAdmins)

type Sport struct {
	// representation of Sport class
	// contains multilingual content
	// and image for showing in app
	id             int
	site           string
	slug           string
	image          []byte
	content        []Content
	moderated      bool
	administrators []SportAdministrator
	meta           Meta
}

type SportInterface interface {
	AvailableBy()
	Id(id int)
	Site(site string)
	Slug(slug string)
	WithImage(image []byte)
	Moderated(moderated bool)
	WithContent(content []Content)
	AdministratedBy([]SportAdministrator)
}

type SportBuilder struct {
	sport *Sport
}

// builder for sport

func (b *SportBuilder) AvailableBy() *SportAvailabilityBuilder {
	// create id, slug and site for sport availability
	return &SportAvailabilityBuilder{*b}
}

type SportAvailabilityBuilder struct {
	// builder for sport credentials
	SportBuilder
}

func (a *SportAvailabilityBuilder) Id(id int) *SportAvailabilityBuilder {
	// add unique identifier
	a.sport.id = id
	return a
}

func (a *SportAvailabilityBuilder) Site(site string) *SportAvailabilityBuilder {
	// add site for sport
	a.sport.site = site
	return a
}

func (a *SportAvailabilityBuilder) Slug(slug string) *SportAvailabilityBuilder {
	// add slug for sport
	a.sport.slug = slug
	return a
}

func (b *SportBuilder) WithImage(image []byte) *SportImageBuilder {
	// add image for sport in builder
	b.sport.image = image
	return &SportImageBuilder{*b}
}

type SportImageBuilder struct {
	// add image for sport in builder
	SportBuilder
}

func (b *SportBuilder) Moderated(moderated bool) *SportModerationBuilder {
	// set moderation status propagated to events
	b.sport.moderated = moderated
	return &SportModerationBuilder{*b}
}

type SportModerationBuilder struct {
	// set moderation status propagated to events
	// false - events not moderated
	// true - events are moderated
	SportBuilder
}

func (b *SportBuilder) WithContent(content []Content) *SportContentBuilder {
	// add content array
	for i := range content {
		b.sport.content = append(b.sport.content, content[i])
	}
	return &SportContentBuilder{*b}
}

type SportContentBuilder struct {
	// builder for country of sport administrator
	SportBuilder
}

func (b *SportBuilder) AdministratedBy(sportAdministrator []SportAdministrator) *SportAdministratorForSportBuilder {
	// add administrators to administrators array
	for i := range sportAdministrator {
		b.sport.administrators = append(b.sport.administrators, sportAdministrator[i])
	}
	return &SportAdministratorForSportBuilder{*b}
}

type SportAdministratorForSportBuilder struct {
	// builder for administrators
	SportBuilder
}

func (b *SportBuilder) Build() *Sport {
	// SportBuilder object
	return b.sport
}
