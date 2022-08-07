package main

import (
	"time"
)

type Qualification struct {
	// representation of Qualification class
	// contains information about user qualification
	// in specific sport and category
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
