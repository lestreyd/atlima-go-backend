package main

import "time"

type EVSK struct {
	// contains information about
	// evsk event status
	name           string
	regionalStatus int
	meta           Meta
}

type EVSKInterface interface {
	Name(name string)
	Status(regionalStatus int)
	WithMeta(meta Meta)
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

type PriceOptionInterface interface {
	Name(name string)
	StaticPrice(staticPrice int)
	Price(price int)
	WithCurrency(currency Currency)
	WithMeta(meta Meta)
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

type EventControlFlagsInterface interface {
	SetBanned(banned bool)
	SetBannedModeration(bannedModeration bool)
	SetApproved(approved bool)
	SetModerated(moderated bool)
	SetHasResults(hasResults bool)
	SetCompleted(completed bool)
	WithMeta(meta Meta)
}

type EventAdministration struct {
	// information about
	// specific administration
	// of event (with referees)
	id             int
	referees       []RefereeSlot
	director       User
	administrators []User
	organizer      Organizer
	meta           Meta
}

type EventAdministrationInterface interface {
	AvailableBy(id int)
	Set()
	Referees(referees []RefereeSlot)
	Director(director User)
	Administrators(administrators []User)
	Organizers(organizer Organizer)
	WithMeta(meta Meta)
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

type EventPropertyInterface interface {
	AvailableBy(id int)
	For(sport Sport)
	Level(level int)
	SquadsAmount(squads int)
	ShootersAmount(shooters int)
	WithPrematch(prematch bool)
	NumberInCalendarPlan(numberInCalendarPlan string)
	WithMeta(meta Meta)
}

type Event struct {
	// provide information about event
	// with administration and participants
	// result information (slots with squads)
	id                   int
	slug                 string
	site                 string
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
	contacts             Contacts
	firstCalculationDate time.Time
	lastCalculationDate  time.Time
	meta                 Meta
}

type EventInterface interface {
	AvailableBy()
	Id(id int)
	Slug(slug string)
	Site(site string)
	For(sport Sport)
	Format(evsk EVSK)
	Where(location Location)
	When(startDate, endDate time.Time)
	AdministratedBy(administration EventAdministration)
	WithContacts(contacts Contacts)
	InStatus(status int)
	Courses(courses []Course)
	Properties(properties EventProperty)
	Price(price PriceOption)
	ControlFlags(flags EventControlFlags)
	EVSK(evsk EVSK)
	AddSlots(slots []Slot)
	AddSquads(squads []Squad)
	SetFirstCalculationDate(firstCalculationDate time.Time)
	SetLastCalculationDate(lastCalculationDate time.Time)
	SetInterestedIn(users []User)
	WithMeta(meta Meta)
}
