package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yoonaji/carbon_test/controllers"
	//"github.com/yoonaji/carbon/middleware"
)

type AuthRouteController struct {
	AuthController controllers.AuthController
}

func NewRouteAuthController(AuthController controllers.AuthController) AuthRouteController {
	return AuthRouteController{AuthController}
}

func (pc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/signup", pc.AuthController.Signup)
		auth.POST("/login", pc.AuthController.Login)
		auth.POST("/logout", pc.AuthController.Logout)
		auth.POST("/refresh", pc.AuthController.Refresh)
	}
}
