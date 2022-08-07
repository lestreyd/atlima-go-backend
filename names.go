package main

type Names struct {
	// represents a names information
	// about user
	firstName        string
	lastName         string
	middleName       string
	nativeFirstName  string
	nativeLastName   string
	nativeMiddleName string
	meta             Meta
}

type NamesInterface interface {
	CalledBy()
	First(firstName string)
	Middle(middleName string)
	Last(lastName string)
	FirstNative(firstNative string)
	LastNative(lastNative string)
	MiddleNative(middleNative string)
	WithMeta(meta Meta)
}
