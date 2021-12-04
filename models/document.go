package models

type Document struct {
	Id      uint    `json:"id"`
	Title   string  `json:"title" binding:"required"`
	Content content `json:"content" gorm:"embedded" binding:"required"`
	Signee  string  `json:"signee" binding:"required"`
}

type content struct {
	Header string `json:"header" binding:"required"`
	Data   string `json:"data" binding:"required"`
}
