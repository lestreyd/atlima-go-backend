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
	content       []Content
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

type CourseBuilder struct {
	course *Course
}

func (b *CourseBuilder) AvailableBy(id int) *CourseIdBuilder {
	// create id for availability
	b.course.id = id
	return &CourseIdBuilder{*b}
}

type CourseIdBuilder struct {
	// builder for id of language
	CourseBuilder
}

func (b *CourseBuilder) WithTitle(title string) *CourseTitleBuilder {
	// create id for availability
	b.course.title = title
	return &CourseTitleBuilder{*b}
}

type CourseTitleBuilder struct {
	// builder for id of language
	CourseBuilder
}

func (b *CourseBuilder) WithImage(image []byte) *CourseImageBuilder {
	// create id for availability
	b.course.image = image
	return &CourseImageBuilder{*b}
}

type CourseImageBuilder struct {
	// builder for id of language
	CourseBuilder
}

func (b *CourseBuilder) WithTargetSets(targetSets []TargetSet) *CourseTargetSetBuilder {
	// create id for availability
	for i := range targetSets {
		b.course.targetSets = append(b.course.targetSets, targetSets[i])
	}
	return &CourseTargetSetBuilder{*b}
}

type CourseTargetSetBuilder struct {
	// builder for id of language
	CourseBuilder
}

func (b *CourseBuilder) WithContent(content []Content) *CourseContentBuilder {
	// create id for availability
	for i := range content {
		b.course.content = append(b.course.content, content[i])
	}
	return &CourseContentBuilder{*b}
}

type CourseContentBuilder struct {
	// builder for id of language
	CourseBuilder
}

func (b *CourseBuilder) WithMeta(meta Meta) *CourseMetaBuilder {
	// create id for availability
	b.course.meta = meta
	return &CourseMetaBuilder{*b}
}

type CourseMetaBuilder struct {
	// builder for id of language
	CourseBuilder
}
