package main

import (
	"github.com/gin-gonic/gin"
	"github.com/perilainen/document-store/handlers"
	"github.com/perilainen/document-store/models"
)

func main() {
	r := gin.Default()

	r.GET("/documents", handlers.AllDocuments)
	r.POST("/document", handlers.CreateDocument)
	r.GET("/document/:id", handlers.GetDocument)
	r.PUT("/document/:id", handlers.PutDocument)
	models.Connect()
	r.Run()

}
