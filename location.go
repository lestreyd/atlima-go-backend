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

type Region struct {
	//representation of region with
	//name and code fields
	id      int
	country Country
	titles  []Content
	code    int
	meta    Meta
}

type City struct {
	// representation of city with
	// region reference and names
	id     int
	region Region
	titles []Content
	meta   Meta
}

type Location struct {
	// represents a location info
	// about user
	country Country
	region  Region
	city    City
	meta    Meta
}
