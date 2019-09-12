/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package context

import (
    "github.com/transchain/sdk-go/os"

    "github.com/transchain/sdk-go/log"
)

// DefaultContext provides a basic stoppable context with a logger.
type DefaultContext struct {
    os.Stopper
    Logger log.Logger
}

// DefaultContext constructor.
func NewDefaultContext(stopper os.Stopper, logger log.Logger) *DefaultContext {
    return &DefaultContext{
        Stopper: stopper,
        Logger:  logger,
    }
}
