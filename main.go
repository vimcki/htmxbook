package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	Flashes  []interface{}
}

type newTemplateData struct {
	First  string `form:"first_name"`
	Last   string `form:"last_name"`
	Phone  string `form:"phone"`
	Email  string `form:"email"`
	Errors map[string]string
}

type editTemplateData struct {
	ID     int
	First  string `form:"first_name"`
	Last   string `form:"last_name"`
	Phone  string `form:"phone"`
	Email  string `form:"email"`
	Errors map[string]string
}

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secretss"))

	r.Use(sessions.Sessions("user", store))

	repo := contactsRepo.New()

	contactsTemplate := template.Must(
		template.ParseFiles("templates/layout.html", "templates/contacts.html"),
	)
	newTemplate := template.Must(template.ParseFiles("templates/layout.html", "templates/new.html"))
	showTemplate := template.Must(
		template.ParseFiles("templates/layout.html", "templates/show.html"),
	)
	editTemplate := template.Must(
		template.ParseFiles("templates/layout.html", "templates/edit.html"),
	)

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

		session := sessions.Default(c)
		flashes := session.Flashes()
		session.Save()

		err := contactsTemplate.Execute(c.Writer, contactsTemplateData{
			Query:    query.Query,
			Contacts: contacts,
			Flashes:  flashes,
		})
		if err != nil {
			panic(err)
		}
	})

	r.GET("/contacts/new", func(c *gin.Context) {
		err := newTemplate.Execute(c.Writer, newTemplateData{})
		if err != nil {
			panic(err)
		}
	})

	r.POST("/contacts/new", func(c *gin.Context) {
		newContact := newTemplateData{}
		if err := c.ShouldBind(&newContact); err != nil {
			panic(err)
		}
		newContact.Errors = map[string]string{}
		if newContact.First == "" {
			newContact.Errors["first_name"] = "First name is required"
		}
		if newContact.Last == "" {
			newContact.Errors["last_name"] = "Last name is required"
		}
		if newContact.Email == "" {
			newContact.Errors["email"] = "Email is required"
		}

		if len(newContact.Errors) > 0 {
			err := newTemplate.Execute(c.Writer, newContact)
			if err != nil {
				panic(err)
			}
			return
		}

		repo.Save(model.Contact{
			First: newContact.First,
			Last:  newContact.Last,
			Phone: newContact.Phone,
			Email: newContact.Email,
		})

		session := sessions.Default(c)
		session.AddFlash("Contact " + newContact.First + " created!")
		session.Save()

		c.Redirect(http.StatusMovedPermanently, "/contacts")
	})

	r.GET("/contacts/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		contact, found := repo.Find(id)
		if !found {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = showTemplate.Execute(c.Writer, contact)
		if err != nil {
			panic(err)
		}
	})

	r.GET("/contacts/:id/edit", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		contact, found := repo.Find(id)
		if !found {
			c.AbortWithStatus(http.StatusNotFound)
		}

		modify := editTemplateData{
			ID:    contact.ID,
			First: contact.First,
			Last:  contact.Last,
			Phone: contact.Phone,
			Email: contact.Email,
		}

		err = editTemplate.Execute(c.Writer, modify)
		if err != nil {
			panic(err)
		}
	})

	r.POST("/contacts/:id/edit", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		contact, found := repo.Find(id)
		if !found {
			c.AbortWithStatus(http.StatusNotFound)
		}

		modify := editTemplateData{
			ID:    contact.ID,
			First: contact.First,
			Last:  contact.Last,
			Phone: contact.Phone,
			Email: contact.Email,
		}

		if err := c.ShouldBind(&modify); err != nil {
			panic(err)
		}

		modify.Errors = map[string]string{}
		if modify.First == "" {
			modify.Errors["first_name"] = "First name is required"
		}
		if modify.Last == "" {
			modify.Errors["last_name"] = "Last name is required"
		}
		if modify.Email == "" {
			modify.Errors["email"] = "Email is required"
		}
		if len(modify.Errors) > 0 {
			err := editTemplate.Execute(c.Writer, modify)
			if err != nil {
				panic(err)
			}
			return
		}

		contact.First = modify.First
		contact.Last = modify.Last
		contact.Phone = modify.Phone
		contact.Email = modify.Email

		repo.Update(contact.ID, contact)

		session := sessions.Default(c)
		session.AddFlash("Contact " + contact.First + " updated!")
		_ = session.Save()

		c.Redirect(http.StatusMovedPermanently, "/contacts")
	})

	r.DELETE("/contacts/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		contact, found := repo.Find(id)
		if !found {
			c.AbortWithStatus(http.StatusNotFound)
		}

		repo.Delete(contact.ID)

		session := sessions.Default(c)
		session.AddFlash("Contact " + contact.First + " deleted!")
		_ = session.Save()

		c.Redirect(http.StatusSeeOther, "/contacts")
	})

	r.GET("/contacts/validate/email", func(c *gin.Context) {
		email := c.Query("email")
		c.Writer.Header().Add("Content-Type", "text/html")
		if !strings.Contains(email, ".") {
			_, err := c.Writer.WriteString("Email must contain .")
			if err != nil {
				panic(err)
			}
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
