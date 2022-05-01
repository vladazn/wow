package hash

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"vladazn/wow/internal/domain"
)

type Hasher struct {
	zeros string
}

func NewHasher(difficulty int) *Hasher {
	return &Hasher{strings.Repeat("0", difficulty)}
}

func (h Hasher) Hash(data interface{}) string {
	m, _ := json.Marshal(data)
	return fmt.Sprintf("%x", sha256.Sum256(m))
}

func (h Hasher) IsValid(hash string) bool {
	return strings.HasPrefix(hash, h.zeros)
}

func (h Hasher) Proof(ch chan bool, r *domain.Challenge) bool {
	solved := false
L:
	for {
		select {
		case <-ch:
			break L
		default:
			if strings.HasPrefix(h.Hash(r), h.zeros) {
				solved = true
				break L
			} else {
				r.Nonce++
			}
		}
	}
	return solved
}
