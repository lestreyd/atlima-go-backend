package main

type Course struct {
	// represents course with specific
	// title and set of attributes
	// for IPSC as scoringShoots or
	// minimumShoots
	id            int
	number        int
	title         string
	scoringShoots int
	minimumShoots int
	image         []byte
	targetSets    []TargetSet
	content       Content
	meta          Meta
}

type CourseInterface interface {
	AvailableBy(id int)
	WithTitle(title string)
	WithImage(image []byte)
	WithTargetSets(targetSet []TargetSet)
	WithContent(content Content)
	WithMeta(meta Meta)
}
