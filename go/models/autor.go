package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Autor struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Nome      string    `json:",omitempty" gorm:"not null;default:null"`
	Sobrenome string    `json:",omitempty" gorm:"not null;default:null"`
	CreatedAt time.Time `json:",omitempty"`
	UpdatedAt time.Time `json:",omitempty"`
}

func (Autor) TableName() string {
	return "autores"
}

func (a *Autor) Update(c *gin.Context, db *gorm.DB, n *Autor) {
	a.Nome = n.Nome
	a.Sobrenome = n.Sobrenome

	Update(c, db, a, nil)
}
