package service

import (
	"context"
	"vladazn/wow/internal/domain"
	"vladazn/wow/internal/pkg/hash"
	"vladazn/wow/internal/repository"
)

type Params struct {
	Hasher *hash.Hasher
	Repo   *repository.Repositories
}

type Wow interface {
	GetChallenge(ctx context.Context) (*domain.Challenge, error)
	GetWisdom(ctx context.Context, c *domain.Challenge) (string, error)
}

type Services struct {
	Wow Wow
}

func InitServices(p *Params) *Services {
	return &Services{Wow: newWowService(p)}
}
