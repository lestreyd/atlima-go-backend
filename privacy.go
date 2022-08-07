package main

type PrivacyConfiguration struct {
	// represent a privacy configuration
	// for user with visibility attributes
	// and block/message control
	phoneVisibility   int
	emailVisibility   int
	wantToGetMails    int
	whoCanSendMessage []User
	blocked           []User
	meta              Meta
}

type PrivacyConfigurationInterface interface {
	Visibility()
	Phone(phoneVisibility int)
	Email(emailVisibility int)
	GetMails(wantToGetMails int)
	Access()
	CanMessage(whoCanSendMessage []User)
	Blocked(blocked []User)
	WithMeta(meta Meta)
}
