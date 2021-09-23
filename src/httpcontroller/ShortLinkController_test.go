package httpcontroller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/andrefebrianto/URL-Shortener-Service/mocks"
	"github.com/andrefebrianto/URL-Shortener-Service/src/httpcontroller"
	model "github.com/andrefebrianto/URL-Shortener-Service/src/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateShortLinkHttpController(t *testing.T) {
	mockShortLinkUseCase := new(mocks.ShortLinkUsecase)

	t.Run("success to create short link", func(t *testing.T) {
		mockShortLinkUseCase.On("Create", mock.Anything, mock.AnythingOfType("*domain.ShortLink")).Return(nil).Once()
		mockShortLink := model.ShortLink{
			Id:             "1",
			Code:           "32fh202je1",
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		}

		bodyRequest, err := json.Marshal(mockShortLink)
		assert.NoError(t, err)

		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/shortlinks", strings.NewReader(string(bodyRequest)))
		assert.NoError(t, err)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks")

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.CreateShortLink(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusCreated, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})

	t.Run("fail to create short link (unprocessed entity)", func(t *testing.T) {
		mockShortLink := model.ShortLink{
			Id:             "1",
			Code:           "32fh202je1",
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		}

		bodyRequest, err := json.Marshal(mockShortLink)
		assert.NoError(t, err)

		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/shortlinks", strings.NewReader(string(bodyRequest)))
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks")

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.CreateShortLink(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusUnprocessableEntity, responseRecorder.Code)
	})

	t.Run("fail to create short link (internal error)", func(t *testing.T) {
		expectedError := errors.New("Internal server error")
		mockShortLinkUseCase.On("Create", mock.Anything, mock.AnythingOfType("*domain.ShortLink")).Return(expectedError).Once()
		mockShortLink := model.ShortLink{
			Id:             "1",
			Code:           "32fh202je1",
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		}

		bodyRequest, err := json.Marshal(mockShortLink)
		assert.NoError(t, err)

		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/shortlinks", strings.NewReader(string(bodyRequest)))
		assert.NoError(t, err)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks")

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.CreateShortLink(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})
}

func TestGetShortlinks(t *testing.T) {
	mockShortLinkUseCase := new(mocks.ShortLinkUsecase)

	t.Run("success to get short links", func(t *testing.T) {
		mockShortLinks := make([]model.ShortLink, 0)
		mockShortLinks = append(mockShortLinks, model.ShortLink{
			Id:             "1",
			Code:           "32fh202je1",
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		})
		mockShortLinks = append(mockShortLinks, model.ShortLink{
			Id:             "2",
			Code:           "89rhf0fwwe",
			Url:            "https://facebook.com",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		})

		mockShortLinkUseCase.On("GetAll", mock.Anything).Return(mockShortLinks, nil).Once()

		ech := echo.New()
		request, err := http.NewRequest(echo.GET, "/api/v1/shortlinks", nil)
		assert.NoError(t, err)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks")

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.GetShortlinks(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})

	t.Run("fail to get short links (not found)", func(t *testing.T) {
		mockShortLinkUseCase.On("GetAll", mock.Anything).Return(nil, errors.New("not found")).Once()

		ech := echo.New()
		request, err := http.NewRequest(echo.GET, "/api/v1/shortlinks", nil)
		assert.NoError(t, err)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks")

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.GetShortlinks(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})

	t.Run("fail to get short links (internal server error)", func(t *testing.T) {
		mockShortLinkUseCase.On("GetAll", mock.Anything).Return(nil, errors.New("internal server error")).Once()

		ech := echo.New()
		request, err := http.NewRequest(echo.GET, "/api/v1/shortlinks", nil)
		assert.NoError(t, err)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks")

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.GetShortlinks(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})
}

func TestUpdateShortLinkHttpController(t *testing.T) {
	mockShortLinkUseCase := new(mocks.ShortLinkUsecase)

	t.Run("success to update short link", func(t *testing.T) {
		mockShortLinkUseCase.On("UpdateByCode", mock.Anything, mock.AnythingOfType("*domain.ShortLink")).Return(nil).Once()
		mockShortLink := model.ShortLink{
			Id:             "1",
			Code:           "32fh202je1",
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		}

		bodyRequest, err := json.Marshal(mockShortLink)
		assert.NoError(t, err)

		ech := echo.New()
		request, err := http.NewRequest(echo.PATCH, "/api/v1/shortlinks", strings.NewReader(string(bodyRequest)))
		assert.NoError(t, err)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks")

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.UpdateShortlinks(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})

	t.Run("fail to update short link (unprocessed entity)", func(t *testing.T) {
		mockShortLink := model.ShortLink{
			Id:             "1",
			Code:           "32fh202je1",
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		}

		bodyRequest, err := json.Marshal(mockShortLink)
		assert.NoError(t, err)

		ech := echo.New()
		request, err := http.NewRequest(echo.PATCH, "/api/v1/shortlinks", strings.NewReader(string(bodyRequest)))
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks")

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.UpdateShortlinks(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusUnprocessableEntity, responseRecorder.Code)
	})

	t.Run("fail to update short link (internal error)", func(t *testing.T) {
		expectedError := errors.New("Internal server error")
		mockShortLinkUseCase.On("UpdateByCode", mock.Anything, mock.AnythingOfType("*domain.ShortLink")).Return(expectedError).Once()
		mockShortLink := model.ShortLink{
			Id:             "1",
			Code:           "32fh202je1",
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		}

		bodyRequest, err := json.Marshal(mockShortLink)
		assert.NoError(t, err)

		ech := echo.New()
		request, err := http.NewRequest(echo.PATCH, "/api/v1/shortlinks", strings.NewReader(string(bodyRequest)))
		assert.NoError(t, err)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks")

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.UpdateShortlinks(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})
}

func TestDeleteShortLinkHttpController(t *testing.T) {
	mockShortLinkUseCase := new(mocks.ShortLinkUsecase)
	code := "32fh202je1"

	t.Run("success to delete short link", func(t *testing.T) {
		mockShortLinkUseCase.On("DeleteByCode", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		ech := echo.New()
		request, err := http.NewRequest(echo.DELETE, "/api/v1/shortlinks/"+code, nil)
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks/:code")
		context.SetParamNames("code")
		context.SetParamValues(code)

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.DeleteShortlinks(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})

	t.Run("fail to delete short links (internal server error)", func(t *testing.T) {
		mockShortLinkUseCase.On("DeleteByCode", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("internal server error")).Once()

		ech := echo.New()
		request, err := http.NewRequest(echo.DELETE, "/api/v1/shortlinks/"+code, nil)
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/shortlinks/:code")
		context.SetParamNames("code")
		context.SetParamValues(code)

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.DeleteShortlinks(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})
}

func TestForwardShortLinkHttpController(t *testing.T) {
	mockShortLinkUseCase := new(mocks.ShortLinkUsecase)
	code := "32fh202je1"

	t.Run("success to redirect short link", func(t *testing.T) {
		mockShortLink := model.ShortLink{
			Id:             "1",
			Code:           "32fh202je1",
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		}
		mockShortLinkUseCase.On("GetByCode", mock.Anything, mock.AnythingOfType("string")).Return(&mockShortLink, nil).Once()
		mockShortLinkUseCase.On("AddCounterByCode", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("uint64")).Return(nil).Once()

		ech := echo.New()
		request, err := http.NewRequest(echo.GET, "/"+code, nil)
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/:code")
		context.SetParamNames("code")
		context.SetParamValues(code)

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.ForwardShortlink(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusFound, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})

	t.Run("fail to redirect short links (not found)", func(t *testing.T) {
		mockShortLinkUseCase.On("GetByCode", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("not found")).Once()

		ech := echo.New()
		request, err := http.NewRequest(echo.GET, "/"+code, nil)
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/:code")
		context.SetParamNames("code")
		context.SetParamValues(code)

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.ForwardShortlink(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})

	t.Run("fail to redirect short links (internal server error)", func(t *testing.T) {
		mockShortLinkUseCase.On("GetByCode", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("internal server error")).Once()

		ech := echo.New()
		request, err := http.NewRequest(echo.GET, "/"+code, nil)
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/:code")
		context.SetParamNames("code")
		context.SetParamValues(code)

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.ForwardShortlink(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})

	t.Run("fail to redirect short links (link expired)", func(t *testing.T) {
		mockShortLink := model.ShortLink{
			Id:             "1",
			Code:           "32fh202je1",
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, -7),
			VisitorCounter: 0,
		}
		mockShortLinkUseCase.On("GetByCode", mock.Anything, mock.AnythingOfType("string")).Return(&mockShortLink, nil).Once()

		ech := echo.New()
		request, err := http.NewRequest(echo.GET, "/"+code, nil)
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/:code")
		context.SetParamNames("code")
		context.SetParamValues(code)

		controller := httpcontroller.ShortLinkHttpController{
			ShortLinkUseCase: mockShortLinkUseCase,
		}
		err = controller.ForwardShortlink(context)
		require.NoError(t, err)

		assert.Equal(t, http.StatusGone, responseRecorder.Code)
		mockShortLinkUseCase.AssertExpectations(t)
	})
}
