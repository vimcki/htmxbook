package handlers

import "github.com/gin-gonic/gin"

func AddArchiveHandlers(r *gin.Engine) {
	r.GET("/contacts/archive", func(c *gin.Context) {
	})
}
