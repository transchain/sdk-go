/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
    "math"
)

// GetNumPage returns the total number of pages for a given total number of items and their maximum number per page.
func GetNumPage(perPage, totalCount int) int {
    return int(math.Ceil(float64(totalCount) / float64(perPage)))
}