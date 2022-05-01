package client

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"os/signal"
	"syscall"
	"vladazn/wow/internal/domain"
	"vladazn/wow/internal/pkg/client"
	"vladazn/wow/internal/pkg/hash"
)

func Run(configPath string) error {

	cfg := &domain.Config{}
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return fmt.Errorf("error creating configs: %w", err)
	}

	printConfigs(cfg)

	h := hash.NewHasher(cfg.Hash.Difficulty)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	cl := client.NewWowClient(cfg.Client.Host)

	var err error

	stopChan := make(chan bool)

	go func() {
		for range quit {
			close(stopChan)
		}
	}()

L:
	for {
		select {

		case <-stopChan:
			break L

		default:
			err = FinProof(stopChan, cl, h)
			if err != nil {
				break L
			}
		}
	}

	return err
}

func printConfigs(cfg *domain.Config) {
	_ = yaml.NewEncoder(os.Stdout).Encode(cfg)
}

func FinProof(stopChan chan bool, cl *client.WowClient, h *hash.Hasher) error {
	c := domain.Challenge{}
	err := cl.GetChallenge(&c)
	if err != nil {
		return err
	}

	solved := h.Proof(stopChan, &c)

	if solved {
		wisdom, err := cl.GetWisdom(&c)
		if err != nil {
			return err
		}
		logrus.Info(wisdom)
	}

	return nil
}
