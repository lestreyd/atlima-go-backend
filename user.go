package main

import (
	"time"
)

type User struct {
	// user must be a valid set of attributes
	// with role and other properties as phone
	// and privacy configuration
	id       int
	slug     string
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

type UserInterface interface {
	AvailableBy()
	Id(id int)
	Slug(slug string)
	WithContacts(contacts Contacts)
	Set()
	Sex(sex int)
	Privacy(privacy PrivacyConfiguration)
	Qualifications(qualification []Qualification)
	StrongHand(strongHand int)
	BirthDate(birthDate time.Time)
	LastLogin(lastLogin time.Time)
	Image(image []byte)
	Role(role int)
	Active(active bool)
	Password(password string)
	AuthToken(authToken string)
	RefreshToken(refreshToken string)
	DeviceId(deviceId string)
	IpAddress(ipAddress string)
	Location()
	InCountry(country Country)
	InRegion(region Region)
	InCity(city City)
	WithMeta(meta Meta)
}
