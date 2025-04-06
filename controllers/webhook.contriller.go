package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// 예: 필요한 필드 추출
	request := models.CreateTransactionRequest{
		TransactionType:   fmt.Sprintf("%v", rawData["transaction_type"]),
		BankAccountID:     fmt.Sprintf("%v", rawData["bank_account_id"]),
		BankAccountNumber: fmt.Sprintf("%v", rawData["bank_account_number"]),
		BankCode:          fmt.Sprintf("%v", rawData["bank_code"]),
		Amount:            int(rawData["amount"].(float64)), // float64 → int 변환
		TransactionDate:   fmt.Sprintf("%v", rawData["transaction_date"]),
		TransactionName:   fmt.Sprintf("%v", rawData["transaction_name"]),
		UserID:            fmt.Sprintf("%v", rawData["user_id"]),
	}

	// 👉 CreateTransaction에 직접 넘겨주기
	tc := controllers.TransactionController{DB: initializers.DB}
	tc.CreateTransactionFromWebhook(ctx, request)
}
