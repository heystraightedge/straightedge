package hd

import (
	"github.com/heystraightedge/straightedge/crypto/keys/sr25519"
	substratebip39 "github.com/sikkatech/go-substrate-bip39"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/types"
)

var (
	// Sr25519 represents the Sr25519Type signature system.
	Sr25519 = sr25519Algo{}
)

type sr25519Algo struct {
}

func (s sr25519Algo) Name() hd.PubKeyType {
	return hd.Sr25519Type
}

// Derive derives and returns the sr25519 private key for the given seed and HD path.
func (s sr25519Algo) Derive() hd.DeriveFn {
	return func(mnemonic string, bip39Passphrase, hdPath string) ([]byte, error) {
		seed, err := substratebip39.SeedFromMnemonic(mnemonic, bip39Passphrase)
		return seed[:], err
	}
}

// Generate generates a sr25519 private key from the given bytes.
func (s sr25519Algo) Generate() hd.GenerateFn {
	return func(bz []byte) types.PrivKey {
		return sr25519.PrivKey{bz}
	}
}
