package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yoonaji/carbon_test/controllers"
	"github.com/yoonaji/carbon_test/initializers"
	"github.com/yoonaji/carbon_test/routes"
)

var (
	server                     *gin.Engine
	TransactionController      controllers.TransactionController
	TransactionRouteController routes.TransactionRouteController
	WebhookController          controllers.WebhookController
	WebhookRouteController     routes.WebhookRouteController
	AuthController             controllers.AuthController
	AuthRouteController        routes.AuthRouteController
	UserController             controllers.UserController
	UserRouteController        routes.UserRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	TransactionController = controllers.NewTransactionController(initializers.DB)
	TransactionRouteController = routes.NewRouteTransactionController(TransactionController)
	WebhookController = controllers.NewWebhookController()
	WebhookRouteController = routes.NewWebhookRouteController(WebhookController)
	AuthController := controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewRouteAuthController(AuthController)
	UserController := controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Carbon API"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	TransactionRouteController.TransactionRoute(router) // 트랜잭션 라우트 연결
	WebhookRouteController.WebhookRoute(router)         // 웹훅 라우트 연결

	AuthRouteController.AuthRoute(router) // 인증 라우트 연결
	UserRouteController.UserRoute(router) // 유저 라우트 연결

	log.Fatal(server.Run(":" + config.ServerPort))
}
