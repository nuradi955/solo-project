package handlers

import (
	data_base "solo-project/data-base"
	"solo-project/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBookHTTP(ctx *gin.Context) {
	var req models.CreateBook

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{
		AuthorID:        req.AuthorID,
		Title:           req.Title,
		ISBN:            req.ISBN,
		TotalCopies:     req.TotalCopies,
		AvailableCopies: req.TotalCopies,
		Category:        req.Category,
	}
	if err := data_base.DB.Create(&book).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, book)
}

// func GetBooks(ctx *gin.Context) {
// 	var books []models.Book

// 	if err := data_base.DB.Find(&books).Error; err != nil {
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(200, books)
// }

func GetBookByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "id должен быть числом"})
		return
	}
	var book models.Book
	if err := data_base.DB.First(&book, id).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var author models.Author

	if err := data_base.DB.Where("id=?", book.AuthorID).Find(&author).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"book": book,
		"author": author})
}

func UpdateBookHTTP(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "id должен быть числом"})
		return
	}

	var req models.UpdateBook
	var book models.Book

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	res := data_base.DB.Model(&book).Where("id = ?", id).Updates(req)
	if res.Error != nil {
		ctx.JSON(500, res.Error)
		return
	}
	if res.RowsAffected == 0 {
		ctx.JSON(400, gin.H{"error": "книга не найдена или не обновлена"})
		return
	}
	ctx.JSON(201, req)
}

func GetAvailableBooks(c *gin.Context) {

	var books []models.Book
	// var availableBooks []models.Book
	res := data_base.DB.Where("available_copies>?", 0).Find(&books)

	if res.Error != nil {
		c.JSON(500, gin.H{"error": res.Error})
		return
	}
	// for i, v := range books {
	// 	if v.AvailableCopies > 0 {
	// 		availableBooks = append(availableBooks, books[i])
	// 	}
	// }
	c.JSON(200, books)
}

func GetBooksQuery(c *gin.Context) {
	category := c.Query("category")
	authorId := c.Query("author_id")

	var books []models.Book

	query := data_base.DB.Model(&models.Book{})
	if category != "" {
		var count int64
		data_base.DB.Where("category=?", category).Find(&models.Book{}).Count(&count)

		if count == 0 {
			c.JSON(400, gin.H{"error": "книга с такой категорией не найдена"})
			return
		}
		query = query.Where("category=?", category)
	}

	if authorId != "" {
		id, err := strconv.Atoi(authorId)

		if err != nil {
			c.JSON(400, gin.H{"error": "id должен быть числом"})
			return
		}
		var author models.Author
		if err := data_base.DB.First(&author, id).Error; err != nil {
			c.JSON(400, gin.H{"error": "автор с таким id не найден"})
			return
		}
		query = query.Where("author_id=?", id)
	}
	if err := query.Find(&books).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(books) == 0 {
		c.JSON(404, gin.H{"error": "список книг пустой"})
		return
	}
	c.JSON(200, books)
}

