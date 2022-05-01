package repository

import (
	"context"
	"vladazn/wow/internal/pkg/redis"
)

type Challenge interface {
	GetChallenge(ctx context.Context, key string) (int, bool)
	SetChallenge(ctx context.Context, key string, value int) error
	RemoveChallenge(ctx context.Context, key string) error
}

type Repositories struct {
	Challenge Challenge
}

func InitRepo(r redis.Redis) *Repositories {
	return &Repositories{
		Challenge: newChallengeRepo(r),
	}
}
