package main

type Confirmation struct {
	// information about User confirmation
	// activity (phone/email)
	action string
	data   string
	target User
	status uint
	meta   Meta
}

type ConfirmationInterface interface {
	ProvideTo(user User)
	Action(action string)
	For(data string)
	InStatus(status uint)
	WithMeta(meta Meta)
}
