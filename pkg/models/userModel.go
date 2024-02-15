package models

import "time"

type User struct {
	ID					uint		`gorm:"primaryKey" json:"id"`
	Name				string	`gorm:"name" json:"name"`
	Email				string	`gorm:"email" json:"email"`
	Password		string	`gorm:"password" json:"password"`
	CreatedAt 	time.Time
	UpdatedAt		time.Time
}