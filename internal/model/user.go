package model

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Fullname  string    `db:"fullname" json:"fullname"`
	Email     string    `db:"email" json:"email"`
	BirthDate time.Time `db:"birth_date" json:"birthDate"`
	Age       int       `db:"-" json:"age"`
}
