package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Password  string `json:"password" validate:"required,min=6"`
	Email     string `json:"email" validate:"required, email" gorm:"unique"`
	Phone     string `json:"phone" validate:"required"`
	// Token        *string `json:"token"`
	// RefreshToken *string `json:"refreshToken"`
}
