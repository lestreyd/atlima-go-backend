package main

type Weapon struct {
	// Weapon object for Division system
	// in IPSC
	id            int
	image         []byte
	abbreviations []Content
	meta          Meta
}

type WeaponInterface interface {
	AvailableBy(id int)
	Image(image []byte)
	Abbreviations(abbreviations []Content)
	WithMeta(meta Meta)
}
