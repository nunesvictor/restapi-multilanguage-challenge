package models

import (
	"gin-restapi/schemas"
	"time"
)

type Autor struct {
	ID        uint `gorm:"primaryKey"`
	Nome      string
	Sobrenome string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Autor) TableName() string {
	return "autores"
}

func (a *Autor) UpdateFromInput(s *schemas.AutorInput) {
	a.Nome = s.Nome
	a.Sobrenome = s.Sobrenome
}
