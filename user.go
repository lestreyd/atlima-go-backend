package main

import (
	"time"
)

type User struct {
	// user must be a valid set of attributes
	// with role and other properties as phone
	// and privacy configuration
	id       int
	names    Names
	contacts Contacts
	image    []byte
	role     string
	active   bool
	// salted password
	password string
	// oauth2 tokens
	authToken    string
	refreshToken string
	// specific device info
	deviceId  string
	ipAddress string
	// location info
	country Country
	region  Region
	city    City
	// privacy configuration
	// must be created by default
	// in user constructor
	sex            int
	privacy        PrivacyConfiguration
	qualifications []Qualification
	strongHand     int
	birthDate      time.Time
	lastLogin      time.Time
	meta           Meta
}
