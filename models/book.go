package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	AuthorID        uint   `json:"author_id" gorm:"not null"`
	Author          Author `json:"-"`
	Title           string `json:"title" gorm:"not null"`
	ISBN            string `json:"isbn" gorm:"not null"`
	TotalCopies     int    `json:"total_copies" gorm:"not null"`
	AvailableCopies int    `json:"available_copies" gorm:"not null"`
	Category        string `json:"category" gorm:"not null"`
}

type CreateBook struct {
	AuthorID    uint   `json:"author_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	TotalCopies int    `json:"total_copies" binding:"required"`
	Category    string `json:"category" binding:"required"`
}

type UpdateBook struct {
	AuthorID        *uint   `json:"author_id" binding:"required"`
	Title           *string `json:"title" binding:"required"`
	ISBN            *string `json:"isbn" binding:"required"`
	TotalCopies     *int    `json:"total_copies" binding:"required"`
	AvailableCopies *int    `json:"available_copies" binding:"required"`
	Category        *string `json:"category" binding:"required"`
}
