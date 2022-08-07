package main

import "time"

type PromoCode struct {
	// represent PromoCode structure
	// with titles and discount
	// properties
	id         int
	title      string
	discount   float32
	limit      int
	applied    int
	slots      []Slot
	finishDate time.Time
	meta       Meta
}

type PromoCodeInterface interface {
	AvailableBy(id int)
	Title(title string)
	Discount(discount float32)
	Limit(limit int)
	Finish(finishDate time.Time)
	WithMeta(meta Meta)
}

type PromoCodeBuilder struct {
	promoCode *PromoCode
}

func (b *PromoCodeBuilder) AvailableBy(id int) *PromoCodeIdBuilder {
	// add id for PromoCode
	b.promoCode.id = id
	return &PromoCodeIdBuilder{*b}
}

type PromoCodeIdBuilder struct {
	// builder for id
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Title(title string) *PromoCodeTitleBuilder {
	// add title
	b.promoCode.title = title
	return &PromoCodeTitleBuilder{*b}
}

type PromoCodeTitleBuilder struct {
	// builder for title
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Discount(discount float32) *PromoCodeDiscountBuilder {
	// add discount to PromoCode
	b.promoCode.discount = discount
	return &PromoCodeDiscountBuilder{*b}
}

type PromoCodeDiscountBuilder struct {
	// builder for discount field
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Limit(limit int) *PromoCodeLimitBuilder {
	// add limit
	b.promoCode.limit = limit
	return &PromoCodeLimitBuilder{*b}
}

type PromoCodeLimitBuilder struct {
	// builder for limit
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Finish(finishDate time.Time) *PromoCodeFinishBuilder {
	// add finishDate. If nil - not installed
	b.promoCode.finishDate = finishDate
	return &PromoCodeFinishBuilder{*b}
}

type PromoCodeFinishBuilder struct {
	// builder for promoCode
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) WithMeta(meta Meta) *PromoCodeMetaBuilder {
	// add meta to PromoCode
	b.promoCode.meta = meta
	return &PromoCodeMetaBuilder{*b}
}

type PromoCodeMetaBuilder struct {
	// builder for PromoCode Meta
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Build() *PromoCode {
	// Build PromoCode object
	return b.promoCode
}
