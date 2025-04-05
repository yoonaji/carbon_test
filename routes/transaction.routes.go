package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
	"github.com/wpcodevo/golang-gorm-postgres/middleware"
)

type TransactionRouteController struct {
	TransactionController controllers.TransactionController
}

func NewRouteTransactionController(TransactionController controllers.TransactionController) TransactionRouteController {
	return TransactionRouteController{TransactionController}
}

func (pc *TransactionRouteController) TransactionRoute(rg *gin.RouterGroup) {

	router := rg.Group("transactions")
	router.Use(middleware.DeserializeUser())
	router.POST("/import", pc.TransactionController.CreateTransaction)
	router.GET("/list", pc.TransactionController.FindTransactions)
	router.GET("", pc.TransactionController.FindTransactionById)
	router.PUT("/:transaction_id/classify", pc.TransactionController.UpdateCarbonScore)
	router.PUT("/:transaction_id/carbonscore", pc.TransactionController.UpdateCategory)
	router.DELETE(":transaction_id", pc.TransactionController.DeleteTransaction)
}
