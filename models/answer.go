package models

import (
	"github.com/google/uuid"
)

type Answer struct {
	Id         uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Name       string    `gorm:"type:text;not null"`
	IsCorrect  bool      `gorm:"not null"`
	QuestionId uuid.UUID `gorm:"type:uuid;not null;index"`
}
