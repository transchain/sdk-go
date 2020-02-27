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
	ErrBadNonceSize = fmt.Errorf("bad nonce size")
)

const BoxNonceSize = 24

// BoxNonce is a Nacl box nonce wrapper (24 bytes)
type BoxNonce [BoxNonceSize]byte

// String returns the base64 representation.
func (bn BoxNonce) String() string {
	return base64.StdEncoding.EncodeToString(bn[:])
}

// MarshalJSON encodes the base64 value of a box nonce.
func (bn BoxNonce) MarshalJSON() ([]byte, error) {
	return json.Marshal(bn[:])
}

// UnmarshalJSON accepts a base64 value to load a box nonce.
func (bn *BoxNonce) UnmarshalJSON(data []byte) error {
	var bytes []byte
	if err := json.Unmarshal(data, &bytes); err != nil {
		return err
	}
	if len(bytes) != BoxNonceSize {
		return ErrBadNonceSize
	}
	copy(bn[:], bytes)
	return nil
}
