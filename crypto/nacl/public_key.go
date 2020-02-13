/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package nacl

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
)

var (
    ErrBadPublicKeySize         = fmt.Errorf("bad x25519 public key size")
    ErrBadPublicKeyBase64Format = fmt.Errorf("bad x25519 public key base64 format")
)

const PublicKeySize = 32

// PublicKey is an x25519 public key wrapper (32 bytes).
type PublicKey [PublicKeySize]byte

// PublicKey constructor.
func NewPublicKey(publicKeyBytes []byte) PublicKey {
    if len(publicKeyBytes) != PublicKeySize {
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

// String returns the base64 representation.
func (pk PublicKey) String() string {
    return base64.StdEncoding.EncodeToString(pk[:])
}

// MarshalJSON encodes the base64 value of an x25519 public key.
func (pk PublicKey) MarshalJSON() ([]byte, error) {
    return json.Marshal(pk[:])
}

// UnmarshalJSON accepts a base64 value to load an x25519 public key.
func (pk *PublicKey) UnmarshalJSON(data []byte) error {
    var bytes []byte
    if err := json.Unmarshal(data, &bytes); err != nil {
        return err
    }
    if len(bytes) != PublicKeySize {
        return ErrBadPublicKeySize
    }
    copy(pk[:], bytes)
    return nil
}
