package httpcontroller

import (
	"net/http"

	usecase "github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/contract"
	model "github.com/andrefebrianto/URL-Shortener-Service/src/model"
	"github.com/labstack/echo/v4"
)

type ShortLinkHttpController struct {
	shortLinkUseCase usecase.ShortLinkUsecase
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
