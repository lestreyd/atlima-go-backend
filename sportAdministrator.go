package main

type SportAdministrator struct {
	// representation of sport admin user
	// with sks and referee committee flags
	id                        int
	user                      User
	country                   Country
	region                    Region
	sksMember                 bool
	sksPresident              bool
	refereeCommitteeMember    bool
	refereeCommitteePresident bool
	meta                      Meta
}

type SportAdministratorBuilder struct {
	// sport admin builder
	sportAdministrator *SportAdministrator
}

func (b *SportAdministratorBuilder) AvailableBy(id int) *SportAdministratorIdBuilder {
	// create id for availability
	b.sportAdministrator.id = id
	return &SportAdministratorIdBuilder{*b}
}

type SportAdministratorIdBuilder struct {
	// builder for id of language
	SportAdministratorBuilder
}

func (b *SportAdministratorBuilder) For(user User) *SportAdministratorUserBuilder {
	// create user for which we delegate administration
	b.sportAdministrator.user = user
	return &SportAdministratorUserBuilder{*b}
}

type SportAdministratorUserBuilder struct {
	// builder for user of sport administrator
	SportAdministratorBuilder
}

func (b *SportAdministratorBuilder) inCountry(country Country) *SportAdministratorCountryBuilder {
	// create coutry for which we delegate administration
	b.sportAdministrator.country = country
	return &SportAdministratorCountryBuilder{*b}
}

type SportAdministratorCountryBuilder struct {
	// builder for country of sport administrator
	SportAdministratorBuilder
}

func (b *SportAdministratorBuilder) inRegion(region Region) *SportAdministratorRegionBuilder {
	// create coutry for which we delegate administration
	b.sportAdministrator.region = region
	return &SportAdministratorRegionBuilder{*b}
}

type SportAdministratorRegionBuilder struct {
	// builder for country of sport administrator
	SportAdministratorBuilder
}

func (b *SportAdministratorBuilder) isSKSMember(sksMember bool) *SportAdministratorSksMemberBuilder {
	// create coutry for which we delegate administration
	b.sportAdministrator.sksMember = sksMember
	return &SportAdministratorSksMemberBuilder{*b}
}

type SportAdministratorSksMemberBuilder struct {
	// builder for country of sport administrator
	SportAdministratorBuilder
}

func (b *SportAdministratorBuilder) isSKSPresident(sksPresident bool) *SportAdministratorSksPresidentBuilder {
	// create coutry for which we delegate administration
	b.sportAdministrator.sksPresident = sksPresident
	return &SportAdministratorSksPresidentBuilder{*b}
}

type SportAdministratorSksPresidentBuilder struct {
	// builder for country of sport administrator
	SportAdministratorBuilder
}

func (b *SportAdministratorBuilder) isRefereeCommitteeMember(refereeCommitteeMember bool) *SportAdministratorRCMemberBuilder {
	// create coutry for which we delegate administration
	b.sportAdministrator.sksPresident = refereeCommitteeMember
	return &SportAdministratorRCMemberBuilder{*b}
}

type SportAdministratorRCMemberBuilder struct {
	// builder for country of sport administrator
	SportAdministratorBuilder
}

func (b *SportAdministratorBuilder) isRefereeCommitteePresident(refereeCommitteeMember bool) *SportAdministratorRCPresidentBuilder {
	// create coutry for which we delegate administration
	b.sportAdministrator.sksPresident = refereeCommitteeMember
	return &SportAdministratorRCPresidentBuilder{*b}
}

type SportAdministratorRCPresidentBuilder struct {
	// builder for country of sport administrator
	SportAdministratorBuilder
}

func (b *SportAdministratorBuilder) Build() *SportAdministrator {
	// provides language object for text blocks
	return b.sportAdministrator
}
