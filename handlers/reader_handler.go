package handlers

import (
	data_base "solo-project/data-base"
	"solo-project/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateReader(c *gin.Context) {

	var req models.CreateReader

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	reader := models.Reader{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}
	res := data_base.DB.Create(&reader)

	if res.Error != nil {
		c.JSON(500, gin.H{"error": res.Error})
		return
	}
	c.JSON(201, reader)
}

func GetReaders(c *gin.Context) {
	var readers []models.Reader

	if err := data_base.DB.Find(&readers).Error; err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, readers)

}

func GetReaderByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "id должен быть числом"})
		return
	}
	var reader models.Reader
	if err := data_base.DB.Preload("Borrowings.Book").First(&reader, id).Error; err != nil {
		c.JSON(404, err.Error())
		return
	}
	var books []models.Book

	if err := data_base.DB.
		Joins("JOIN borrowings ON borrowings.book_id = books.id").
		Where("borrowings.reader_id = ?", id).
		Find(&books).Error; err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, books)

	c.JSON(200, gin.H{"reader": reader.Name, "books": books})
}

func GetRearedHistory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "id должен быть числом"})
		return
	}
	var history []models.Borrowing

	if err := data_base.DB.Where("reader_id=?", id).Find(&history).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, history)
}
