/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package log

// Logger interface defines the methods a concrete logger must implement.
type Logger interface {
    Debug(message string, keyValues ...interface{})
    Info(message string, keyValues ...interface{})
    Error(message string, keyValues ...interface{})
    Warn(message string, keyValues ...interface{})
    With(keyValues ...interface{}) Logger
}