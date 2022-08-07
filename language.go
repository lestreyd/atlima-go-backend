package main

type Language struct {
	// representation of language structure
	// for multilanguage text and description
	// inside of majority system objects
	id     int
	name   string
	code   string
	weight int
	meta   Meta
}

type LanguageBuilder struct {
	//builder for language
	language *Language
}

func (b *LanguageBuilder) AvailableBy(id int) *LanguageIdBuilder {
	// create id for availability
	b.language.id = id
	return &LanguageIdBuilder{*b}
}

func (b *LanguageBuilder) Name(name string) *LanguageNameBuilder {
	// name for language
	b.language.name = name
	return &LanguageNameBuilder{*b}
}

func (b *LanguageBuilder) Code(code string) *LanguageCodeBuilder {
	// language code
	b.language.code = code
	return &LanguageCodeBuilder{*b}
}

func (b *LanguageBuilder) Weight(weight int) *LanguageWeightBuilder {
	// language weight
	b.language.weight = weight
	return &LanguageWeightBuilder{*b}
}

func (b *LanguageBuilder) Meta(meta Meta) *LanguageMetaBuilder {
	// metadata for language
	b.language.meta = meta
	return &LanguageMetaBuilder{*b}
}

type LanguageIdBuilder struct {
	// builder for id of language
	LanguageBuilder
}

type LanguageNameBuilder struct {
	// builder for name
	LanguageBuilder
}

type LanguageCodeBuilder struct {
	// builder for code
	LanguageBuilder
}

type LanguageWeightBuilder struct {
	// builder for weight
	LanguageBuilder
}

type LanguageMetaBuilder struct {
	// builder for meta
	LanguageBuilder
}

func (b *LanguageBuilder) Build() *Language {
	// provides language object for text blocks
	return b.language
}
