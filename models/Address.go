package models

type Address struct {
	line1 string
	line2 string
	line3 string
	city  string
	state string
	zip   string
	country string
	countryCode string
	isBilling bool
	user User
}
