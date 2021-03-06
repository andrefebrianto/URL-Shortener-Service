package httpcontroller

import (
	"net/http"
	"time"

	"github.com/andrefebrianto/URL-Shortener-Service/domain/ShortLink/contract"
	model "github.com/andrefebrianto/URL-Shortener-Service/model"
	"github.com/labstack/echo/v4"
)

type ShortLinkHttpController struct {
	ShortLinkUseCase contract.ShortLinkUsecase
}

func CreateShortLinkHttpController(server *echo.Echo, shortLinkUseCase contract.ShortLinkUsecase) {
	controller := ShortLinkHttpController{ShortLinkUseCase: shortLinkUseCase}

	server.POST("/api/v1/shortlinks", controller.CreateShortLink)
	server.GET("/api/v1/shortlinks", controller.GetShortlinks)
	server.PATCH("/api/v1/shortlinks", controller.UpdateShortlinks)
	server.DELETE("/api/v1/shortlinks/:code", controller.DeleteShortlinks)
	server.GET("/:code", controller.ForwardShortlink)
}

func (controller ShortLinkHttpController) CreateShortLink(context echo.Context) error {
	var shortLink model.ShortLink
	err := context.Bind(&shortLink)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, model.HttpResponseObject{Message: err.Error()})
	}

	ctx := context.Request().Context()
	err = controller.ShortLinkUseCase.Create(ctx, &shortLink)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.HttpResponseObject{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, model.HttpResponseObject{Message: "Short link created", Data: shortLink})
}

func (controller ShortLinkHttpController) GetShortlinks(context echo.Context) error {
	ctx := context.Request().Context()

	shortLinks, err := controller.ShortLinkUseCase.GetAll(ctx)
	if err != nil {
		if err.Error() == "not found" {
			return context.JSON(http.StatusNotFound, model.HttpResponseObject{Message: "Short link not found"})
		}
		return context.JSON(http.StatusInternalServerError, model.HttpResponseObject{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, model.HttpResponseObject{Message: "Short links retrieved", Data: shortLinks})
}

func (controller ShortLinkHttpController) UpdateShortlinks(context echo.Context) error {
	var shortLink model.ShortLink
	err := context.Bind(&shortLink)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, model.HttpResponseObject{Message: err.Error()})
	}

	ctx := context.Request().Context()
	err = controller.ShortLinkUseCase.UpdateByCode(ctx, &shortLink)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.HttpResponseObject{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, model.HttpResponseObject{Message: "Short link updated"})
}

func (controller ShortLinkHttpController) DeleteShortlinks(context echo.Context) error {
	code := context.Param("code")
	ctx := context.Request().Context()

	err := controller.ShortLinkUseCase.DeleteByCode(ctx, code)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.HttpResponseObject{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, model.HttpResponseObject{Message: "Short link deleted"})
}

func (controller ShortLinkHttpController) ForwardShortlink(context echo.Context) error {
	code := context.Param("code")
	ctx := context.Request().Context()

	shortLink, err := controller.ShortLinkUseCase.GetByCode(ctx, code)
	if err != nil {
		if err.Error() == "not found" {
			return context.JSON(http.StatusNotFound, model.HttpResponseObject{Message: "Short link not found"})
		}
		return context.JSON(http.StatusInternalServerError, model.HttpResponseObject{Message: err.Error()})
	}

	if time.Now().Local().After(shortLink.ExpiredAt) {
		return context.JSON(http.StatusGone, model.HttpResponseObject{Message: "Short link expired"})
	}

	_ = controller.ShortLinkUseCase.AddCounterByCode(ctx, shortLink.Code, shortLink.VisitorCounter+1)

	return context.Redirect(http.StatusFound, shortLink.Url)
}
