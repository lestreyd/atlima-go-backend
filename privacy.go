package main

type PrivacyConfiguration struct {
	// represent a privacy configuration
	// for user with visibility attributes
	// and block/message control
	id                int
	phoneVisibility   int
	emailVisibility   int
	wantToGetMails    int
	whoCanSendMessage []User
	blocked           []User
	meta              Meta
}

type PrivacyConfigurationInterface interface {
	AvailableBy(id int)
	Visibility()
	Phone(phoneVisibility int)
	Email(emailVisibility int)
	GetMails(wantToGetMails int)
	Access()
	CanMessage(whoCanSendMessage []User)
	Blocked(blocked []User)
	WithMeta(meta Meta)
}

type PrivacyBuilder struct {
	privacy *PrivacyConfiguration
}

func (b *PrivacyBuilder) AvailableBy(id int) *PrivacyIdBuilder {
	// create id
	b.privacy.id = id
	return &PrivacyIdBuilder{*b}
}

type PrivacyIdBuilder struct {
	// builder id
	PrivacyBuilder
}

func (b *PrivacyBuilder) Visibility() *PrivacyVisBuilder {
	// create visibility subset
	return &PrivacyVisBuilder{*b}
}

type PrivacyVisBuilder struct {
	// builder meta
	PrivacyBuilder
}

func (b *PrivacyBuilder) Phone(phoneVisibility int) *PrivacyVisPhoneBuilder {
	// create Phone
	b.privacy.phoneVisibility = phoneVisibility
	return &PrivacyVisPhoneBuilder{*b}
}

type PrivacyVisPhoneBuilder struct {
	// builder meta
	PrivacyBuilder
}

func (b *PrivacyBuilder) Email(emailVisibility int) *PrivacyVisEmailBuilder {
	// create Email
	b.privacy.emailVisibility = emailVisibility
	return &PrivacyVisEmailBuilder{*b}
}

type PrivacyVisEmailBuilder struct {
	// builder meta
	PrivacyBuilder
}

func (b *PrivacyBuilder) GetMails(wantToGetMails int) *PrivacyGetMailsBuilder {
	// want to get mails builder
	b.privacy.wantToGetMails = wantToGetMails
	return &PrivacyGetMailsBuilder{*b}
}

type PrivacyGetMailsBuilder struct {
	// builder want to get mails
	PrivacyBuilder
}

func (b *PrivacyBuilder) Access() *PrivacyAccessBuilder {
	return &PrivacyAccessBuilder{*b}
}

type PrivacyAccessBuilder struct {
	// builder access
	PrivacyBuilder
}

func (b *PrivacyBuilder) Blocked(blocked []User) *PrivacyBlockedBuilder {
	// blocked users
	for i := range blocked {
		b.privacy.blocked = append(b.privacy.blocked, blocked[i])
	}
	return &PrivacyBlockedBuilder{*b}
}

type PrivacyBlockedBuilder struct {
	// builder meta
	PrivacyBuilder
}

func (b *PrivacyBuilder) CanMessage(whoCanSendMessage []User) *PrivacyCanMessageBuilder {
	// who can messge?
	for i := range whoCanSendMessage {
		b.privacy.whoCanSendMessage = append(b.privacy.whoCanSendMessage, whoCanSendMessage[i])
	}
	return &PrivacyCanMessageBuilder{*b}
}

type PrivacyCanMessageBuilder struct {
	// builder who can message
	PrivacyBuilder
}

func (b *PrivacyBuilder) WithMeta(meta Meta) *PrivacyMetaBuilder {
	// builder for meta
	b.privacy.meta = meta
	return &PrivacyMetaBuilder{*b}
}

type PrivacyMetaBuilder struct {
	// builder meta
	PrivacyBuilder
}

func (b *PrivacyBuilder) Build() *PrivacyConfiguration {
	// Privacy object
	return b.privacy
}
