package main

type CourseResult struct {
	// course fix result for slot
	// in specific event
	id     int
	course Course
	A      int
	C      int
	D      int
	M      int
	NS     int
	T      float32
	image  []byte
	// bunch of penalties
	penalties []Penalty
	referee   RefereeSlot
	meta      Meta
}

type CourseResultInterface interface {
	AvailableBy(id int)
	For(course Course)
	Set()
	A(A int)
	C(C int)
	D(D int)
	M(M int)
	NS(NS int)
	T(T float32)
	Reference(image []byte)
	Penalties(penalties []Penalty)
	Referee(referee RefereeSlot)
	WithMeta(meta Meta)
}

type CourseResultBuilder struct {
	result *CourseResult
}

func (b *CourseResultBuilder) AvailableBy(id int) *CourseResultIdBuilder {
	// create Meta object for Slot
	b.result.id = id
	return &CourseResultIdBuilder{*b}
}

type CourseResultIdBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) For(course Course) *CourseResultCourseBuilder {
	// create Meta object for Slot
	b.result.course = course
	return &CourseResultCourseBuilder{*b}
}

type CourseResultCourseBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) Set() *CourseResultSetBuilder {
	// create Meta object for Slot
	return &CourseResultSetBuilder{*b}
}

type CourseResultSetBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) A(A int) *CourseResultABuilder {
	b.result.A = A
	// create Meta object for Slot
	return &CourseResultABuilder{*b}
}

type CourseResultABuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) C(C int) *CourseResultCBuilder {
	// create Meta object for Slot
	b.result.C = C
	return &CourseResultCBuilder{*b}
}

type CourseResultCBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) D(D int) *CourseResultDBuilder {
	// create Meta object for Slot
	b.result.D = D
	return &CourseResultDBuilder{*b}
}

type CourseResultDBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) M(M int) *CourseResultMBuilder {
	// create Meta object for Slot
	b.result.M = M
	return &CourseResultMBuilder{*b}
}

type CourseResultMBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) NS(NS int) *CourseResultNSBuilder {
	// create Meta object for Slot
	b.result.NS = NS
	return &CourseResultNSBuilder{*b}
}

type CourseResultNSBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) T(T float32) *CourseResultTBuilder {
	// create Meta object for Slot
	b.result.T = T
	return &CourseResultTBuilder{*b}
}

type CourseResultTBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) Reference(image []byte) *CourseResultImageBuilder {
	// create Meta object for Slot
	b.result.image = image
	return &CourseResultImageBuilder{*b}
}

type CourseResultImageBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) Penalties(penalties []Penalty) *CourseResultPenaltiesBuilder {
	// create Meta object for Slot
	for i := range penalties {
		b.result.penalties = append(b.result.penalties, penalties[i])
	}
	return &CourseResultPenaltiesBuilder{*b}
}

type CourseResultPenaltiesBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) Referee(referee RefereeSlot) *CourseResultRefereeBuilder {
	// create Meta object for Slot
	b.result.referee = referee
	return &CourseResultRefereeBuilder{*b}
}

type CourseResultRefereeBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) WithMeta(meta Meta) *CourseResultMetaBuilder {
	// create Meta object for Slot
	b.result.meta = meta
	return &CourseResultMetaBuilder{*b}
}

type CourseResultMetaBuilder struct {
	// builder meta
	CourseResultBuilder
}

func (b *CourseResultBuilder) Build() *CourseResult {
	// Slot object
	return b.result
}
