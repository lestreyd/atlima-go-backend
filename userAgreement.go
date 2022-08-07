package main

type UserAgreement struct {
	id      int
	slug    string
	content []Content
	meta    Meta
}

type UserAgreementInterface interface {
	AvailableById(id int)
	Slug(slug string)
	WithContent(content []Content)
	WithMeta(meta Meta)
}

type UserAgreementBuilder struct {
	userAgreement *UserAgreement
}

func (b *UserAgreementBuilder) AvailableBy(id int) *UserAgreementIdBuilder {
	// set id for offer
	b.userAgreement.id = id
	return &UserAgreementIdBuilder{*b}
}

type UserAgreementIdBuilder struct {
	// offer id builder
	UserAgreementBuilder
}

func (b *UserAgreementBuilder) Slug(slug string) *UserAgreementSlugBuilder {
	// set id for offer
	b.userAgreement.slug = slug
	return &UserAgreementSlugBuilder{*b}
}

type UserAgreementSlugBuilder struct {
	// offer id builder
	UserAgreementBuilder
}

func (b *UserAgreementBuilder) WithContent(content []Content) *UserAgreementContentBuilder {
	// set id for offer
	for i := range content {
		b.userAgreement.content = append(b.userAgreement.content, content[i])
	}
	return &UserAgreementContentBuilder{*b}
}

type UserAgreementContentBuilder struct {
	// offer id builder
	UserAgreementBuilder
}

func (b *UserAgreementBuilder) WithMeta(meta Meta) *UserAgreementMetaBuilder {
	// set id for offer
	b.userAgreement.meta = meta
	return &UserAgreementMetaBuilder{*b}
}

type UserAgreementMetaBuilder struct {
	// offer id builder
	UserAgreementBuilder
}

func (b *UserAgreementBuilder) Build() *UserAgreement {
	// offer builder
	return b.userAgreement
}
