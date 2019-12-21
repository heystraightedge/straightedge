package main

import (
	"crypto/sha512"
	"errors"

	bip39 "github.com/cosmos/go-bip39"

	"golang.org/x/crypto/pbkdf2"

	schnorrkel "github.com/ChainSafe/go-schnorrkel"
)

// MiniSecretFromMnemonic returns a go-schnorrkel MiniSecretKey from a bip39 mnemonic
func MiniSecretFromMnemonic(mnemonic string, password string) (*schnorrkel.MiniSecretKey, error) {
	seed, err := SeedFromMnemonic(mnemonic, password)
	return schnorrkel.NewMiniSecretKey(seed), err
}

// SeedFromMnemonic returns a 64-byte seed from a bip39 mnemonic
func SeedFromMnemonic(mnemonic string, password string) ([64]byte, error) {
	entropy, err := bip39.MnemonicToByteArray(mnemonic)
	if err != nil {
		return [64]byte{}, err
	}

	if len(entropy) < 16 || len(entropy) > 32 || len(entropy)%4 != 0 {
		return [64]byte{}, errors.New("invalid entropy")
	}

	bz := pbkdf2.Key([]byte(mnemonic), []byte("mnemonic"+password), 2048, 64, sha512.New)
	var bzArr [64]byte
	copy(bzArr[:], bz[:64])

	return bzArr, nil
}
