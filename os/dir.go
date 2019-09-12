/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package os

import (
    "os"
    "path/filepath"
)

// EnsureDir creates a directory if it does not exist.
func EnsureDir(dirPath string, perm os.FileMode) error {
    if dirPath != "" && !PathExists(dirPath) {
        return os.MkdirAll(dirPath, perm)
    }
    return nil
}

// EndureFileDir creates a file's directory if it does not exist.
func EnsureFileDir(filePath string, perm os.FileMode) error {
    dirPath, _ := filepath.Split(filePath)
    return EnsureDir(dirPath, perm)
}