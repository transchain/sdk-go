/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package os

import (
    "os"
)

// PathExists indicates if a path exists.
func PathExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

// RemovePathIfExists removes a directory or a file if it exists.
func RemovePathIfExists(path string) error {
    if PathExists(path) {
        return os.RemoveAll(path)
    }
    return nil
}
