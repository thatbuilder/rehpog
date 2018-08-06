package models

type rfp struct {
	user       User
	title      string
	images     int
	priceLow   Price
	priceHighr Price
}