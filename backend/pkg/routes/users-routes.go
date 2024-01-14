package routes

import (
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/controllers"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func UsersRouter(usersRouter *gin.RouterGroup) {
	usersRouter.POST("/register", controllers.Register)
	usersRouter.POST("/login", controllers.Login)
	usersRouter.GET("/control",middlewares.AuthMiddleware(), controllers.Control)
}