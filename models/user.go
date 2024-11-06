package models

import (
	"github.com/google/uuid"
)

type Users struct {
	Id       uuid.UUID     `gorm:"type:uuid;primaryKey"`
	Nickname string        `gorm:"type:text;unique;not null;unique"`
	Password string        `gorm:"type:text;not null"`
	Tests    []TestToUsers `gorm:"foreignKey:user_id"`
}

type UpdatePassword struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type Authorize struct {
	ID uuid.UUID 
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
