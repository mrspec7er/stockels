package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID int `json:"id" gorm:"primaryKey"`
	FullName string `json:"fullName"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	IsVerified bool `json:"isVerified"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Subscribtion []Subscribtion `json:"subscribtions"`
}