/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ed25519

import (
	"bytes"
	"crypto/rand"

	"github.com/oasislabs/ed25519"
)

// PublicKeysAreEqual checks if two public keys are equal.
func PublicKeysAreEqual(key1 PublicKey, key2 PublicKey) bool {
	return bytes.Equal(key1[:], key2[:])
}

// GenerateNewPrivateKey generates a new ed25519 private key.
func GenerateNewPrivateKey() PrivateKey {
	_, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	return NewPrivateKey(privKey)
}
