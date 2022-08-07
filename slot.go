package main

import "time"

type PromoCode struct {
	// represent PromoCode structure
	// with titles and discount
	// properties
	id         int
	title      string
	discount   float32
	limit      int
	applied    int
	slots      []Slot
	finishDate time.Time
	meta       Meta
}

type PromoCodeInterface interface {
	AvailableBy(id int)
	Title(title string)
	Discount(discount float32)
	Limit(limit int)
	Finish(finishDate time.Time)
	WithMeta(meta Meta)
}

type PromoCodeBuilder struct {
	promoCode *PromoCode
}

func (b *PromoCodeBuilder) AvailableBy(id int) *PromoCodeIdBuilder {
	// add id for PromoCode
	b.promoCode.id = id
	return &PromoCodeIdBuilder{*b}
}

type PromoCodeIdBuilder struct {
	// builder for id
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Title(title string) *PromoCodeTitleBuilder {
	// add title
	b.promoCode.title = title
	return &PromoCodeTitleBuilder{*b}
}

type PromoCodeTitleBuilder struct {
	// builder for title
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Discount(discount float32) *PromoCodeDiscountBuilder {
	// add discount to PromoCode
	b.promoCode.discount = discount
	return &PromoCodeDiscountBuilder{*b}
}

type PromoCodeDiscountBuilder struct {
	// builder for discount field
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Limit(limit int) *PromoCodeLimitBuilder {
	// add limit
	b.promoCode.limit = limit
	return &PromoCodeLimitBuilder{*b}
}

type PromoCodeLimitBuilder struct {
	// builder for limit
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Finish(finishDate time.Time) *PromoCodeFinishBuilder {
	// add finishDate. If nil - not installed
	b.promoCode.finishDate = finishDate
	return &PromoCodeFinishBuilder{*b}
}

type PromoCodeFinishBuilder struct {
	// builder for promoCode
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) WithMeta(meta Meta) *PromoCodeMetaBuilder {
	// add meta to PromoCode
	b.promoCode.meta = meta
	return &PromoCodeMetaBuilder{*b}
}

type PromoCodeMetaBuilder struct {
	// builder for PromoCode Meta
	PromoCodeBuilder
}

func (b *PromoCodeBuilder) Build() *PromoCode {
	// Build PromoCode object
	return b.promoCode
}

type Squad struct {
	// represent Squad structure
	// with number and comment info
	// about squad
	id       int
	number   int
	comment  string
	blocked  bool
	password string
	meta     Meta
}

type SquadInterface interface {
	AvailableBy(id int)
	SetNumber(number string)
	SetComment(comment string)
	Blocked(blocked bool)
	WithPassword(password string)
	WithMeta(meta Meta)
}

type SquadBuilder struct {
	squad *Squad
}

type Team struct {
	// represents Team structure
	// with title and disciplines info
	id         int
	title      string
	discipline Discipline
	meta       Meta
}

type TeamInterface interface {
	AvailableBy(id int)
	Title(title string)
	Discipline(discipline Discipline)
	WithMeta(meta Meta)
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

type SlotResultInterface interface {
	AvailableBy(id int)
	For(course Course)
	Result(coursePoints int)
	Stage(stagePoints float32)
	HitFactor(hitFactor float32)
	WithMeta(meta Meta)
}

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
