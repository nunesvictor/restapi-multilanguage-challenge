package models

import (
	"fmt"
	"gin-restapi/httputil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Livro struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Titulo    string    `json:",omitempty" gorm:"not null;default:null"`
	EditoraID int       `json:",omitempty"`
	Editora   Editora   `json:",omitempty" gorm:"constraint:OnDelete:SET NULL"`
	Autores   []Autor   `json:",omitempty" gorm:"many2many:livro_autores;constraint:OnDelete:CASCADE"`
	Generos   []Genero  `json:",omitempty" gorm:"many2many:livro_generos;constraint:OnDelete:CASCADE"`
	Edicao    int       `json:",omitempty"`
	Ano       int       `json:",omitempty"`
	Paginas   int       `json:",omitempty"`
	ISBN      string    `json:",omitempty"`
	CreatedAt time.Time `json:",omitempty"`
	UpdatedAt time.Time `json:",omitempty"`
}

func CreateLivro(c *gin.Context, db *gorm.DB, livro *Livro) error {
	if result := db.Create(&livro); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return fmt.Errorf(result.Error.Error())
	}

	if result := db.First(&livro); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return fmt.Errorf(result.Error.Error())
	}

	return nil
}

func (l *Livro) Update(c *gin.Context, db *gorm.DB, n *Livro) {
	l.Titulo = n.Titulo
	l.EditoraID = n.EditoraID
	l.Edicao = n.Edicao
	l.Ano = n.Ano
	l.Paginas = n.Paginas
	l.ISBN = n.ISBN

	postSave := func() {
		log.Println("============> UPDATE ASSOCIATION")
		log.Printf("%+v", n.Autores)
		log.Printf("%+v", n.Generos)

		db.Model(&l).Association("Autores").Replace(n.Autores)
		db.Model(&l).Association("Generos").Replace(n.Generos)
	}

	Update(c, db, l, postSave)
}
