package main

type Weapon struct {
	// Weapon object for Division system
	// in IPSC
	id            int
	image         []byte
	abbreviations []string
	meta          Meta
}

type WeaponInterface interface {
	AvailableBy(id int)
	Image(image []byte)
	Abbreviations(abbreviations []string)
	WithMeta(meta Meta)
}

type WeaponBuilder struct {
	weapon *Weapon
}

func (b *WeaponBuilder) AvailableBy(id int) *WeaponIdBuilder {
	// build availability by id
	b.weapon.id = id
	return &WeaponIdBuilder{*b}
}

type WeaponIdBuilder struct {
	WeaponBuilder
}

func (b *WeaponBuilder) Image(image []byte) *WeaponImageBuilder {
	// build meta for user
	b.weapon.image = image
	return &WeaponImageBuilder{*b}
}

type WeaponImageBuilder struct {
	WeaponBuilder
}

func (b *WeaponBuilder) Abbreviations(abbreviations []string) *WeaponAbbreviationsBuilder {
	// build abbreviations list
	for i := range abbreviations {
		b.weapon.abbreviations = append(b.weapon.abbreviations, abbreviations[i])
	}
	return &WeaponAbbreviationsBuilder{*b}
}

type WeaponAbbreviationsBuilder struct {
	WeaponBuilder
}

func (b *WeaponBuilder) Meta(meta Meta) *WeaponMetaBuilder {
	// build meta for user
	b.weapon.meta = meta
	return &WeaponMetaBuilder{*b}
}

type WeaponMetaBuilder struct {
	WeaponBuilder
}

func (b *WeaponBuilder) Build() *Weapon {
	// WeaponBuilder object
	return b.weapon
}
