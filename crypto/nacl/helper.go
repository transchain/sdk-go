/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package nacl

import (
    "crypto/rand"

    "golang.org/x/crypto/nacl/box"
)

// GenerateNewPrivateKey generates a new x25519 private key.
func GenerateNewPrivateKey() PrivateKey {
    pubKey, privKey, err := box.GenerateKey(rand.Reader)
    if err != nil {
        panic(err)
    }
    privateKeyBytes := make([]byte, PrivateKeySize)
    copy(privateKeyBytes, privKey[:])
    copy(privateKeyBytes[32:], pubKey[:])
    return NewPrivateKey(privateKeyBytes)
}
