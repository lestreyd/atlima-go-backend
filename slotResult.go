package main

type SlotResult struct {
	// representation of Course result
	// for user in event
	id           int
	course       Course
	coursePoints int
	stagePoints  float32
	hitFactor    float32
	meta         Meta
}

type SlotResultInterface interface {
	AvailableBy(id int)
	For(course Course)
	Result(coursePoints int)
	StagePoints(stagePoints float32)
	HitFactor(hitFactor float32)
	WithMeta(meta Meta)
}

type SlotResultBuilder struct {
	slotResult *SlotResult
}

func (b *SlotResultBuilder) AvailableBy(id int) *SlotResultIdBuilder {
	// add id to slot result
	b.slotResult.id = id
	return &SlotResultIdBuilder{*b}
}

type SlotResultIdBuilder struct {
	// add id to slot result
	SlotResultBuilder
}

func (b *SlotResultBuilder) For(course Course) *SlotResultCourseBuilder {
	// add course to slot result
	b.slotResult.course = course
	return &SlotResultCourseBuilder{*b}
}

type SlotResultCourseBuilder struct {
	// builder for slot result
	SlotResultBuilder
}

func (b *SlotResultBuilder) Result(coursePoints int) *SlotResultResultBuilder {
	// add course points (course absolute points)
	b.slotResult.coursePoints = coursePoints
	return &SlotResultResultBuilder{*b}
}

type SlotResultResultBuilder struct {
	// builder for slot result
	SlotResultBuilder
}

func (b *SlotResultBuilder) StagePoints(coursePoints int) *SlotResultStagePointsBuilder {
	// add stage points
	b.slotResult.coursePoints = coursePoints
	return &SlotResultStagePointsBuilder{*b}
}

type SlotResultStagePointsBuilder struct {
	// builder for stage points
	SlotResultBuilder
}

func (b *SlotResultBuilder) HitFactor(hitFactor float32) *SlotResultHFBuilder {
	// add hit factor
	b.slotResult.hitFactor = hitFactor
	return &SlotResultHFBuilder{*b}
}

type SlotResultHFBuilder struct {
	// builder for hit factor
	SlotResultBuilder
}

func (b *SlotResultBuilder) WithMeta(meta Meta) *SlotResultMetaBuilder {
	// add Meta object to SlotResult
	b.slotResult.meta = meta
	return &SlotResultMetaBuilder{*b}
}

type SlotResultMetaBuilder struct {
	// builder for meta
	SlotResultBuilder
}

func (b *SlotResultBuilder) Build() *SlotResult {
	// Slot result by specific Course
	return b.slotResult
}
