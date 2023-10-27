package handlers

import (
	"fmt"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/vimcki/htmxbook/internal/archiver"
)

func AddArchiveHandlers(r *gin.Engine, a *archiver.Archiver) {
	template := template.Must(
		template.ParseFiles("templates/archive_base.html", "templates/archive.html"),
	)
	r.POST("/contacts/archive", func(c *gin.Context) {
		session := sessions.Default(c)
		archive := a.Get(session.ID())
		archive.Run()
		a := archiver.Archive{
			Status:   archive.Status,
			Progress: archive.Progress * 100,
		}
		err := template.Execute(c.Writer, a)
		if err != nil {
			panic(err)
		}
	})

	r.GET("/contacts/archive", func(c *gin.Context) {
		session := sessions.Default(c)
		fmt.Println(session.ID())
		archive := a.Get(session.ID())
		a := archiver.Archive{
			Status:   archive.Status,
			Progress: archive.Progress * 100,
		}
		err := template.Execute(c.Writer, a)
		if err != nil {
			panic(err)
		}
	})

	r.GET("/contacts/archive/file", func(c *gin.Context) {
		session := sessions.Default(c)
		archive := a.Get(session.ID())
		file := archive.ArchiveFile()
		c.Header("Content-Disposition", "attachment; filename=contacts.json")
		c.Data(200, "application/json", file)
	})

	r.DELETE("/contacts/archive", func(c *gin.Context) {
		session := sessions.Default(c)
		archive := a.Get(session.ID())
		archive.Reset()
		err := template.Execute(c.Writer, archive)
		if err != nil {
			panic(err)
		}
	})
}
