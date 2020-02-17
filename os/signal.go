/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package os

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/transchain/sdk-go/log"
)

// Stopper interface defines the methods a concrete stopper must implement.
type Stopper interface {
    Stop() error
    IsRunning() bool
}

// TrapSignal listens os signals SIGINT and SIGTERM in a goroutine.
// If trapped once, the stopper's Stop method is called to cleanup.
// If trapped twice, the program is interrupted.
func TrapSignal(logger log.Logger, stopper Stopper) {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        for sig := range c {
            if stopper.IsRunning() {
                go func() {
                    logger.Info(fmt.Sprintf("captured [%s] signal, properly exiting...", sig.String()))
                    if err := stopper.Stop(); err != nil {
                        logger.Error(fmt.Sprintf("unable to properly exit: %s", err.Error()))
                        os.Exit(130)
                    }
                } ()
            } else {
                logger.Info(fmt.Sprintf("captured [%s] signal, exiting forced...", sig.String()))
                os.Exit(130)
            }
        }
    }()
}
