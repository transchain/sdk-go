/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
	"math"
	"strconv"
)

const (
	PageParam           = "page"
	PerPageParam        = "per_page"
	DefaultPerPageParam = 10
)

// GetPaginationQueryParams returns the query params map to request a pagination.
func GetPaginationQueryParams(page int, txPerPage int) map[string]string {
	return map[string]string{
		PageParam:    strconv.Itoa(page),
		PerPageParam: strconv.Itoa(txPerPage),
	}
}

// GetNumPage returns the total number of pages for a given total number of items and their maximum number per page.
func GetNumPage(perPage, totalCount int) int {
	return int(math.Ceil(float64(totalCount) / float64(perPage)))
}

// GetPaginationOffsets returns the minimum and maximum offset of a specific page.
func GetPaginationOffsets(page, perPage, totalCount int) (int, int) {
	start := (page - 1) * perPage
	if start < 0 {
		start = 0
	}
	end := page * perPage
	if end > totalCount {
		end = totalCount
	}
	return start, end
}
