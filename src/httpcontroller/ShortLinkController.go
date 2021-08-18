package httpcontroller

import (
	"net/http"
	"time"

	"github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/contract"
	"github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/repository/command"
	"github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/repository/query"
	"github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/usecase"
	model "github.com/andrefebrianto/URL-Shortener-Service/src/model"
	"github.com/labstack/echo/v4"
)

type ShortLinkHttpController struct {
	shortLinkUseCase contract.ShortLinkUsecase
}

var (
	shortLinkCommandRepo = command.CreateCassandraCommandRepository()
	shortLinkQueryRepo   = query.CreateCassandraQueryRepository()
	shortLinkUseCase     = usecase.CreateShortLinkUseCase(shortLinkCommandRepo, shortLinkQueryRepo, 2*time.Second)
	controller           = ShortLinkHttpController{}
)

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

func (controller ShortLinkHttpController) GetShortlinks(context echo.Context) error {
	return context.JSON(http.StatusOK, model.CustomError{Message: "Short links retrieved"})
}

func (controller ShortLinkHttpController) UpdateShortlinks(context echo.Context) error {
	return context.JSON(http.StatusOK, model.CustomError{Message: "Short link updated"})
}

func (controller ShortLinkHttpController) DeleteShortlinks(context echo.Context) error {
	return context.JSON(http.StatusOK, model.CustomError{Message: "Short link deleted"})
}

func (controller ShortLinkHttpController) ForwardShortlink(context echo.Context) error {
	return context.JSON(http.StatusOK, model.CustomError{Message: "Short link forwarded"})
}
