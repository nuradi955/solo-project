package handlers

import (
	config "solo-project/data-base"
	data_base "solo-project/data-base"
	"solo-project/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAuthorHTTP(ctx *gin.Context) {
	req := models.CreateAuthor{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res := models.Author{
		Name: req.Name,
		Bio:  req.Bio,
	}
	result := config.DB.Create(&res)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(201, gin.H{"id": res.ID, "name": req.Name, "bio": req.Bio})
}

func GetAuthorsHTTP(ctx *gin.Context) {
	var allAuthors []models.Author

	res := config.DB.Find(&allAuthors)

	if res.Error != nil {
		ctx.JSON(500, gin.H{"error": res.Error})
		return
	}

	if len(allAuthors) == 0 {
		ctx.JSON(404, gin.H{"error": "список авторов пустой"})
		return
	}

	ctx.JSON(200, gin.H{"authors": allAuthors})
}

func GetAuthorByID(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "id должен быть числом"})
		return
	}
	var author models.Author
	res := config.DB.First(&author, id)
	if res.Error != nil {
		ctx.JSON(404, gin.H{"error": "автор не найден"})
		return
	}

	var books []models.Book

	if err := data_base.DB.Where("author_id=?", id).Find(&books).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"author": author, "books": books})
}
