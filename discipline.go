package main

type Division struct {
	// represents Division for IPSC
	// contains information about
	// weapons and applied flags
	id          int
	content     []Content
	image       []byte
	canBeMinor  bool
	canBeMajor  bool
	customStyle bool
	weapon      Weapon
	meta        Meta
}

type DivisionInterface interface {
	AvailableBy(id int)
	WithContent(content Content)
	WithImage(image []byte)
	Minor(canBeMinor bool)
	Major(canBeMajor bool)
	CustomStyle(customStyle bool)
	Weapon(weapon Weapon)
	WithMeta(meta Meta)
}

type Discipline struct {
	//represents Discipline for IPSC
	// contains information about
	// division and competition type
	id                   int
	sport                Sport
	division             Division
	competitionType      int
	code                 string
	active               bool
	standardSpeedCourses bool
	meta                 Meta
}

type DisciplineInterface interface {
	AvailableBy(id int)
	For(sport Sport)
	In(division Division)
	Competition(competitionType int)
	Code(code string)
	SSC(standardSpeedCourses bool)
	Active(active bool)
	WithMeta(meta Meta)
}
