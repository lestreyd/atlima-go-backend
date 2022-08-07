package main

import (
	"time"
)

type FrontendLog struct {
	// grouped by hash error log for
	// frontend with hash and user
	// information
	id         int
	hash       string
	user       User
	message    string
	stackTrace string
	build      int
	logDate    time.Time
	errorCode  string
	counts     int
	meta       Meta
}

type FrontendLogInterface interface {
	AvailableBy(id int)
	WithHash(hash string)
	For(user User)
	Build(build int)
	Message(message string)
	Trace(stackTrace string)
	Logged(logDate time.Time)
	ErrorCode(errorCode string)
	Counts(int)
	WithMeta(meta Meta)
}
