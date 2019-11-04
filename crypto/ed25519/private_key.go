/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ed25519

import (
    "encoding/base64"

    "github.com/oasislabs/ed25519"
)

// PrivateKey is an ed25519 private key wrapper (64 bytes).
type PrivateKey [ed25519.PrivateKeySize]byte

// PrivateKey constructor.
func NewPrivateKey(privateKeyBytes []byte) PrivateKey {
    var privateKey PrivateKey
    copy(privateKey[:], privateKeyBytes[:])
    return privateKey
}

// PrivateKey constructor from a base64 string.
func NewPrivateKeyFromBase64(privateKeyBase64 string) (PrivateKey, error) {
    privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
    if err != nil {
        return PrivateKey{}, err
    }
    return NewPrivateKey(privateKeyBytes), nil
}

// Sign accepts a message and returns its corresponding ed25519 signature.
func (pk PrivateKey) Sign(message []byte) Signature {
    var signature Signature
    copy(signature[:], ed25519.Sign(pk[:], message)[:])
    return signature
}

// GetPublicKey returns the underlying ed25519 public key.
func (pk PrivateKey) GetPublicKey() PublicKey {
    return NewPublicKey(pk[32:])
}

// String returns the base64 representation.
func (pk PrivateKey) String() string {
    return base64.StdEncoding.EncodeToString(pk[:])
}
