package domain

type Challenge struct {
	Key   string `json:"key"`
	Check int    `json:"check"`
	Nonce int    `json:"nonce,omitempty"`
}
