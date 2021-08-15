package contract

import (
	"context"

	model "github.com/andrefebrianto/URL-Shortener-Service/src/model"
)

type ShortLinkUsecase interface {
	Fetch(ctx context.Context) ([]model.ShortLink, error)
	GetByCode(ctx context.Context, code string) (model.ShortLink, error)
	Create(ctx context.Context, shortlink *model.ShortLink) error
	UpdateByCode(ctx context.Context, shortlink *model.ShortLink) error
	DeleteByCode(ctx context.Context, code string) error
}

type ShortLinkRepository interface {
	Fetch(ctx context.Context) ([]model.ShortLink, error)
	GetByCode(ctx context.Context, code string) (model.ShortLink, error)
	Create(ctx context.Context, shortlink *model.ShortLink) error
	UpdateByCode(ctx context.Context, shortlink *model.ShortLink) error
	DeleteByCode(ctx context.Context, code string) error
}
