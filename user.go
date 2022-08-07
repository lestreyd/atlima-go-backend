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

type UserBuilder struct {
	user *User
}

func (b *UserBuilder) AvailableBy() *UserAvailabilityBuilder {
	// set all availability parameters
	return &UserAvailabilityBuilder{*b}
}

type UserAvailabilityBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Id(id int) *UserIdBuilder {
	// set id for user
	b.user.id = id
	return &UserIdBuilder{*b}
}

type UserIdBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Slug(slug string) *UserSlugBuilder {
	// set slug for user
	b.user.slug = slug
	return &UserSlugBuilder{*b}
}

type UserSlugBuilder struct {
	UserBuilder
}

func (b *UserBuilder) WithContacts(contacts Contacts) *UserContactsBuilder {
	// set contact for User
	b.user.contacts = contacts
	return &UserContactsBuilder{*b}
}

type UserContactsBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Set() *UserPropertiesBuilder {
	// set user properties
	return &UserPropertiesBuilder{*b}
}

type UserPropertiesBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Sex(sex int) *UserSexBuilder {
	// set user sex
	b.user.sex = sex
	return &UserSexBuilder{*b}
}

type UserSexBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Privacy(privacy PrivacyConfiguration) *UserPrivacyBuilder {
	// set privacy configuration
	b.user.privacy = privacy
	return &UserPrivacyBuilder{*b}
}

type UserPrivacyBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Qualifications(qualifications []Qualification) *UserQualificationBuilder {
	// set qualifications for user
	for i := range qualifications {
		b.user.qualifications = append(b.user.qualifications, qualifications[i])
	}
	return &UserQualificationBuilder{*b}
}

type UserQualificationBuilder struct {
	UserBuilder
}

func (b *UserBuilder) StrongHand(strongHand int) *UserStrongHandBuilder {
	// set strong hand
	b.user.strongHand = strongHand
	return &UserStrongHandBuilder{*b}
}

type UserStrongHandBuilder struct {
	UserBuilder
}

func (b *UserBuilder) BirthDate(birthDate time.Time) *UserBirthDateBuilder {
	// set birthDate for user
	b.user.birthDate = birthDate
	return &UserBirthDateBuilder{*b}
}

type UserBirthDateBuilder struct {
	UserBuilder
}

func (b *UserBuilder) LastLogin(lastLogin time.Time) *UserLastLoginBuilder {
	// set last login datetime for user
	b.user.lastLogin = lastLogin
	return &UserLastLoginBuilder{*b}
}

type UserLastLoginBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Image(image []byte) *UserImageBuilder {
	// set image for user
	b.user.image = image
	return &UserImageBuilder{*b}
}

type UserImageBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Role(image []byte) *UserRoleBuilder {
	// set role for user
	b.user.image = image
	return &UserRoleBuilder{*b}
}

type UserRoleBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Active(active bool) *UserActivityBuilder {
	// set user active/non-active status
	b.user.active = active
	return &UserActivityBuilder{*b}
}

type UserActivityBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Password(password string) *UserPasswordBuilder {
	// set password for user (salted)
	b.user.password = password
	return &UserPasswordBuilder{*b}
}

type UserPasswordBuilder struct {
	UserBuilder
}

func (b *UserBuilder) AuthToken(authToken string) *UserAuthTokenBuilder {
	// set auth token for user
	b.user.authToken = authToken
	return &UserAuthTokenBuilder{*b}
}

type UserAuthTokenBuilder struct {
	UserBuilder
}

func (b *UserBuilder) RefreshToken(refreshToken string) *UserRefreshTokenBuilder {
	// set refresh token for oauth2
	b.user.refreshToken = refreshToken
	return &UserRefreshTokenBuilder{*b}
}

type UserRefreshTokenBuilder struct {
	UserBuilder
}

func (b *UserBuilder) DeviceId(deviceId string) *UserDeviceIdBuilder {
	// set user device id
	b.user.deviceId = deviceId
	return &UserDeviceIdBuilder{*b}
}

type UserDeviceIdBuilder struct {
	UserBuilder
}

func (b *UserBuilder) IpAddress(ipAddress string) *UserIpAddressBuilder {
	// set user ip address
	b.user.ipAddress = ipAddress
	return &UserIpAddressBuilder{*b}
}

type UserIpAddressBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Location() *UserLocationBuilder {
	// set user Location
	return &UserLocationBuilder{*b}
}

type UserLocationBuilder struct {
	UserBuilder
}

func (b *UserBuilder) InCountry(country Country) *UserCountryBuilder {
	// set user Country
	b.user.country = country
	return &UserCountryBuilder{*b}
}

type UserCountryBuilder struct {
	UserBuilder
}

func (b *UserBuilder) InRegion(region Region) *UserRegionBuilder {
	// set user Region
	b.user.region = region
	return &UserRegionBuilder{*b}
}

type UserRegionBuilder struct {
	UserBuilder
}

func (b *UserBuilder) InCity(city City) *UserCityBuilder {
	// set user City
	b.user.city = city
	return &UserCityBuilder{*b}
}

type UserCityBuilder struct {
	UserBuilder
}

func (b *UserBuilder) WithMeta(meta Meta) *UserMetaBuilder {
	// build meta for user
	b.user.meta = meta
	return &UserMetaBuilder{*b}
}

type UserMetaBuilder struct {
	UserBuilder
}

func (b *UserBuilder) Build() *User {
	// UserBuilder object
	return b.user
}
