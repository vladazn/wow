package grpcserver

import (
	"context"
	"vladazn/wow/internal/domain"
	"vladazn/wow/internal/service"
	"vladazn/wow/proto/gen/go/proto/wow"
)

type WowServer struct {
	wow.UnimplementedWowServer
	s *service.Services
}

func NewWowServer(s *service.Services) *WowServer {
	return &WowServer{
		s: s,
	}
}

func (w *WowServer) GetChallenge(ctx context.Context, _ *wow.Empty) (*wow.ChallengeResponse,
	error) {

	ch, err := w.s.Wow.GetChallenge(ctx)
	if err != nil {
		return nil, err
	}

	return &wow.ChallengeResponse{
		Response: &wow.Challenge{
			Key:   ch.Key,
			Check: int32(ch.Check),
		},
	}, nil
}

func (w *WowServer) GetWisdom(ctx context.Context, req *wow.WisdomRequest) (*wow.WisdomResponse,
	error) {

	wisdom, err := w.s.Wow.GetWisdom(ctx, &domain.Challenge{
		Key:   req.Key,
		Check: int(req.Check),
		Nonce: int(req.Nonce),
	})

	if err != nil {
		return nil, err
	}

	return &wow.WisdomResponse{
		Response: &wow.Wisdom{
			Quote: wisdom,
		},
	}, nil
}
