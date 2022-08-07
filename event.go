package main

import "time"

type EVSK struct {
	// contains information about
	// evsk event status
	name           string
	regionalStatus int
	meta           Meta
}

type PriceOption struct {
	// contains information about price
	// behaviour inside event
	// startDate - finishDate
	name        string
	staticPrice bool
	price       int
	currency    Currency
	meta        Meta
}

type EventControlFlags struct {
	// all control flags for event
	banned           bool
	bannedModeration bool
	approved         bool
	moderated        bool
	hasResults       bool
	completed        bool
	meta             Meta
}

type EventAdministration struct {
	// information about
	// specific administration
	// of event (with referees)
	referees       []RefereeSlot
	director       User
	administrators []User
	organizer      Organizer
	meta           Meta
}

type EventProperty struct {
	// information about details
	// inside event - courses, disciplines
	// and other additional attributes
	id                   int
	sport                Sport
	level                int
	squads               int
	shooters             int
	prematch             bool
	numberInCalendarPlan string
	meta                 Meta
}

type Event struct {
	// provide information about event
	// with administration and participants
	// result information (slots with squads)
	id                   int
	sport                Sport
	format               int
	evsk                 EVSK
	slots                []Slot
	squads               []Squad
	location             Location
	administration       EventAdministration
	status               int
	priceInformation     PriceOption
	startDate            time.Time
	endDate              time.Time
	flags                EventControlFlags
	createdBy            User
	properties           EventProperty
	promoCodes           []PromoCode
	courses              []Course
	interestedIn         []User
	firstCalculationDate time.Time
	lastCalculationDate  time.Time
	meta                 Meta
}
