package hash

import (
	"github.com/stretchr/testify/require"
	"testing"
	"vladazn/wow/internal/domain"
)

const difficulty = 3

func TestHash(t *testing.T) {
	h := NewHasher(difficulty)
	hashed := "5aa762ae383fbb727af3c7a36d4940a5b8c40a989452d2304fc958ff3f354e7a"
	require.Equal(t, hashed, h.Hash("hello"))
}

func TestIsValid(t *testing.T) {

	h := NewHasher(difficulty)

	hashed := "000762ae383fbb727af3c7a36d4940a5b8c40a989452d2304fc958ff3f354e7a"
	require.True(t, h.IsValid(hashed))

	hashed2 := "00762ae383fbb727af3c7a36d4940a5b8c40a989452d2304fc958ff3f354e7a"
	require.False(t, h.IsValid(hashed2))

}

func TestProof(t *testing.T) {

	c := make(chan bool)

	h := NewHasher(difficulty)

	el := domain.Challenge{
		Key:   "aaa",
		Check: 5,
	}
	solved := h.Proof(c, &el)
	require.Equal(t, 578, el.Nonce)
	require.True(t, solved)

	el2 := domain.Challenge{
		Key:   "bbb",
		Check: 25,
	}
	solved = h.Proof(c, &el2)
	require.Equal(t, 960, el2.Nonce)
	require.True(t, solved)

}
