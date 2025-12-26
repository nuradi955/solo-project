package handlers

import (
	"fmt"
	data_base "solo-project/data-base"
	"solo-project/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func BorrowBook(bookID, readerID uint) (*models.Borrowing, error) {
	// 1. Проверить что книга существует
	var borrowing models.Borrowing
	var book models.Book
	if err := data_base.DB.First(&book, bookID).Error; err != nil {
		return nil, fmt.Errorf("книга с таким id не найдена")
	}
	// 2. Проверить что available
	if book.AvailableCopies < 1 {
		return nil, fmt.Errorf("книга не доступна")
	}

	// _
	// copies > 0
	// 3. Проверить что читатель существует
	var reader models.Reader
	if err := data_base.DB.First(&reader, readerID).Error; err != nil {
		return nil, fmt.Errorf("читатель с таким id не найден")
	}
	// 4. Проверить что у читателя нет активных просрочек
	var overdueCount int64
	if err := data_base.DB.Model(&borrowing).Where("reader_ID=? AND status=? AND due_date<?",
		readerID, "active", time.Now()).Count(&overdueCount).Error; err != nil {
		return nil, err
	}

	if overdueCount > 0 {
		return nil, fmt.Errorf("у читателя есть просрочка")
	}

	// 5. Создать Borrowing с due
	borrow := models.Borrowing{
		BookID:   bookID,
		ReaderId: readerID,
		DueDate:  time.Now().Add(14 * 24 * time.Hour),
		Status:   "active",
	}

	if err := data_base.DB.Create(&borrow).Error; err != nil {
		return nil, err
	}

	// date = сегодня + 14 дней
	// _
	// 6. Уменьшить available

	book.AvailableCopies -= 1

	if err := data_base.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return &borrow, nil

	// _
	// copies на 1
}

func BorrowBookHTTP(c *gin.Context) {
	var req models.CreateBorrowing

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	borrow, err := BorrowBook(req.BookID, req.ReaderId)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(201, borrow)

}

func GetBorrowings(c *gin.Context) {
	var borrowings []models.Borrowing

	status := c.Query("status")
	readerID := c.Query("reader_id")
	bookID := c.Query("book_id")

	query := data_base.DB.Model(&models.Borrowing{})

	if status != "" {
		query = query.Where("status=?", status)
	}

	if readerID != "" {
		id, err := strconv.Atoi(readerID)
		if err != nil {
			c.JSON(400, gin.H{"error": "id должен быть числом"})
			return
		}
		query = query.Where("reader_id=?", id)
	}

	if bookID != "" {
		id, err := strconv.Atoi(bookID)
		if err != nil {
			c.JSON(400, gin.H{"error": "id должен быть числом"})
			return
		}
		query = query.Where("book_id=?", id)
	}

	if err := query.Find(&borrowings).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if len(borrowings) == 0 {
		c.JSON(500, gin.H{"error": "выдачи не найдены"})
		return
	}
	c.JSON(200, borrowings)
}

func GetOverdueBorrowings() ([]models.Borrowing, error) {

	var borrowings []models.Borrowing
	if err := data_base.DB.Model(&models.Borrowing{}).
		Where("due_date<? AND status=?", time.Now(), "active").
		Update("status", "overdue").Error; err != nil {
		return nil, err
	}

	if err := data_base.DB.Where("status=? AND due_date<?", "overdue", time.Now()).Find(&borrowings).Error; err != nil {
		return nil, err
	}
	if len(borrowings) == 0 {
		return nil, fmt.Errorf("просроченных выдач нет")
	}
	return borrowings, nil
}

func GetBorrowingsOverdueHTTP(c *gin.Context) {
	res, err := GetOverdueBorrowings()
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
