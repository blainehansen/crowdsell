package main

import (
	"time"
)


type Model struct {
	Id uint32 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type User struct {
	Model
	Name string `json:"name"`
	// https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address
	Email string `gorm:"type:varchar(254);unique;not null" json:"email"`
	Slug string `gorm:"unique;not null" json:"slug"`
	Password string `gorm:"not null"`
}

type Project struct {
	Model
	Name string `json:"name"`
}
