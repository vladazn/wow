package domain

type ProofRequest struct {
	Nonce int    `json:"nonce"`
	Key   string `json:"key"`
}
