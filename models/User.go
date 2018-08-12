package models

type User struct {
	firstName string
	age int
	username string
	lastName string
	rfpCreations rfpCreations
	Addresses [Address{}]Address
}

type rfpCreations map[rfp]rfp
d