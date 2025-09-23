package routes

import (
	"go-stripe/handlers"
	"go-stripe/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	authHandler := &handlers.AuthHandler{DB: db}
	paymentHandler := &handlers.PaymentHandler{DB: db}

	// public routes
	r.GET("/", authHandler.ShowLogin)
	r.GET("/register", authHandler.ShowRegister)
	r.POST("/register", authHandler.Register)
	r.GET("/login", authHandler.ShowLogin)
	r.POST("/login", authHandler.Login)

	// preotected routes
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthRequired())
	{

		authGroup.GET("/dashboard", authHandler.Dashboard)
		authGroup.GET("/send", paymentHandler.ShowSendMoney)
		authGroup.POST("/create-checkout-session", paymentHandler.CreateCheckoutSession)
		authGroup.GET("/success", paymentHandler.HandleSuccess)
		authGroup.POST("/logout", authHandler.Logout)
	}

	// Webhook (no auth required)
	r.POST("/webhook", paymentHandler.HandleWebhook)
}
