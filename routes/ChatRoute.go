package routes

import (
	"chat-app-portfolio/controllers"

	"github.com/labstack/echo/v4"
)

func InitChatRoute(e *echo.Echo) {

	e.POST("/chat", controllers.CreateNewChatRoom)

	e.GET("/chat", controllers.FetchExistingChatRoom)
	e.DELETE("/chat/:id", controllers.DeleteExistingChatRoom)

}
