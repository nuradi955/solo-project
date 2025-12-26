package handlers

import (
	data_base "solo-project/data-base"
	"solo-project/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReturnBook(borrowingID uint) error {
	// 1. Проверить что borrowing существует и status = "active"
	var borrowing models.Borrowing
	if err := data_base.DB.Where("status=? AND id =?", "active", borrowingID).First(&borrowing).Error; err != nil {
		return err
	}
	// 2. Установить returned_at = сейчас
	// 3. Установить status ="returned"
	if err := data_base.DB.Model(&borrowing).Updates(map[string]interface{}{"returned_at": time.Now(), "status": "returned"}).Error; err != nil {
		return err
	}
	// 4. Увеличить book.available_copies на 1

	if err := data_base.DB.Model(&models.Book{}).Where("id=?", borrowing.BookID).Update("available_copies",
		gorm.Expr("available_copies + 1")).Error; err != nil {
		return err
	}
	return nil
}

func BorrowingReturn(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = ReturnBook(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "книга успешно возврашена"})
}
