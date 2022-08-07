package main

type Guncheck struct {
	// guncheck fix result for slot
	// in specific event
	id          int
	discipline  Discipline
	category    int
	powerFactor int
	strongHand  int
	image       []byte
	referee     RefereeSlot
	meta        Meta
}

type GuncheckInterface interface {
	AvailableBy(id int)
	In(discipline Discipline)
	Category(category int)
	PowerFactor(powerFactor int)
	StrongHand(strongHand int)
	Reference(image []byte)
	WithMeta(meta Meta)
}

type GunCheckBuilder struct {
	gunCheck *Guncheck
}

func (b *GunCheckBuilder) AvailableBy(id int) *GunCheckIdBuilder {
	// create id
	b.gunCheck.id = id
	return &GunCheckIdBuilder{*b}
}

type GunCheckIdBuilder struct {
	// builder for Id
	GunCheckBuilder
}

func (b *GunCheckBuilder) In(discipline Discipline) *GunCheckDisciplineBuilder {
	// create id
	b.gunCheck.discipline = discipline
	return &GunCheckDisciplineBuilder{*b}
}

type GunCheckDisciplineBuilder struct {
	// builder for Id
	GunCheckBuilder
}

func (b *GunCheckBuilder) Category(category int) *GunCheckCategoryBuilder {
	// create id
	b.gunCheck.category = category
	return &GunCheckCategoryBuilder{*b}
}

type GunCheckCategoryBuilder struct {
	// builder for Id
	GunCheckBuilder
}

func (b *GunCheckBuilder) PowerFactor(powerFactor int) *GunCheckPowerFactorBuilder {
	// create id
	b.gunCheck.powerFactor = powerFactor
	return &GunCheckPowerFactorBuilder{*b}
}

type GunCheckPowerFactorBuilder struct {
	// builder for Id
	GunCheckBuilder
}

func (b *GunCheckBuilder) StrongHand(strongHand int) *GunCheckStrongHandBuilder {
	// create id
	b.gunCheck.strongHand = strongHand
	return &GunCheckStrongHandBuilder{*b}
}

type GunCheckStrongHandBuilder struct {
	// builder for Id
	GunCheckBuilder
}

func (b *GunCheckBuilder) Reference(image []byte) *GunCheckImageBuilder {
	// create id
	b.gunCheck.image = image
	return &GunCheckImageBuilder{*b}
}

type GunCheckImageBuilder struct {
	// builder for Id
	GunCheckBuilder
}

func (b *GunCheckBuilder) WithMeta(meta Meta) *GunCheckMetaBuilder {
	// create id
	b.gunCheck.meta = meta
	return &GunCheckMetaBuilder{*b}
}

type GunCheckMetaBuilder struct {
	// builder for Id
	GunCheckBuilder
}

func (b *GunCheckBuilder) Build() *Guncheck {
	// Slot object
	return b.gunCheck
}
