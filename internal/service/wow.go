package service

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"vladazn/wow/internal/domain"
	"vladazn/wow/internal/pkg/hash"
	"vladazn/wow/internal/repository"
)

type WowService struct {
	hasher *hash.Hasher
	m      sync.Map
	repo   *repository.Repositories
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var wisdoms = []string{
	"The fool doth think he is wise, but the wise man knows himself to be a fool.",
	"It is better to remain silent at the risk of being thought a fool, than to talk and remove all doubt of it.",
	"Knowing yourself is the beginning of all wisdom.",
	"Think before you speak. Read before you think.",
	"Never let your sense of morals prevent you from doing what is right.",
}

func newWowService(p *Params) Wow {
	return &WowService{
		hasher: p.Hasher,
		m:      sync.Map{},
		repo:   p.Repo,
	}
}

// GetChallenge creates a challenge with uniq key and saves it to cache
func (w *WowService) GetChallenge(ctx context.Context) (*domain.Challenge, error) {
	ex := true
	challenge := &domain.Challenge{}

	for ex { // make sure to be unique
		crateChallenge(challenge)
		_, ex = w.repo.Challenge.GetChallenge(ctx, challenge.Key)
	}

	err := w.repo.Challenge.SetChallenge(ctx, challenge.Key, challenge.Check)
	if err != nil {
		return nil, err
	}

	return challenge, err
}

// GetWisdom validates hash and compares the key data, in case of success returns one of the quotes
func (w *WowService) GetWisdom(ctx context.Context, c *domain.Challenge) (string, error) {
	if v := w.hasher.IsValid(w.hasher.Hash(c)); !v {
		return "", fmt.Errorf("invalid hash")
	}

	n, ex := w.repo.Challenge.GetChallenge(ctx, c.Key)
	if !ex || n != c.Check {
		return "", fmt.Errorf("invalid key")
	}

	err := w.repo.Challenge.RemoveChallenge(ctx, c.Key)
	if err != nil {
		return "", nil
	}

	return wisdoms[rand.Intn(len(wisdoms))], nil
}

//crateChallenge randomly generates 10 letter key and random int value for double check
func crateChallenge(c *domain.Challenge) {

	k := make([]byte, 10)
	for i := range k {
		k[i] = letters[rand.Intn(len(letters))]
	}

	v := 1000000 + rand.Intn(8999999)

	c.Key = string(k)
	c.Check = v
}
