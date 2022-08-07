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
