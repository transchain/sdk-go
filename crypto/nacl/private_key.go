/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package nacl

import (
    "crypto/rand"
    "encoding/base64"
    "io"

    "golang.org/x/crypto/nacl/box"
)

const PrivateKeySize = 64

// PrivateKey is an x25519 private key wrapper (64 bytes).
type PrivateKey [PrivateKeySize]byte

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

// String returns the base64 representation.
func (pk PrivateKey) String() string {
    return base64.StdEncoding.EncodeToString(pk[:])
}

// Seal encrypts a plain text message decipherable afterwards by the recipient private key.
func (pk PrivateKey) Seal(message []byte, recipientPublicKey PublicKey) ([]byte, BoxNonce, error) {
    var nonce [24]byte
    if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
        return nil, BoxNonce{}, err
    }
    var privKey [32]byte
    copy(privKey[:], pk[:32])
    encryptedMessage := box.Seal(nil, message, &nonce, (*[32]byte)(&recipientPublicKey), &privKey)
    return encryptedMessage, nonce, nil
}

// Open decrypts an encrypted message with the appropriate sender information.
func (pk PrivateKey) Open(encryptedMessage []byte, nonce BoxNonce, senderPublicKey PublicKey) ([]byte, bool) {
    var privKey [32]byte
    copy(privKey[:], pk[:32])
    return box.Open(nil, encryptedMessage, (*[24]byte)(&nonce), (*[32]byte)(&senderPublicKey), &privKey)
}

// GetPublicKey returns the underlying x25519 public key.
func (pk PrivateKey) GetPublicKey() PublicKey {
    return NewPublicKey(pk[32:])
}
