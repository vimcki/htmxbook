package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/vimcki/htmxbook/internal/archiver"
	"github.com/vimcki/htmxbook/internal/handlers"
	contactsRepo "github.com/vimcki/htmxbook/internal/model/repo"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secretss"))

	r.Use(sessions.Sessions("user", store))

	repo := contactsRepo.New()

	a := archiver.New()

	handlers.AddArchiveHandlers(r, a)
	handlers.AddMainApp(r, repo, a)

	r.Static("/static", "./static")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
