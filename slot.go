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
	// Result contains MatchResult array
	// with overall result in Event
	Result(result MatchResult)
	Active(active bool)
	Paid(paid bool)
	DontIncludeInRating(dontIncludeInRatingCalculation bool)
	// ResultsByCourses contains SlotResult
	// with calculated stage points
	ResultsByCourses(results []SlotResult)
	// RAWCoursesResults contains CourseResult
	// Raw data from Frontend
	RAWCoursesResults(courseResult []CourseResult)
	GunCheck(gunCheck Guncheck)
	Cancellation(cancellation Cancellation)
	WithMeta(meta Meta)
}

type SlotBuilder struct {
	slot *Slot
}

func (b *SlotBuilder) AvailableBy(id int) *SlotAvailabilityBuilder {
	// create id
	b.slot.id = id
	return &SlotAvailabilityBuilder{*b}
}

type SlotAvailabilityBuilder struct {
	// builder for Id
	SlotBuilder
}

func (b *SlotBuilder) For(user User) *SlotUserBuilder {
	// User builder for Slot
	b.slot.user = user
	return &SlotUserBuilder{*b}
}

type SlotUserBuilder struct {
	// add User in Slot
	SlotBuilder
}

func (b *SlotBuilder) In(disciplines []Discipline) *SlotDisciplinesBuilder {
	// add disciplines
	for i := range disciplines {
		b.slot.disciplines = append(b.slot.disciplines, disciplines[i])
	}
	return &SlotDisciplinesBuilder{*b}
}

type SlotDisciplinesBuilder struct {
	// builder for disciplines in slot
	SlotBuilder
}

func (b *SlotBuilder) PromoCode(promoCode PromoCode) *SlotPromoCodeBuilder {
	// PromoCode builder
	b.slot.promoCode = promoCode
	return &SlotPromoCodeBuilder{*b}
}

type SlotPromoCodeBuilder struct {
	// promocode builder
	SlotBuilder
}

func (b *SlotBuilder) Squad(squad Squad) *SlotSquadBuilder {
	// add info about squad
	b.slot.squad = squad
	return &SlotSquadBuilder{*b}
}

type SlotSquadBuilder struct {
	// builder for Slot Squad
	SlotBuilder
}

func (b *SlotBuilder) Category(category int) *SlotCategoryBuilder {
	// set Category
	b.slot.category = category
	return &SlotCategoryBuilder{*b}
}

type SlotCategoryBuilder struct {
	// category builder
	SlotBuilder
}

func (b *SlotBuilder) PowerFactor(powerFactor int) *SlotPFBuilder {
	// set power factor
	b.slot.powerFactor = powerFactor
	return &SlotPFBuilder{*b}
}

type SlotPFBuilder struct {
	// builder for slot power factor
	SlotBuilder
}

func (b *SlotBuilder) InTeam(team Team) *SlotTeamBuilder {
	// add team
	b.slot.team = team
	return &SlotTeamBuilder{*b}
}

type SlotTeamBuilder struct {
	// builder for team
	SlotBuilder
}

func (b *SlotBuilder) Rating(rating RatingInformation) *SlotRatingBuilder {
	// slot rating information builder
	b.slot.rating = rating
	return &SlotRatingBuilder{*b}
}

type SlotRatingBuilder struct {
	// builder for slot rating
	SlotBuilder
}

func (b *SlotBuilder) Result(result MatchResult) *SlotResultArrayBuilder {
	// add result for this event
	b.slot.result = result
	return &SlotResultArrayBuilder{*b}
}

type SlotResultArrayBuilder struct {
	// builder for slot results
	SlotBuilder
}

func (b *SlotBuilder) Active(active bool) *SlotActivityBuilder {
	// set active flag
	b.slot.active = active
	return &SlotActivityBuilder{*b}
}

type SlotActivityBuilder struct {
	// active flag constructor
	SlotBuilder
}

func (b *SlotBuilder) Paid(paid bool) *SlotPaidBuilder {
	// paid flag
	b.slot.paid = paid
	return &SlotPaidBuilder{*b}
}

type SlotPaidBuilder struct {
	// builder for paid flag
	SlotBuilder
}

func (b *SlotBuilder) DontIncludeInRating(dontIncludeInRatingCalculation bool) *SlotExcludeBuilder {
	// set dontIncludeInRatingCalculation
	b.slot.dontIncludeInRatingCalculation = dontIncludeInRatingCalculation
	return &SlotExcludeBuilder{*b}
}

type SlotExcludeBuilder struct {
	// builder dont-include-in-rating-calculation
	// this flag is true when participant is
	// disqualified or not started
	SlotBuilder
}

func (b *SlotBuilder) ResultsByCourses(slotResult []SlotResult) *SlotRBCBuilder {
	// append results by courses to slot result
	for i := range slotResult {
		b.slot.results = append(b.slot.results, slotResult[i])
	}
	return &SlotRBCBuilder{*b}
}

type SlotRBCBuilder struct {
	// builder for result by courses
	SlotBuilder
}

func (b *SlotBuilder) RAWCoursesResults(courseResult []CourseResult) *SlotCourseResultBuilder {
	// append calculated course result
	for i := range courseResult {
		b.slot.courseResults = append(b.slot.courseResults, courseResult[i])
	}
	return &SlotCourseResultBuilder{*b}
}

type SlotCourseResultBuilder struct {
	// builder for course result
	SlotBuilder
}

func (b *SlotBuilder) GunCheck(gunCheck Guncheck) *SlotGunCheckBuilder {
	// create GunCheck
	b.slot.gunCheck = gunCheck
	return &SlotGunCheckBuilder{*b}
}

type SlotGunCheckBuilder struct {
	// SlotGunCheckBuilder - builder for slot GunCheck
	SlotBuilder
}

func (b *SlotBuilder) Cancellation(cancellation Cancellation) *SlotCancellationBuilder {
	// cancellation (DNS or DQ)
	b.slot.cancellation = cancellation
	return &SlotCancellationBuilder{*b}
}

type SlotCancellationBuilder struct {
	// builder for cancellation
	SlotBuilder
}

func (b *SlotBuilder) WithMeta(meta Meta) *SlotMetaBuilder {
	// create Meta object for Slot
	b.slot.meta = meta
	return &SlotMetaBuilder{*b}
}

type SlotMetaBuilder struct {
	// builder meta
	SlotBuilder
}

func (b *SlotBuilder) Build() *Slot {
	// Slot object
	return b.slot
}
