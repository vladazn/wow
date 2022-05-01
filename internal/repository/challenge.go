package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
	"vladazn/wow/internal/pkg/redis"
)

type ChallengeRepo struct {
	r   redis.Redis
	ttl time.Duration
}

const defaultTtl = 10 * time.Minute

func newChallengeRepo(r redis.Redis) *ChallengeRepo {
	return &ChallengeRepo{r: r, ttl: defaultTtl}
}

func (c *ChallengeRepo) GetChallenge(ctx context.Context, key string) (int, bool) {
	var v int
	err, ex := c.r.Get(ctx, key, &v)
	if !ex {
		return 0, false
	}

	if err != nil {
		logrus.Error(err)
		return 0, false
	}

	return v, true
}

func (c *ChallengeRepo) SetChallenge(ctx context.Context, key string, value int) error {
	return c.r.Set(ctx, key, value, c.ttl)
}

func (c *ChallengeRepo) RemoveChallenge(ctx context.Context, key string) error {
	return c.r.Del(ctx, key)
}
