package main

type TargetType struct {
	// represent Target type
	// structure with allowedResults
	// as "A" or "ACD"
	paper          bool
	image          byte
	allowedResults []string
	content        Content
	meta           Meta
}

type TargetSet struct {
	// represent set of targets
	// with specific amount and
	// alphaCost for target
	targetType TargetType
	course     Course
	amount     int
	alphaCost  int
	meta       Meta
}
