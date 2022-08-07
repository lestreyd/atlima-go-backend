package main

import "time"

// meta attached at every object in system
// and can be created with constructor
// as MetaBuilder.showedAs('Hello').CreatedAt().UpdatedAt(now)
// or meta.UpdatedAt(now) if I want update timestamp
// on meta only.
// So we can get sport.meta.createdAt or sport.meta.updatedAt
// or, if we want to represent object name, call
// sport.meta.showedAs

type Meta struct {
	// string representation information
	// and datetime attributes
	id        int
	showedAs  string
	createdAt time.Time
	updatedAt time.Time
}

type MetaBuilder struct {
	meta *Meta
}

func (b *MetaBuilder) AvailableBy(id int) *MetaIdBuilder {
	// builder for showedAs field
	b.meta.id = id
	return &MetaIdBuilder{*b}
}

func (b *MetaBuilder) Show(showedAs string) *MetaStringBuilder {
	// builder for showedAs field
	b.meta.showedAs = showedAs
	return &MetaStringBuilder{*b}
}

func (b *MetaBuilder) CreatedAt() *MetaCreatedAtBuilder {
	// builder for createdAt field
	b.meta.createdAt = time.Now()
	return &MetaCreatedAtBuilder{*b}
}

func (b *MetaBuilder) UpdatedAt(updatedAt time.Time) *MetaUpdatedAtBuilder {
	b.meta.updatedAt = updatedAt
	// builder for updatedAt field
	return &MetaUpdatedAtBuilder{*b}
}

type MetaIdBuilder struct {
	// builder for updated timestamp
	MetaBuilder
}

type MetaStringBuilder struct {
	// builder for showedAs
	MetaBuilder
}

type MetaCreatedAtBuilder struct {
	// builder for updated timestamp
	MetaBuilder
}

type MetaUpdatedAtBuilder struct {
	// builder for updated timestamp
	MetaBuilder
}

func (b *MetaBuilder) Build() *Meta {
	// provides meta object for entity
	return b.meta
}

type FrontendLog struct {
	// grouped by hash error log for
	// frontend with hash and user
	// information
	id         int
	hash       string
	user       User
	message    string
	stackTrace string
	build      int
	logDate    time.Time
	errorCode  string
	counts     int
	meta       Meta
}
