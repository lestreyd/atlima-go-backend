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
	weapons     Weapon
	meta        Meta
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
