package main

type RefereeGrade struct {
	// shows info about grade from
	// main referee for referees
	// in event
	refereeSlot RefereeSlot
	event       Event
	grade       int
	meta        Meta
}

type RefereeGradeInterface interface {
	For(refereeSlot RefereeSlot)
	On(event Event)
	Grade(grade int)
	WithMeta(meta Meta)
}

type Penalty struct {
	// represent penalty structure
	// with content
	content       []Content
	clause        string
	disciplines   []Discipline
	costInSeconds int
	amount        int
	active        bool
	meta          Meta
}

type PenaltyInterface interface {
	Clause(clause string)
	Content(content []Content)
	In(disciplines []Discipline)
	Cost(costInSeconds int)
	Amount(amount int)
	Active(active bool)
	WithMeta(meta Meta)
}

type Guncheck struct {
	// guncheck fix result for slot
	// in specific event
	discipline  Discipline
	category    int
	powerFactor int
	strongHand  int
	image       []byte
	referee     RefereeSlot
	meta        Meta
}

type GuncheckInterface interface {
	In(discipline Discipline)
	Category(category int)
	PowerFactor(powerFactor int)
	StrongHand(strongHand int)
	Reference(image []byte)
	WithMeta(meta Meta)
}

type CourseResult struct {
	// course fix result for slot
	// in specific event
	course Course
	A      int
	C      int
	D      int
	M      int
	NS     int
	T      float32
	image  []byte
	// bunch of penalties
	penalties []Penalty
	referee   RefereeSlot
	meta      Meta
}

type CourseResultInterface interface {
	For(course Course)
	Set()
	A(A int)
	C(C int)
	D(D int)
	M(M int)
	NS(NS int)
	T(T float32)
	Reference(image []byte)
	Penalties(penalties []Penalty)
	Referee(referee RefereeSlot)
	WithMeta(meta Meta)
}

type Cancellation struct {
	// disqualification from event
	course           Course
	cancellationType int    // DNS/DQ
	reason           string // Why?
	image            []byte
	referee          RefereeSlot
	meta             Meta
}

type CancellationInterface interface {
	For(course Course)
	What(cancellationType int)
	Why(reason string)
	Reference(image []byte)
	Referee(referee RefereeSlot)
	WithMeta(meta Meta)
}
