package main

// import (
// 	"gopkg.in/src-d/go-kallax.v1"
// )

// type User struct {
// 	kallax.Model `table:"users"`
// 	kallax.Timestamps

// 	Id int64 `pk:"autoincr"`
// 	Slug string `unique:"true"`
// 	UrlSlug string `unique:"true"`

// 	Name *string
// 	Email string `unique:"true"`
// 	Password []byte

// 	ProfilePhotoSlug *string
// 	ForgotPasswordToken *string

// 	Projects []*Project `fk:"user_id"`
// }

// type Project struct {
// 	kallax.Model `table:"projects"`
// 	kallax.Timestamps

// 	Id int64 `pk:"autoincr"`
// 	Slug string `unique:"true"`
// 	UrlSlug string `unique:"true"`

// 	Name *string
// 	Description *string

// 	User User `fk:",inverse"`
// }
