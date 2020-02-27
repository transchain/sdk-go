/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package json

import (
	"encoding/json"
)

// MarshalWrapper wraps an interface value and its corresponding string type.
// It is useful to convert an interface field to unmarshal it back afterwards.
type MarshalWrapper struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

// UnmarshalWrapper wraps a json raw message and its corresponding string type.
// It is useful to delay the unmarshal step once a concrete type can be determined by its string type.
type UnmarshalWrapper struct {
	Value json.RawMessage `json:"value"`
	Type  string          `json:"type"`
}

// MarshalAndSortJSON sorts alphabetically the json representation of an interface and returns its marshaled value.
func MarshalAndSortJSON(jsonValue interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(jsonValue)
	if err != nil {
		return nil, err
	}
	// The json package sorts by structure's fields order.
	// Marshal and Unmarshal back to an interface does the trick.
	var sortedJsonValue interface{}
	err = json.Unmarshal(jsonBytes, &sortedJsonValue)
	if err != nil {
		return nil, err
	}
	return json.Marshal(sortedJsonValue)
}

// MustMarshalAndSortJSON panics if MarshalAndSortJSON fails.
func MustMarshalAndSortJSON(jsonValue interface{}) []byte {
	jsonBytes, err := MarshalAndSortJSON(jsonValue)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}
