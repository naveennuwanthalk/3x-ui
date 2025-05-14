package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string    `json:"username" gorm:"unique;not null"`
	Password    string    `json:"password,omitempty" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique"`
	LastLogin   time.Time `json:"last_login"`
	LoginSecret string    `json:"login_secret,omitempty"`
	IsAdmin     bool      `json:"is_admin" gorm:"default:false"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	// Add any pre-save hooks here, like password hashing
	return nil
}