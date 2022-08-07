package main

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
	gunCheck                       Guncheck
	cancellation                   Cancellation
	meta                           Meta
}

type SlotInterface interface {
	AvailableBy(id int)
	For(user User)
	In(disciplines []Discipline)
	PromoCode(promoCode PromoCode)
	Squad(squad Squad)
	Category(category int)
	PowerFactor(powerFactor int)
	InTeam(team Team)
	Rating(rating RatingInformation)
	Result(result MatchResult)
	Active(active bool)
	Paid(paid bool)
	DontIncludeInRating(dontIncludeInRatingCalculation bool)
	Results(results []SlotResult)
	CoursesResults(courseResult []CourseResult)
	Guncheck(gunCheck Guncheck)
	Cancellation(cancellation Cancellation)
	WithMeta(meta Meta)
}
