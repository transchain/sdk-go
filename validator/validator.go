/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package validator

import (
	"regexp"
	"sync"

	"github.com/go-playground/validator/v10"
)

const fqidTagName = "fqid"

var fqidRegexp = regexp.MustCompile("^[a-z]{6}-[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")

var instance *validator.Validate
var once sync.Once

// Get returns a single custom validator instance.
func Get() *validator.Validate {
	once.Do(func() {
		instance = validator.New()
		// Define the goValidate handler
		_ = instance.RegisterValidation(fqidTagName, fqidHandler)
	})
	return instance
}

// fqidHandler validates that a string field contains a company bcid and a uuid.
// [6 lowercase alpha chars]-[uuid4].
func fqidHandler(fl validator.FieldLevel) bool {
	return fqidRegexp.MatchString(fl.Field().String())
}
