package models

import (
	"time"

	"gorm.io/gorm"
)

type Borrowing struct {
	gorm.Model
	BookID uint `json:"book_id" binding:"required"`
	Book   Book `json:"book"`

	ReaderId   uint       `json:"reader_id" binding:"required"`
	Reader     Reader     `json:"-"`
	BorrowedAt time.Time  `json:"borrowed_at" gorm:"autoCreateTime"`
	DueDate    time.Time  `json:"due_date" gorm:"not null"`
	ReturnedAt *time.Time `json:"returned_at"`
	Status     string     `json:"status" binding:"required, oneof= active returned overdue" gorm:"not null; check:(status IN ('active', 'returned', 'overdue'))"`
}

type CreateBorrowing struct {
	BookID   uint `json:"book_id" binding:"required"`
	ReaderId uint `json:"reader_id" binding:"required"`
}
