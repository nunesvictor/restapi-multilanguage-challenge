package models

import (
	"time"
)

type Editora struct {
	ID         uint `gorm:"primaryKey"`
	Nome       string
	Localidade string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
