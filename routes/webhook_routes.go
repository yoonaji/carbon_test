package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yoonaji/carbon_test/controllers"
)

type WebhookRouteController struct {
	WebhookController controllers.WebhookController
}

func NewWebhookRouteController(webhookController controllers.WebhookController) WebhookRouteController {
	return WebhookRouteController{webhookController}
}

func (wc *WebhookRouteController) WebhookRoute(rg *gin.RouterGroup) {
	rg.POST("/webhook/payaction", wc.WebhookController.ReceiveWebhook)
}
