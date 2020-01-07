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

    "gopkg.in/go-playground/validator.v9"
)

var txidRegexp = regexp.MustCompile("^[a-z]{6}-[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")

var instance *validator.Validate
var once sync.Once

// Get returns a single custom validator instance.
func Get() *validator.Validate {
    once.Do(func() {
        instance = validator.New()
        // Define the goValidate handler
        _ = instance.RegisterValidation("txid", txidHandler)
    })
    return instance
}

// txidHandler validates that a string field is a txid field.
// [6 lowercase alpha chars]-[uuidv4].
func txidHandler(fl validator.FieldLevel) bool {
    return txidRegexp.MatchString(fl.Field().String())
}
