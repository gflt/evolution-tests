package models

import (
	"github.com/google/uuid"
)

type TestToUsers struct {
	UserId   uuid.UUID `gorm:"type:uuid;not null;index"`
	TestId   uuid.UUID `gorm:"type:uuid;not null;index"`
	User     Users     `gorm:"foreignKey:UserId"`
	Test     Tests     `gorm:"foreignKey:TestId"`
	IsPassed bool      `gorm:"default:false"`
}

type TestResult struct {
	TestId   uuid.UUID    `gorm:"type:uuid;not null;index"`
	Test     Result_Tests `gorm:"foreignKey:TestId"`
	IsPassed bool         `gorm:"default:false"`
}
