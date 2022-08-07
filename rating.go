package main

type RatingInformation struct {
	// contains all initial and calculated
	// values from this event
	id             int
	initialRating  int
	deviation      float32
	handicap       float32
	performance    float32
	ratingIncrease int
	meta           Meta
}

type RatingInformationInterface interface {
	AvailableBy(id int)
	IR(initialRating int)
	Deviation(deviation float32)
	Handicap(handicap float32)
	Performance(performance float32)
	Increase(ratingIncrease int)
	WithMeta(meta Meta)
}
