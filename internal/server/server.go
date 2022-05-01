package server

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"os/signal"
	"syscall"
	"vladazn/wow/internal/api"
	"vladazn/wow/internal/domain"
	grpcserver "vladazn/wow/internal/pkg/grpc/server"
	"vladazn/wow/internal/pkg/hash"
	"vladazn/wow/internal/pkg/redis"
	"vladazn/wow/internal/repository"
	"vladazn/wow/internal/service"
)

func Run(configPath string) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := &domain.Config{}
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return fmt.Errorf("error creating configs: %w", err)
	}

	printConfigs(cfg)

	r, err := redis.NewRedisConnection(cfg.Redis)
	if err != nil {
		return errors.Wrap(err, "error connecting to redis")
	}

	repo := repository.InitRepo(r)

	h := hash.NewHasher(cfg.Hash.Difficulty)

	s := service.InitServices(&service.Params{
		Hasher: h,
		Repo:   repo,
	})

	wowServer := grpcserver.NewWowServer(s)

	errChan := make(chan error)

	go func() {
		errChan <- api.RunApiServer(ctx, wowServer, cfg.Listen)
	}()

	grpcServer := grpcserver.NewGrpcServer(cfg.Listen.Grpc, s)
	go func() {
		errChan <- grpcServer.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err = <-errChan:
		logrus.Info("stopping gracefully...")
	case q := <-quit:
		logrus.Infof("%s signal received, stopping gracefully...", q.String())
	}

	if err := r.Close(); err != nil {
		logrus.Error(err)
	}
	logrus.Info("redis has closed")

	grpcServer.Stop()

	return nil
}

func printConfigs(cfg *domain.Config) {
	_ = yaml.NewEncoder(os.Stdout).Encode(cfg)
}
