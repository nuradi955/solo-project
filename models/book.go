package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	AuthorID        uint   `json:"author_id"`
	Author          Author `json:"-"`
	Title           string `json:"title" binding:"required" gorm:"not null"`
	ISBN            string `json:"isbn" binding:"required" gorm:"not null"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
	Category        string `json:"category"`
}

type CreateBook struct {
	AuthorID    uint   `json:"author_id"`
	Title       string `json:"title" binding:"required" gorm:"not null"`
	ISBN        string `json:"isbn" binding:"required" gorm:"not null"`
	TotalCopies int    `json:"total_copies"`
	Category    string `json:"category"`
}

type UpdateBook struct {
	AuthorID        *uint   `json:"author_id"`
	Title           *string `json:"title" binding:"required" gorm:"not null"`
	ISBN            *string `json:"isbn" binding:"required" gorm:"not null"`
	TotalCopies     *int    `json:"total_copies"`
	AvailableCopies *int    `json:"available_copies"`
	Category        *string `json:"category"`
}
