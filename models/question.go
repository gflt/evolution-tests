package models

import (
	"github.com/google/uuid"
)

type Question struct {
	Id      uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Name    string    `gorm:"type:text;not null"`
	TestsId uuid.UUID `gorm:"type:uuid;not null;index"`
	Answers []Answer  `gorm:"foreignKey:QuestionId"`
}
