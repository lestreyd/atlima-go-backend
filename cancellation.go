package main

type Cancellation struct {
	// disqualification from event
	// can be 0 (DNS - don't started) or 1 (DQ - disqualified)
	id               int
	course           Course
	cancellationType int    // DNS/DQ
	reason           string // Why?
	image            []byte
	referee          RefereeSlot
	meta             Meta
}

type CancellationInterface interface {
	AvailableBy(id int)
	For(course Course)
	What(cancellationType int)
	Why(reason string)
	Reference(image []byte)
	Referee(referee RefereeSlot)
	WithMeta(meta Meta)
}

type CancellationBuilder struct {
	cancellation *Cancellation
}

func (b *CancellationBuilder) AvailableBy(id int) *CancellationIdBuilder {
	// create id
	b.cancellation.id = id
	return &CancellationIdBuilder{*b}
}

type CancellationIdBuilder struct {
	// builder id
	CancellationBuilder
}

func (b *CancellationBuilder) For(course Course) *CancellationCourseBuilder {
	// add course
	b.cancellation.course = course
	return &CancellationCourseBuilder{*b}
}

type CancellationCourseBuilder struct {
	// builder Course
	CancellationBuilder
}

func (b *CancellationBuilder) What(cancellationType int) *CancellationTypeBuilder {
	// add type of Cancellation
	b.cancellation.cancellationType = cancellationType
	return &CancellationTypeBuilder{*b}
}

type CancellationTypeBuilder struct {
	// builder for Cancellation type
	CancellationBuilder
}

func (b *CancellationBuilder) Why(reason string) *CancellationWhyBuilder {
	// contains reason why Cancellation was
	// initiated by referee
	b.cancellation.reason = reason
	return &CancellationWhyBuilder{*b}
}

type CancellationWhyBuilder struct {
	// builder meta
	CancellationBuilder
}

func (b *CancellationBuilder) Reference(image []byte) *CancellationReferenceBuilder {
	// image with photo from phone display for proof
	b.cancellation.image = image
	return &CancellationReferenceBuilder{*b}
}

type CancellationReferenceBuilder struct {
	// builder image
	CancellationBuilder
}

func (b *CancellationBuilder) Referee(referee RefereeSlot) *CancellationRefereeBuilder {
	// add RefereeSlot
	b.cancellation.referee = referee
	return &CancellationRefereeBuilder{*b}
}

type CancellationRefereeBuilder struct {
	// add RefereeSlot builder
	CancellationBuilder
}

func (b *CancellationBuilder) WithMeta(meta Meta) *CancellationMetaBuilder {
	// add meta builder
	b.cancellation.meta = meta
	return &CancellationMetaBuilder{*b}
}

type CancellationMetaBuilder struct {
	// builder meta
	CancellationBuilder
}

func (b *CancellationBuilder) Build() *Cancellation {
	// Cancellation object
	return b.cancellation
}
