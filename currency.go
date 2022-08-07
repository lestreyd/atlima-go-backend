package main

type Currency struct {
	// information about currency in order
	id          int
	digitalCode int
	code        string
	content     []Content
	weight      int
	meta        Meta
}
