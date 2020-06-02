/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ed25519

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/oasisprotocol/ed25519"
)

var (
	ErrBadPublicKeySize         = fmt.Errorf("bad ed25519 public key size")
	ErrBadPublicKeyBase64Format = fmt.Errorf("bad ed25519 public key base64 format")
)

// PublicKey is an ed25519 public key wrapper (32 bytes).
type PublicKey [ed25519.PublicKeySize]byte

// PublicKey constructor.
func NewPublicKey(publicKeyBytes []byte) PublicKey {
	if len(publicKeyBytes) != ed25519.PublicKeySize {
		panic(ErrBadPublicKeySize)
	}
	var publicKey PublicKey
	copy(publicKey[:], publicKeyBytes[:])
	return publicKey
}

// PublicKey constructor from a base64 string.
func NewPublicKeyFromBase64(publicKeyBase64 string) PublicKey {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		panic(ErrBadPublicKeyBase64Format)
	}
	return NewPublicKey(publicKeyBytes)
}

// Verify indicates if a message and a signature match.
func (pk PublicKey) Verify(message []byte, signature Signature) bool {
	return ed25519.Verify(pk[:], message, signature[:])
}

// String returns the base64 representation.
func (pk PublicKey) String() string {
	return base64.StdEncoding.EncodeToString(pk[:])
}

// MarshalJSON encodes the base64 value of an ed25519 public key.
func (pk PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(pk[:])
}

// UnmarshalJSON accepts a base64 value to load an ed25519 public key.
func (pk *PublicKey) UnmarshalJSON(data []byte) error {
	var bytes []byte
	if err := json.Unmarshal(data, &bytes); err != nil {
		return err
	}
	if len(bytes) != ed25519.PublicKeySize {
		return ErrBadPublicKeySize
	}
	copy(pk[:], bytes)
	return nil
}
