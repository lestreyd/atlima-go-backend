package main

type Country struct {
	// representation of country with
	// name and code fields
	id       int
	titles   []Content
	alpha2   string
	alpha3   string
	iso      int
	location string
	precise  string
	weight   int
	meta     Meta
}

type CountryInterface interface {
	AvailableBy(id int)
	Titles(titles []Content)
	A2(alpha2 string)
	A3(alpha3 string)
	ISO(iso int)
	Location(location Location)
	Precise(precise string)
	Weight(weight int)
	WithMeta(meta Meta)
}

type Region struct {
	//representation of region with
	//name and code fields
	id      int
	country Country
	titles  []Content
	code    int
	meta    Meta
}

type RegionInterface interface {
	AvailableBy(id int)
	InCountry(country Country)
	Titles(titles []Content)
	Code(code int)
	WithMeta(meta Meta)
}

type City struct {
	// representation of city with
	// region reference and names
	id     int
	region Region
	titles []Content
	meta   Meta
}

type CityInterface interface {
	AvailableBy(id int)
	InRegion(region Region)
	Titles(titles []Content)
	WithMeta(meta Meta)
}

type Location struct {
	// represents a location info
	// about user
	id      int
	country Country
	region  Region
	city    City
	meta    Meta
}

type LocationInterface interface {
	AvailableBy(id int)
	InCountry(country Country)
	InRegion(region Region)
	InCity(city City)
	WithMeta(meta Meta)
}
