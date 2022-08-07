package main

type Offer struct {
	id          int
	destination int
	content     []Content
	meta        Meta
}

type OfferInterface interface {
	AvailableBy(id int)
	Destination(destination int) // 1 - пользак, 2 - организатор
	WithContent(content []Content)
	WithMeta(meta Meta)
}

type OfferBuilder struct {
	offer *Offer
}

func (b *OfferBuilder) AvailableBy(id int) *OfferIdBuilder {
	// set id for offer
	b.offer.id = id
	return &OfferIdBuilder{*b}
}

type OfferIdBuilder struct {
	// offer id builder
	OfferBuilder
}

func (b *OfferBuilder) Destination(destination int) *OfferDestinationBuilder {
	// set destination
	b.offer.destination = destination
	return &OfferDestinationBuilder{*b}
}

type OfferDestinationBuilder struct {
	// destination builder
	OfferBuilder
}

func (b *OfferBuilder) WithContent(content []Content) *OfferContentBuilder {
	// set content array with titles and descriptions
	for i := range content {
		b.offer.content = append(b.offer.content, content[i])
	}
	return &OfferContentBuilder{*b}
}

type OfferContentBuilder struct {
	// content builder
	OfferBuilder
}

func (b *OfferBuilder) WithMeta(meta Meta) *OfferMetaBuilder {
	// set meta for offer
	b.offer.meta = meta
	return &OfferMetaBuilder{*b}
}

type OfferMetaBuilder struct {
	// meta builder for offer
	OfferBuilder
}

func (b *OfferBuilder) Build() *Offer {
	// offer builder
	return b.offer
}
