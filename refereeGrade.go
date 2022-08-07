package main

type RefereeGrade struct {
	// shows info about grade from
	// main referee for referees
	// in event
	id          int
	refereeSlot RefereeSlot
	event       Event
	grade       int
	meta        Meta
}

type RefereeGradeInterface interface {
	AvailableBy(id int)
	For(refereeSlot RefereeSlot)
	On(event Event)
	Grade(grade int)
	WithMeta(meta Meta)
}

type RefereeGradeBuilder struct {
	refereeGrade *RefereeGrade
}

func (b *RefereeGradeBuilder) AvailableBy(id int) *RefereeGradeIdBuilder {
	// add course to slot result
	b.refereeGrade.id = id
	return &RefereeGradeIdBuilder{*b}
}

type RefereeGradeIdBuilder struct {
	// builder for slot result
	RefereeGradeBuilder
}

func (b *RefereeGradeBuilder) For(refereeSlot RefereeSlot) *RefereeGradeForBuilder {
	// add course to slot result
	b.refereeGrade.refereeSlot = refereeSlot
	return &RefereeGradeForBuilder{*b}
}

type RefereeGradeForBuilder struct {
	// builder for slot result
	RefereeGradeBuilder
}

func (b *RefereeGradeBuilder) On(event Event) *RefereeGradeEventBuilder {
	// add course to slot result
	b.refereeGrade.event = event
	return &RefereeGradeEventBuilder{*b}
}

type RefereeGradeEventBuilder struct {
	// builder for slot result
	RefereeGradeBuilder
}

func (b *RefereeGradeBuilder) Grade(grade int) *RefereeGradeGradeBuilder {
	// add course to slot result
	b.refereeGrade.grade = grade
	return &RefereeGradeGradeBuilder{*b}
}

type RefereeGradeGradeBuilder struct {
	// builder for slot result
	RefereeGradeBuilder
}

func (b *RefereeGradeBuilder) WithMeta(meta Meta) *RefereeGradeMetaBuilder {
	// add course to slot result
	b.refereeGrade.meta = meta
	return &RefereeGradeMetaBuilder{*b}
}

type RefereeGradeMetaBuilder struct {
	// builder for slot result
	RefereeGradeBuilder
}

func (b *RefereeGradeBuilder) Build() *RefereeGrade {
	// Slot result by specific Course
	return b.refereeGrade
}
