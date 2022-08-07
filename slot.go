package main

import "time"

type PromoCode struct {
	// represent PromoCode structure
	// with titles and discount
	// properties
	title      string
	discount   string
	limit      int
	applied    int
	slots      []Slot
	finishDate time.Time
	createdAt  time.Time
	updatedAt  time.Time
	meta       Meta
}

type Squad struct {
	// represent Squad structure
	// with number and comment info
	// about squad
	number  int
	comment string
	blocked bool
	meta    Meta
}

type Team struct {
	// represents Team structure
	// with title and disciplines info
	title      string
	discipline Discipline
	meta       Meta
}

type SlotResult struct {
	// representation of Course result
	// for user in event
	id           int
	course       Course
	coursePoints int
	stagePoints  float32
	hitFactor    float32
	meta         Meta
}

type RatingInformation struct {
	// contains all initial and calculated
	// values from this event
	initialRating  int
	deviation      float32
	handicap       float32
	performance    float32
	ratingIncrease int
	meta           Meta
}

type MatchResult struct {
	// contains information about match result
	percentage  float32
	stagePoints float32
	points      float32
	meta        Meta
}

type Slot struct {
	// slot information about
	// event participation results
	id                             int
	user                           User
	disciplines                    []Discipline
	promoCode                      PromoCode
	squad                          Squad
	category                       int
	powerFactor                    int
	team                           Team
	rating                         RatingInformation
	result                         MatchResult
	active                         bool
	paid                           bool
	dontIncludeInRatingCalculation bool
	results                        []SlotResult
	courseResults                  []CourseResult
	guncheck                       Guncheck
	cancellation                   Cancellation
	meta                           Meta
}
