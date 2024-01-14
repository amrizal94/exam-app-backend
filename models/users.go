package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Code     uuid.UUID `gorm:"size:191;unique"`
	Name     string
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	RoleID   uint   `gorm:"default:2"`
	Role     Role
}

type Role struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"unique"`
}
