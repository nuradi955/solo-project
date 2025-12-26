package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Bio  string `json:"bio"`
	Book []Book `json:"book"`
}

type CreateAuthor struct {
	Name string `json:"name" binding:"required"`
	Bio  string `json:"bio"`
}
