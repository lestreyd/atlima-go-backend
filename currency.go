package main

type Currency struct {
	// information about currency in order
	id          int
	digitalCode uint
	code        string
	content     []Content
	weight      int
	meta        Meta
}

type CurrencyInterface interface {
	AvailableBy()
	Id(id int)
	DigitalCode(digitalCode uint)
	Code(code string)
	WithContent(content Content)
	WithWeight(weight int)
	WithMeta(meta Meta)
}
