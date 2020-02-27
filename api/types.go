/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

// Client interface defines the methods a concrete client must implement.
type Client interface {
	Get(route string, headers map[string]string, queryValues map[string]string) (*RawResponse, error)
	Post(route string, body []byte, headers map[string]string, queryValues map[string]string) (*RawResponse, error)
	AddHeader(key string, value string)
	RemoveHeader(key string)
}

// Response is a fasthttp.Response wrapper.
type RawResponse struct {
	StatusCode int
	Body       []byte
}
