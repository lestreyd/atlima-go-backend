package main

type Invite struct {
	// representation of Event Referee invite
	user          User
	role          int
	status        int
	dismissReason string
	meta          Meta
	createdBy     User
}

type InviteInterface interface {
	ProvideTo(user User)
	Role(role int)
	InStatus(status int)
	WithDismissReason(dismissReason string)
	withMeta(meta Meta)
	CreatedBy(user User)
}
