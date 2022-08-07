package main

import (
	"time"
)

type Qualification struct {
	// representation of Qualification class
	// contains information about user qualification
	// in specific sport and category
	id                int
	sport             Sport
	qualificationType int
	category          int
	document          []byte
	IROA              bool
	approved          bool
	dismissed         bool
	dismissReason     string
	approvedDate      time.Time
	meta              Meta
}

type QualificationInterface interface {
	AvailableBy(id int)
	For(sport Sport)
	Qualification(qualificationType int)
	Category(category int)
	Document(document []byte)
	IROA(IROA bool)
	Approved(approved bool)
	Dismissed(dismissed bool)
	Reason(dismissReason string)
	ApproveDate(approvedDate time.Time)
	WithMeta(meta Meta)
}
