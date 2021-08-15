package httpcontroller

import (
	"net/http"

	"github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/usecase"
	model "github.com/andrefebrianto/URL-Shortener-Service/src/model"
	"github.com/labstack/echo/v4"
)

type ShortLinkHttpController struct {
	shortLinkUseCase usecase.ShortLinkUseCase
}

func CreateShortLinkHttpController(server *echo.Echo) {
	server.POST("/api/v1/shortlinks", nil)
	server.GET("/api/v1/shortlinks", nil)
	server.GET("/api/v1/shortlinks/:id", nil)
	server.DELETE("/api/v1/shortlinks/:code", nil)
	server.GET("/:code", nil)
}

func (controller ShortLinkHttpController) CreateShortLink(context echo.Context) error {
	return context.JSON(http.StatusOK, model.CustomError{Message: "Short link created"})
}
