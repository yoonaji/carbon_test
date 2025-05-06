package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoonaji/carbon_test/models"
	"gorm.io/gorm"
)

type TransactionController struct {
	DB *gorm.DB
}

func NewTransactionController(DB *gorm.DB) TransactionController {
	return TransactionController{DB}
}

func (pc *TransactionController) CreateTransaction(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateTransactionRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "잘못된 요청 형식입니다",
			"error":   err.Error(),
		})
		return
	}

	newTransaction := models.TransactionModel{
		TransactionID:     uuid.New().String(), // UUID 생성
		TransactionType:   payload.TransactionType,
		BankAccountID:     payload.BankAccountID,
		BankAccountNumber: payload.BankAccountNumber,
		BankCode:          payload.BankCode,
		Amount:            payload.Amount,
		TransactionDate:   payload.TransactionDate,
		TransactionName:   payload.TransactionName,
		Category:          "",
		CarbonScore:       0.0,
		UserID:            payload.UserID, //currentUser.ID.String(),
	}

	result := pc.DB.Create(&newTransaction)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "DB 저장 실패",
			"error":   result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"message": "거래내역 생성 성공",
		"error":   "",
	})
}

func (pc *TransactionController) UpdateCategory(ctx *gin.Context) {
	TransactionId := ctx.Param("transaction_id")
	//currentUser := ctx.MustGet("currentUser").(models.User)

	var payload struct {
		Category string `json:"category"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "잘못된 요청 형식입니다.",
			"error":   err.Error(),
		})
		return
	}
	var updatedTransaction models.TransactionModel
	result := pc.DB.First(&updatedTransaction, "transaction_id = ?", TransactionId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "거래내역을 찾을 수 없습니다.",
			"error":   result.Error.Error(),
		})
		return
	}

	updatedTransaction.Category = payload.Category
	pc.DB.Save(&updatedTransaction)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "카테고리 수정 완료",
		"error":   "",
	})
}

func (pc *TransactionController) UpdateCarbonScore(ctx *gin.Context) {
	transactionId := ctx.Param("transaction_id")

	var payload struct {
		CarbonScore float64 `json:"carbon_score"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "잘못된 요청 형식입니다.",
			"error":   err.Error(),
		})
		return
	}

	var transaction models.TransactionModel
	result := pc.DB.First(&transaction, "transaction_id = ?", transactionId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "거래내역을 찾을 수 없습니다.",
			"error":   result.Error.Error(),
		})
		return
	}

	transaction.CarbonScore = payload.CarbonScore
	if err := pc.DB.Save(&transaction).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "탄소 점수 업데이트 실패",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "탄소 점수 수정 완료",
		"error":   "",
	})
}

func (pc *TransactionController) FindTransactionById(ctx *gin.Context) {
	TransactionId := ctx.Query("transaction_id")

	var Transaction models.TransactionModel
	result := pc.DB.First(&Transaction, "transaction_id = ?", TransactionId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "거래내역이 존재하지 않습니다.",
			"error":   result.Error.Error(), // 수정: result.Error 사용
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "거래내역 확인",
		"error":   "",
	}) // 수정: post 대신 Transaction 사용
}

func (pc *TransactionController) FindTransactions(ctx *gin.Context) {
	UserId := ctx.Query("user_id")
	var transactions []models.TransactionModel
	result := pc.DB.Where("user_id = ?", UserId).Find(&transactions) //find를 해줘야 전체에서 조회해줌

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "서버 오류",
			"error":   result.Error.Error(),
		})
		return
	}

	if len(transactions) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "거래내역이 존재하지 않습니다.",
			"error":   "",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"result": transactions,
	})

}

func (pc *TransactionController) DeleteTransaction(ctx *gin.Context) {
	TransactionId := ctx.Param("transaction_id")

	result := pc.DB.Delete(&models.TransactionModel{}, "transaction_id = ?", TransactionId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "거래내역을 찾을 수 없습니다",
			"error":   "",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "거래내역 확인",
		"error":   "",
	}) // 수정: post 대신 Transaction 사용
}
