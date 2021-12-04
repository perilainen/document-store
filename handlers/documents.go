package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/perilainen/document-store/models"
)

// Get /documents
// Get all documents
func AllDocuments(c *gin.Context) {
	var docs []models.Document
	models.DB.Find(&docs)
	c.JSON(http.StatusOK, gin.H{"data": docs})

}

// POST /docuement
// Creata a new document
func CreateDocument(c *gin.Context) {
	var body models.Document
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	models.DB.Create(&body)
	c.JSON(http.StatusCreated, gin.H{"data": body})

}

// Get /document:id
// Get document by id
func GetDocument(c *gin.Context) {
	doc, err := getById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": doc})
}

// PUT /docuemnt:id
// Alters docuemnt with id
func PutDocument(c *gin.Context) {

	doc, err := getById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
	}
	var body models.Document
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.DB.Model(&doc).Updates(body)
	c.JSON(http.StatusOK, gin.H{"data": doc})
}

func getById(id string) (models.Document, error) {
	var doc models.Document

	if err := models.DB.Where("id = ?", id).First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil

}
