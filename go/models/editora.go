package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Editora struct {
	ID         uint   `gorm:"primaryKey"`
	Nome       string `gorm:"not null;default:null"`
	Localidade string `gorm:"not null;default:null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (e *Editora) Update(c *gin.Context, db *gorm.DB, n *Editora) {
	e.Nome = n.Nome
	e.Localidade = n.Localidade

	Update(c, db, e, nil)
}
