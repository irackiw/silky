package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	// ZROBIC TO POZNIEJ
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")

	// POPRAWNE UZYCIE REQUESTU - JAKO WZOR
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Stan początkowy",
		})
	})
	router.Run(os.Getenv("LISTEN"))
}
