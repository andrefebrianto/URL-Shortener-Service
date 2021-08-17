package usecase

import (
	"context"
	"time"

	"github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/repository/command"
	"github.com/andrefebrianto/URL-Shortener-Service/src/domain/ShortLink/repository/query"
	model "github.com/andrefebrianto/URL-Shortener-Service/src/model"
)

type ShortLinkUseCase struct {
	cassandraCommandRepository command.CassandraCommandRepository
	cassandraQueryRepository   query.CassandraQueryRepository
	contextTimeout             time.Duration
}

func CreateShortLinkUseCase(command command.CassandraCommandRepository, query query.CassandraQueryRepository, timeout time.Duration) ShortLinkUseCase {
	return ShortLinkUseCase{
		cassandraCommandRepository: command,
		cassandraQueryRepository:   query,
		contextTimeout:             timeout,
	}
}

func (usecase ShortLinkUseCase) Create(mainContext context.Context, shortlink *model.ShortLink) (*model.ShortLink, error) {
	return nil, nil
}

func (usecase ShortLinkUseCase) GetAll(ctx context.Context) ([]model.ShortLink, error) {
	return nil, nil
}

func (usecase ShortLinkUseCase) GetByCode(ctx context.Context, code string) (model.ShortLink, error) {
	return model.ShortLink{}, nil
}

func (usecase ShortLinkUseCase) UpdateByCode(ctx context.Context, shortlink *model.ShortLink) error {
	return nil
}

func (usecase ShortLinkUseCase) DeleteByCode(ctx context.Context, code string) error {
	return nil
}
