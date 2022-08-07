package main

type Content struct {
	// represent set of title and description
	// with possibility choose appropriate language
	id          int
	language    Language
	title       string
	description string
	meta        Meta
}

type ContentInterface interface {
	AvailableBy(id int)
	WithLanguage(language Language)
	WithTitle(title, description string)
	WithMeta(meta Meta)
}

type ContentBuilder struct {
	content *Content
}

func (b *ContentBuilder) AvailableBy(id int) *ContentIdBuilder {
	// create id for availability
	b.content.id = id
	return &ContentIdBuilder{*b}
}

type ContentIdBuilder struct {
	// builder for id of language
	ContentBuilder
}

func (b *ContentBuilder) WithLanguage(language Language) *ContentLanguageBuilder {
	// create id for availability
	b.content.language = language
	return &ContentLanguageBuilder{*b}
}

type ContentLanguageBuilder struct {
	// builder for id of language
	ContentBuilder
}

func (b *ContentBuilder) Title(title string, description string) *ContentTitleBuilder {
	//create title for multilingual text
	b.content.title = title
	b.content.description = description
	return &ContentTitleBuilder{*b}
}

type ContentTitleBuilder struct {
	// builder for id of language
	ContentBuilder
}

func (b *ContentBuilder) Meta(meta Meta) *ContentMetaBuilder {
	//create meta for multilingual text
	b.content.meta = meta
	return &ContentMetaBuilder{*b}
}

type ContentMetaBuilder struct {
	// builder for id of language
	ContentBuilder
}
