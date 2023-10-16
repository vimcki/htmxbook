package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vimcki/htmxbook/internal/model"
	contactsRepo "github.com/vimcki/htmxbook/internal/model/repo"
)

type contactsQuery struct {
	Query string `form:"q"`
}

type contactsTemplateData struct {
	Query    string
	Contacts []model.Contact
}

type newTemplateData struct {
	First  string
	Last   string
	Phone  string
	Email  string
	Errors map[string]string
}

func main() {
	r := gin.Default()

	repo := contactsRepo.New()

	contactsTemplate := template.Must(
		template.ParseFiles("templates/layout.html", "templates/contacts.html"),
	)
	newTemplate := template.Must(template.ParseFiles("templates/layout.html", "templates/new.html"))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/contacts")
	})

	r.GET("/contacts", func(c *gin.Context) {
		query := contactsQuery{}
		if err := c.ShouldBindQuery(&query); err != nil {
			panic(err)
		}

		var contacts []model.Contact
		if query.Query == "" {
			contacts = repo.All()
		} else {
			contacts = repo.Search(query.Query)
		}

		err := contactsTemplate.Execute(c.Writer, contactsTemplateData{
			Query:    query.Query,
			Contacts: contacts,
		})
		if err != nil {
			panic(err)
		}
	})

	r.GET("/contacts/new", func(c *gin.Context) {
		newTemplate.Execute(c.Writer, newTemplateData{})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
