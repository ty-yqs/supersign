package model

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Salt     string `gorm:"not null" json:"-"`
}

type UserStore interface {
	Query(username string) (*User, error)
	Update(oldUsername, newUsername, password string) error
}
