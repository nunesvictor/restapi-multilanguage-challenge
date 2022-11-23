package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Editora struct {
	ID         uint      `json:"id,omitempty" gorm:"primaryKey"`
	Nome       string    `json:",omitempty" gorm:"not null;default:null"`
	Localidade string    `json:",omitempty" gorm:"not null;default:null"`
	CreatedAt  time.Time `json:",omitempty"`
	UpdatedAt  time.Time `json:",omitempty"`
}

func (e *Editora) Update(c *gin.Context, db *gorm.DB, n *Editora) {
	e.Nome = n.Nome
	e.Localidade = n.Localidade

	Update(c, db, e, nil)
}
