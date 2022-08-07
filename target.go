package main

type TargetType struct {
	// represent Target type
	// structure with allowedResults
	// as "A" or "ACD"
	id             int
	paper          bool
	image          []byte
	allowedResults []string
	content        Content
	meta           Meta
}

type TargetInterface interface {
	AvailableById(id int)
	SetContent(content Content)
	Paper(paper bool)
	Image(image []byte)
	Allowed(allowedResults []string)
	WithMeta(meta Meta)
}

type TargetSet struct {
	// represent set of targets
	// with specific amount and
	// alphaCost for target
	id         int
	targetType TargetType
	course     Course
	amount     int
	alphaCost  int
	meta       Meta
}

type TargetSetInterface interface {
	AvailableBy(id int)
	For(course Course)
	Type(targetType TargetType)
	Amount(amount int)
	AlphaCost(alphaCost int)
	WithMeta(meta Meta)
}
