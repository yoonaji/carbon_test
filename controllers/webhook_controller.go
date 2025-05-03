package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoonaji/carbon/initializers"
	"github.com/yoonaji/carbon/models"
)

type WebhookController struct{}

func NewWebhookController() WebhookController {
	return WebhookController{}
}

func (wc *WebhookController) ReceiveWebhook(ctx *gin.Context) {
	var rawData map[string]interface{}
	if err := ctx.ShouldBindJSON(&rawData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "잘못된 웹훅 형식입니다"})
		return
	}

	transactionDateStr, ok := rawData["transaction_date"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "transaction_date가 문자열이 아닙니다"})
		return
	}

	transactionDate, err := time.Parse(time.RFC3339, transactionDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "날짜 포맷 오류", "error": err.Error()})
		return
	}

	webhookTx := models.WebhookTransaction{
		TransactionType:   fmt.Sprintf("%v", rawData["transaction_type"]),
		BankAccountID:     fmt.Sprintf("%v", rawData["bank_account_id"]),
		BankAccountNumber: fmt.Sprintf("%v", rawData["bank_account_number"]),
		BankCode:          fmt.Sprintf("%v", rawData["bank_code"]),
		Amount:            uint(rawData["amount"].(float64)),
		TransactionDate:   transactionDate,
		TransactionName:   fmt.Sprintf("%v", rawData["transaction_name"]),
	}

	if err := initializers.DB.Create(&webhookTx).Error; err != nil {
		ctx.JSON(500, gin.H{"message": "웹훅 데이터 저장 실패", "error": err.Error()})
		return
	}

	// 직접 TransactionModel 생성
	newTransaction := models.TransactionModel{
		TransactionID:     uuid.New().String(),
		TransactionType:   webhookTx.TransactionType,
		BankAccountID:     webhookTx.BankAccountID,
		BankAccountNumber: webhookTx.BankAccountNumber,
		BankCode:          webhookTx.BankCode,
		Amount:            int(webhookTx.Amount),
		TransactionDate:   webhookTx.TransactionDate,
		TransactionName:   webhookTx.TransactionName,
		UserID:            "system",
	}

	if err := initializers.DB.Create(&newTransaction).Error; err != nil {
		ctx.JSON(500, gin.H{"message": "거래내역 생성 실패", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "웹훅 데이터와 거래내역이 성공적으로 저장되었습니다",
	})
}
