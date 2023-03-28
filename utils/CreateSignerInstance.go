package utils

import (
	"crypto"
	"crypto/ecdsa"
)

type Signer struct {
	privateKey *ecdsa.PrivateKey
}

func (s *Signer) Public() crypto.PublicKey {
	return s.privateKey.Public()
}

func (s *Signer) Sign()
