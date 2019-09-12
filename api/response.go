/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

// Response is a fasthttp.Response wrapper.
type RawResponse struct {
    StatusCode int
    Body       []byte
}