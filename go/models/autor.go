package models

import (
	"time"

	"gorm.io/gorm"
)

type Autor struct {
	ID        uint `gorm:"primaryKey"`
	Nome      string
	Sobrenome string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Autor) TableName() string {
	return "autores"
}
