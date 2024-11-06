package models

import (
	"github.com/google/uuid"
)

type Tests struct {
	Id        uuid.UUID  `gorm:"type:uuid;primaryKey;not null"`
	Name      string     `gorm:"type:text;not null"`
	Questions []Question `gorm:"foreignKey:tests_id"`
}
 