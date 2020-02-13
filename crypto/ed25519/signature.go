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

    "github.com/oasislabs/ed25519"
)

var (
    ErrBadSignatureSize = fmt.Errorf("bad ed25519 signature size")
)

// Signature is an ed25519 signature wrapper (64 bytes).
type Signature [ed25519.SignatureSize]byte

// String returns the base64 representation.
func (s Signature) String() string {
    return base64.StdEncoding.EncodeToString(s[:])
}

// MarshalJSON encodes the base64 value of an ed25519 signature.
func (s Signature) MarshalJSON() ([]byte, error) {
    return json.Marshal(s[:])
}

// UnmarshalJSON accepts a base64 value to load an ed25519 signature.
func (s *Signature) UnmarshalJSON(data []byte) error {
    var bytes []byte
    if err := json.Unmarshal(data, &bytes); err != nil {
        return err
    }
    if len(bytes) != ed25519.SignatureSize {
        return ErrBadSignatureSize
    }
    copy(s[:], bytes)
    return nil
}
