package models

import "time"

type Image struct {
	location string
	user User
	createdAt time.Time
}
