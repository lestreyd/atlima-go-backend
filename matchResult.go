package main

type MatchResult struct {
	// contains information about match result
	id          int
	percentage  float32
	stagePoints float32
	points      float32
	meta        Meta
}

type MatchResultInterface interface {
	AvailableBy(id int)
	Percentage(percentage float32)
	Stage(stagePoints float32)
	Points(points float32)
	WithMeta(meta Meta)
}
