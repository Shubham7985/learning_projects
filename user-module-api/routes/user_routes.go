package routes

import (
	"user-module-api/controllers"
	"user-module-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signup", controllers.SignUp)
			auth.POST("/signin", controllers.SignIn)
			auth.POST("/forgot-password", controllers.ForgotPassword)
		}

		user := api.Group("/users", middlewares.AuthMiddleware("Secret_key"))
		{
			user.POST("/profile", controllers.CreateProfile)
			user.GET("/list", controllers.ListUsers)
			user.GET("/:id", controllers.GetUserByID) // url param
		}
	}
}
