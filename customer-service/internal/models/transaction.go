package models

import "time"

type Customer struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
