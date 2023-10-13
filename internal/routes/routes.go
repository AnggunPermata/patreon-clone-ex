package routes

import (
	"github.com/anggunpermata/patreon-clone/internal/controller"
	"github.com/labstack/echo/v4"
)

func NewRoutes(e *echo.Echo, handler *controller.BackendHandler) {
	e.GET("/healthcheck", handler.HealthcheckHandler)
	e.GET("/upload", handler.RouteUploadFile)
	e.POST("/upload/images", handler.UploadFile)
	e.GET("/file/show/:file_name", handler.GetFile)
	e.GET("/signup", handler.UserSignup)
	e.POST("/signup", handler.UserSignup)
	e.GET("/login", handler.UserLogin)
	e.POST("/login", handler.UserLogin)
}
