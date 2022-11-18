package models

import (
	"time"

	"gorm.io/gorm"
)

type Editora struct {
	ID         uint `gorm:"primaryKey"`
	Nome       string
	Localidade string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
