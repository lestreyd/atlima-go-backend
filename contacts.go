package main

type Contacts struct {
	// represents a contacts information
	// about user
	id        int
	email     string
	phone     string
	facebook  string
	vk        string
	instagram string
	meta      Meta
}

type ContactInterface interface {
	AvailableBy(id int)
	Email(email string)
	Phone(phone string)
	Facebook(facebook string)
	Vk(vk string)
	Instagram(instagram string)
	WithMeta(meta Meta)
}

type ContactBuilder struct {
	contact *Contacts
}

func (b *ContactBuilder) AvailableBy(id int) *ContactIdBuilder {
	// add Meta object to SlotResult
	b.contact.id = id
	return &ContactIdBuilder{*b}
}

type ContactIdBuilder struct {
	// builder for meta
	ContactBuilder
}

func (b *ContactBuilder) Email(email string) *ContactEmailBuilder {
	// add Meta object to SlotResult
	b.contact.email = email
	return &ContactEmailBuilder{*b}
}

type ContactEmailBuilder struct {
	// builder for meta
	ContactBuilder
}

func (b *ContactBuilder) Phone(phone string) *ContactPhoneBuilder {
	// add Meta object to SlotResult
	b.contact.phone = phone
	return &ContactPhoneBuilder{*b}
}

type ContactPhoneBuilder struct {
	// builder for meta
	ContactBuilder
}

func (b *ContactBuilder) Facebook(facebook string) *ContactFBBuilder {
	// add Meta object to SlotResult
	b.contact.facebook = facebook
	return &ContactFBBuilder{*b}
}

type ContactFBBuilder struct {
	// builder for meta
	ContactBuilder
}

func (b *ContactBuilder) Vk(vk string) *ContactVkBuilder {
	// add Meta object to SlotResult
	b.contact.vk = vk
	return &ContactVkBuilder{*b}
}

type ContactVkBuilder struct {
	// builder for meta
	ContactBuilder
}

func (b *ContactBuilder) Instagram(instagram string) *ContactInstagramBuilder {
	// add Meta object to SlotResult
	b.contact.instagram = instagram
	return &ContactInstagramBuilder{*b}
}

type ContactInstagramBuilder struct {
	// builder for meta
	ContactBuilder
}

func (b *ContactBuilder) WithMeta(meta Meta) *ContactMetaBuilder {
	// add Meta object to SlotResult
	b.contact.meta = meta
	return &ContactMetaBuilder{*b}
}

type ContactMetaBuilder struct {
	// builder for meta
	ContactBuilder
}

func (b *ContactBuilder) Build() *Contacts {
	// Slot result by specific Course
	return b.contact
}
