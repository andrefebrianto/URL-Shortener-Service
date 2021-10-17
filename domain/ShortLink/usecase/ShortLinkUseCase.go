package usecase

import (
	"context"
	"time"

	"github.com/andrefebrianto/URL-Shortener-Service/domain/ShortLink/contract"
	"github.com/andrefebrianto/URL-Shortener-Service/generator"
	model "github.com/andrefebrianto/URL-Shortener-Service/model"
)

type ShortLinkUseCase struct {
	cassandraCommandRepository contract.ShortLinkCommandRepository
	cassandraQueryRepository   contract.ShortLinkQueryRepository
	contextTimeout             time.Duration
}

func CreateShortLinkUseCase(command contract.ShortLinkCommandRepository, query contract.ShortLinkQueryRepository, timeout time.Duration) ShortLinkUseCase {
	return ShortLinkUseCase{
		cassandraCommandRepository: command,
		cassandraQueryRepository:   query,
		contextTimeout:             timeout,
	}
}

func (usecase ShortLinkUseCase) Create(ctx context.Context, shortlink *model.ShortLink) error {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	shortlink.Code, _ = generator.GenerateRandomString(10)
	shortlink.CreatedAt = time.Now().Local()
	shortlink.UpdatedAt = time.Now().Local()
	shortlink.ExpiredAt = time.Now().Local().AddDate(0, 0, 7)
	shortlink.VisitorCounter = 0

	err := usecase.cassandraCommandRepository.Create(contextWithTimeOut, shortlink)
	if err != nil {
		return err
	}
	return nil
}

func (usecase ShortLinkUseCase) GetAll(ctx context.Context) ([]model.ShortLink, error) {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	shortLinks, err := usecase.cassandraQueryRepository.GetAll(contextWithTimeOut)
	if err != nil {
		return nil, err
	}
	return shortLinks, nil
}

func (usecase ShortLinkUseCase) GetByCode(ctx context.Context, code string) (*model.ShortLink, error) {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	shortLink, err := usecase.cassandraQueryRepository.GetByCode(contextWithTimeOut, code)
	if err != nil {
		return nil, err
	}

	return shortLink, nil
}

func (usecase ShortLinkUseCase) UpdateByCode(ctx context.Context, shortlink *model.ShortLink) error {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	shortlink.ExpiredAt = time.Now().Local().AddDate(0, 0, 7)

	err := usecase.cassandraCommandRepository.UpdateByCode(contextWithTimeOut, shortlink)
	if err != nil {
		return err
	}
	return nil
}

func (usecase ShortLinkUseCase) DeleteByCode(ctx context.Context, code string) error {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	err := usecase.cassandraCommandRepository.DeleteByCode(contextWithTimeOut, code)
	if err != nil {
		return err
	}
	return nil
}

func (usecase ShortLinkUseCase) AddCounterByCode(ctx context.Context, code string, counter uint64) error {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	err := usecase.cassandraCommandRepository.AddCounterByCode(contextWithTimeOut, code, counter)
	if err != nil {
		return err
	}
	return nil
}
