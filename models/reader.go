package models

import (
	"time"

	"gorm.io/gorm"
)

type Reader struct {
	gorm.Model
	Name           string    `json:"name" binding:"required" gorm:"not null"`
	Email          string    `json:"email" binding:"required,email" gorm:"not null,uniqueIndex"`
	Phone          string    `json:"phone" binding:"required" gorm:"not null, uniqueIndex"`
	MembershipDate time.Time `json:"membership_date" gorm:"autoCreateTime"`
	Borrowings []Borrowing	 
}

type CreateReader struct {
	Name  string `json:"name" binding:"required" gorm:"not null"`
	Email string `json:"email" binding:"required,email" gorm:"not null,uniqueIndex"`
	Phone string `json:"phone" binding:"required" gorm:"not null, uniqueIndex"`
}
