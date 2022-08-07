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

type Cancellation struct {
	// disqualification from event
	course           Course
	cancellationType int    // DNS/DQ
	reason           string // Why?
	image            byte
	referee          RefereeSlot
	meta             Meta
}
