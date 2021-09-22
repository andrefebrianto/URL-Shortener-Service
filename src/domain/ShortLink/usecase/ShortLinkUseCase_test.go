package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/andrefebrianto/URL-Shortener-Service/mocks"
	"github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/usecase"
	ShortLinkUseCase "github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/usecase"
	model "github.com/andrefebrianto/URL-Shortener-Service/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateShortLink(t *testing.T) {
	mockCassandraCommandRepo := new(mocks.ShortLinkCommandRepository)
	mockCassandraQueryRepo := new(mocks.ShortLinkQueryRepository)

	t.Run("success to create short link", func(t *testing.T) {
		mockCassandraCommandRepo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		usecase := ShortLinkUseCase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		mockShortLink := model.ShortLink{
			Url: "https://inpoin.id",
		}
		err := usecase.Create(context.TODO(), &mockShortLink)
		assert.Nil(t, err)
		mockCassandraCommandRepo.AssertExpectations(t)
	})

	t.Run("fail to create short link", func(t *testing.T) {
		expectedError := errors.New("Database error")
		mockCassandraCommandRepo.On("Create", mock.Anything, mock.Anything).Return(expectedError).Once()

		usecase := ShortLinkUseCase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		mockShortLink := model.ShortLink{
			Url: "https://inpoin.id",
		}
		err := usecase.Create(context.TODO(), &mockShortLink)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		mockCassandraCommandRepo.AssertExpectations(t)
	})
}

func TestGetAllShortLink(t *testing.T) {
	mockCassandraCommandRepo := new(mocks.ShortLinkCommandRepository)
	mockCassandraQueryRepo := new(mocks.ShortLinkQueryRepository)

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

		mockCassandraQueryRepo.On("GetAll", mock.Anything).Return(mockShortLinks, nil).Once()

		usecase := ShortLinkUseCase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		shortLinks, err := usecase.GetAll(context.TODO())
		assert.ObjectsAreEqualValues(mockShortLinks, shortLinks)
		assert.Nil(t, err)
		mockCassandraQueryRepo.AssertExpectations(t)
	})

	t.Run("fail to get short links", func(t *testing.T) {
		expectedError := errors.New("Short link(s) not found")
		mockCassandraQueryRepo.On("GetAll", mock.Anything).Return(nil, expectedError).Once()

		usecase := ShortLinkUseCase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		shortLinks, err := usecase.GetAll(context.TODO())
		assert.Nil(t, shortLinks)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		mockCassandraQueryRepo.AssertExpectations(t)
	})
}

func TestGetShortLinkByCode(t *testing.T) {
	mockCassandraCommandRepo := new(mocks.ShortLinkCommandRepository)
	mockCassandraQueryRepo := new(mocks.ShortLinkQueryRepository)
	code := "32fh202je1"

	t.Run("success to get short link", func(t *testing.T) {
		mockShortLink := &model.ShortLink{
			Id:             "1",
			Code:           code,
			Url:            "https://google.co.id",
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			ExpiredAt:      time.Now().Local().AddDate(0, 0, 7),
			VisitorCounter: 0,
		}

		mockCassandraQueryRepo.On("GetByCode", mock.Anything, mock.AnythingOfType("string")).Return(mockShortLink, nil).Once()

		usecase := usecase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		shortLink, err := usecase.GetByCode(context.TODO(), code)
		assert.ObjectsAreEqualValues(mockShortLink, shortLink)
		assert.Nil(t, err)
		mockCassandraQueryRepo.AssertExpectations(t)
	})

	t.Run("fail to get short link", func(t *testing.T) {
		expectedError := errors.New("Short link not found")
		mockCassandraQueryRepo.On("GetByCode", mock.Anything, mock.AnythingOfType("string")).Return(nil, expectedError).Once()

		usecase := usecase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		shortLink, err := usecase.GetByCode(context.TODO(), code)
		assert.Nil(t, shortLink)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		mockCassandraQueryRepo.AssertExpectations(t)
	})
}

func TestUpdateShortLinkByCode(t *testing.T) {
	mockCassandraCommandRepo := new(mocks.ShortLinkCommandRepository)
	mockCassandraQueryRepo := new(mocks.ShortLinkQueryRepository)
	code := "32fh202je1"

	t.Run("success to update short link", func(t *testing.T) {
		mockCassandraCommandRepo.On("UpdateByCode", mock.Anything, mock.Anything).Return(nil).Once()

		usecase := usecase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		shortLink := &model.ShortLink{
			Code: code,
			Url:  "https://google.co.id",
		}
		err := usecase.UpdateByCode(context.TODO(), shortLink)
		updatedYear, updatedMonth, updatedDate := shortLink.ExpiredAt.Local().Date()
		expectedYear, expectedMonth, expectedDate := time.Now().Local().AddDate(0, 0, 7).Date()
		assert.Equal(t, expectedYear, updatedYear)
		assert.Equal(t, expectedMonth, updatedMonth)
		assert.Equal(t, expectedDate, updatedDate)
		assert.Nil(t, err)
		mockCassandraCommandRepo.AssertExpectations(t)
	})

	t.Run("fail to update short link", func(t *testing.T) {
		expectedError := errors.New("Short link not found")
		mockCassandraCommandRepo.On("UpdateByCode", mock.Anything, mock.Anything).Return(expectedError).Once()

		usecase := usecase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		shortLink := &model.ShortLink{
			Code: code,
			Url:  "https://google.co.id",
		}
		err := usecase.UpdateByCode(context.TODO(), shortLink)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		mockCassandraCommandRepo.AssertExpectations(t)
	})
}

func TestDeleteShortLinkByCode(t *testing.T) {
	mockCassandraCommandRepo := new(mocks.ShortLinkCommandRepository)
	mockCassandraQueryRepo := new(mocks.ShortLinkQueryRepository)
	code := "32fh202je1"

	t.Run("success to delete short link", func(t *testing.T) {
		mockCassandraCommandRepo.On("DeleteByCode", mock.Anything, mock.Anything).Return(nil).Once()

		usecase := usecase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		err := usecase.DeleteByCode(context.TODO(), code)
		assert.Nil(t, err)
		mockCassandraCommandRepo.AssertExpectations(t)
	})

	t.Run("fail to delete short link", func(t *testing.T) {
		expectedError := errors.New("Short link not found")
		mockCassandraCommandRepo.On("DeleteByCode", mock.Anything, mock.Anything).Return(expectedError).Once()

		usecase := usecase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		err := usecase.DeleteByCode(context.TODO(), code)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		mockCassandraCommandRepo.AssertExpectations(t)
	})
}

func TestAddShortLinkCounterByCode(t *testing.T) {
	mockCassandraCommandRepo := new(mocks.ShortLinkCommandRepository)
	mockCassandraQueryRepo := new(mocks.ShortLinkQueryRepository)
	code := "32fh202je1"

	t.Run("success to add short link counter", func(t *testing.T) {
		mockCassandraCommandRepo.On("AddCounterByCode", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("uint64")).Return(nil).Once()

		usecase := usecase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		err := usecase.AddCounterByCode(context.TODO(), code, 1)
		assert.Nil(t, err)
		mockCassandraCommandRepo.AssertExpectations(t)
	})

	t.Run("fail to add short link counter", func(t *testing.T) {
		expectedError := errors.New("Short link not found")
		mockCassandraCommandRepo.On("AddCounterByCode", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("uint64")).Return(expectedError).Once()

		usecase := usecase.CreateShortLinkUseCase(mockCassandraCommandRepo, mockCassandraQueryRepo, time.Second*2)
		err := usecase.AddCounterByCode(context.TODO(), code, 1)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		mockCassandraCommandRepo.AssertExpectations(t)
	})
}
