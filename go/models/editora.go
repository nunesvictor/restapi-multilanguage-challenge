package models

import (
	"gin-restapi/schemas"
	"time"
)

type Editora struct {
	ID         uint   `gorm:"primaryKey"`
	Nome       string `gorm:"not null;default:null"`
	Localidade string `gorm:"not null;default:null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (e *Editora) UpdateFromInput(s *schemas.EditoraInput) {
	e.Nome = s.Nome
	e.Localidade = s.Localidade
}
