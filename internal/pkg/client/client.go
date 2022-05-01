package client

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
	"vladazn/wow/internal/domain"
)

type WowClient struct {
	host string
}

type ErrorMsg struct {
	Msg string
}

type Response struct {
	Response json.RawMessage
	Err      *ErrorMsg
}

const (
	challengeUrl = "/challenge"
	wisdomUrl    = "/wisdom"
)

func NewWowClient(host string) *WowClient {
	return &WowClient{
		host: host,
	}
}

func (w WowClient) GetChallenge(c *domain.Challenge) error {
	d := Response{}
	client := http.Client{Timeout: 10 * time.Minute}

	url := w.host + challengeUrl

	res, err := client.Get(url)
	if err != nil {
		logrus.Errorf("http request err => %v", err)
		return err
	}

	defer res.Body.Close()

	blob, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(blob, &d)
	if err != nil {
		return err
	}

	return json.Unmarshal(d.Response, &c)
}

func (w WowClient) GetWisdom(c *domain.Challenge) (string, error) {
	d := Response{}
	client := http.Client{Timeout: 10 * time.Minute}

	url := fmt.Sprintf("%v?key=%v&check=%v&nonce=%v",
		w.host+wisdomUrl, c.Key, c.Check, c.Nonce)

	res, err := client.Get(url)
	if err != nil {
		logrus.Errorf("http request err => %v", err)
		return "", err
	}

	defer res.Body.Close()

	blob, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(blob, &d)
	if err != nil {
		return "", err
	}

	wr := domain.Wisdom{}
	err = json.Unmarshal(d.Response, &wr)

	return wr.Quote, err
}
