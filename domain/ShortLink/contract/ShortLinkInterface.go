package contract

import (
	"context"

	model "github.com/andrefebrianto/URL-Shortener-Service/model"
)

type ShortLinkUsecase interface {
	GetAll(ctx context.Context) ([]model.ShortLink, error)
	GetByCode(ctx context.Context, code string) (*model.ShortLink, error)
	Create(ctx context.Context, shortlink *model.ShortLink) error
	UpdateByCode(ctx context.Context, shortlink *model.ShortLink) error
	DeleteByCode(ctx context.Context, code string) error
	AddCounterByCode(ctx context.Context, code string, counter uint64) error
}

type ShortLinkCommandRepository interface {
	Create(ctx context.Context, shortlink *model.ShortLink) error
	UpdateByCode(ctx context.Context, shortlink *model.ShortLink) error
	DeleteByCode(ctx context.Context, code string) error
	AddCounterByCode(ctx context.Context, code string, counter uint64) error
}

type ShortLinkQueryRepository interface {
	GetAll(ctx context.Context) ([]model.ShortLink, error)
	GetByCode(ctx context.Context, code string) (*model.ShortLink, error)
}
