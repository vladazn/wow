package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"vladazn/wow/internal/domain"
	"vladazn/wow/internal/pkg/hash"
	"vladazn/wow/internal/repository"
	mock_repository "vladazn/wow/test/mocks/repository"
)

func TestCreateChallenge(t *testing.T) {
	rand.Seed(101)

	c := domain.Challenge{}
	crateChallenge(&c)
	require.Equal(t, "bCXQhnFnkD", c.Key)
	require.Equal(t, 1384714, c.Check)
	require.Equal(t, 0, c.Nonce)
}

func TestGetChallenge_challenge_does_not_exist(t *testing.T) {
	rand.Seed(101)

	ctx := context.Background()

	mc := mock_repository.NewMockChallenge(gomock.NewController(t))
	s := newWowService(&Params{
		Repo: &repository.Repositories{Challenge: mc},
	})

	mc.EXPECT().GetChallenge(gomock.Any(), "bCXQhnFnkD").Return(0, false)
	mc.EXPECT().SetChallenge(gomock.Any(), "bCXQhnFnkD", gomock.Any())

	c, err := s.GetChallenge(ctx)

	require.NoError(t, err)
	require.Equal(t, "bCXQhnFnkD", c.Key)
	require.Equal(t, 1384714, c.Check)
	require.Equal(t, 0, c.Nonce)
}

func TestGetChallenge_challenge_does_exist(t *testing.T) {
	rand.Seed(101)

	ctx := context.Background()

	mc := mock_repository.NewMockChallenge(gomock.NewController(t))
	s := newWowService(&Params{
		Repo: &repository.Repositories{Challenge: mc},
	})

	mc.EXPECT().GetChallenge(gomock.Any(), "bCXQhnFnkD").Return(0, true)
	mc.EXPECT().GetChallenge(gomock.Any(), "wXxkQZQnpu").Return(0, false)
	mc.EXPECT().SetChallenge(gomock.Any(), "wXxkQZQnpu", gomock.Any())

	c, err := s.GetChallenge(ctx)

	require.NoError(t, err)
	require.Equal(t, "wXxkQZQnpu", c.Key)
	require.Equal(t, 9156171, c.Check)
	require.Equal(t, 0, c.Nonce)
}

func TestGetWisdom_valid_response(t *testing.T) {

	rand.Seed(202)

	ctx := context.Background()

	mc := mock_repository.NewMockChallenge(gomock.NewController(t))
	s := newWowService(&Params{
		Repo:   &repository.Repositories{Challenge: mc},
		Hasher: hash.NewHasher(3),
	})

	mc.EXPECT().GetChallenge(gomock.Any(), "bCXQhnFnkD").Return(1384714, true)
	mc.EXPECT().RemoveChallenge(gomock.Any(), gomock.Any())

	v, err := s.GetWisdom(ctx, &domain.Challenge{
		Key:   "bCXQhnFnkD",
		Check: 1384714,
		Nonce: 1639,
	})
	require.NoError(t, err)
	require.Equal(t, "wisdom 4", v)
}

func TestGetWisdom_invalid_nonce(t *testing.T) {

	rand.Seed(202)

	ctx := context.Background()

	mc := mock_repository.NewMockChallenge(gomock.NewController(t))
	s := newWowService(&Params{
		Repo:   &repository.Repositories{Challenge: mc},
		Hasher: hash.NewHasher(3),
	})

	_, err := s.GetWisdom(ctx, &domain.Challenge{
		Key:   "bCXQhnFnkD",
		Check: 1384714,
		Nonce: 0,
	})
	require.Error(t, err)
	require.Equal(t, "invalid hash", err.Error())

}

func TestGetWisdom_invalid_key(t *testing.T) {

	rand.Seed(202)

	ctx := context.Background()

	mc := mock_repository.NewMockChallenge(gomock.NewController(t))
	s := newWowService(&Params{
		Repo:   &repository.Repositories{Challenge: mc},
		Hasher: hash.NewHasher(3),
	})

	mc.EXPECT().GetChallenge(gomock.Any(), "bCXQhnFnkD").Return(138471555, true)

	_, err := s.GetWisdom(ctx, &domain.Challenge{
		Key:   "bCXQhnFnkD",
		Check: 1384714,
		Nonce: 1639,
	})
	require.Error(t, err)
	require.Equal(t, "invalid key", err.Error())
}

func TestGetWisdom_invalid_key_not_exist(t *testing.T) {

	rand.Seed(202)

	ctx := context.Background()

	mc := mock_repository.NewMockChallenge(gomock.NewController(t))
	s := newWowService(&Params{
		Repo:   &repository.Repositories{Challenge: mc},
		Hasher: hash.NewHasher(3),
	})

	mc.EXPECT().GetChallenge(gomock.Any(), "bCXQhnFnkD").Return(0, false)

	_, err := s.GetWisdom(ctx, &domain.Challenge{
		Key:   "bCXQhnFnkD",
		Check: 1384714,
		Nonce: 1639,
	})
	require.Error(t, err)
	require.Equal(t, "invalid key", err.Error())
}
