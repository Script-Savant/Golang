package routes

import (
	"go-blog/controllers"
	"go-blog/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// initialize all routes for the application
func SetupRoutes(db *gorm.DB) *gin.Engine {
	/*
	- create a new gin router with default params
	- initialize controllers
	- define public routes
	- define protected routes
	*/

	// 1. gin router
	r := gin.Default()

	// 2. controllers
	authController := controllers.NewAuthController(db)
	userController := controllers.NewUserController(db)
	postController := controllers.NewPostController(db)
	commentController := controllers.NewCommentcontroller(db)

	// 3. public routes - no authentication required
	public := r.Group("/api")
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
		public.GET("/posts", postController.GetPosts)
		public.GET("/posts/:id", postController.GetPost)
		public.POST("/posts/:id/share", postController.SharePost)
		public.GET("/posts/:postId/comments", commentController.GetComments)
	}

	// 4. protected routes - require authentication
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		protected.GET("/profile", userController.GetProfile)
		protected.PUT("/profile", userController.UpdateProfile)

		// Post routes
		protected.POST("/posts", postController.CreatePost)
		protected.PUT("/posts/:id", postController.UpdatePost)
		protected.DELETE("/posts/:id", postController.DeletePost)
		protected.POST("/posts/:id/:action", postController.LikePost) 

		// Comment routes
		protected.POST("/posts/:postId/comments", commentController.CreateComment)
		protected.POST("/comments/:commentId/:action", commentController.LikeComment) 
	}

	return r
}