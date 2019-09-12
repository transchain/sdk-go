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
)

// Stopper interface defines the methods a concrete stopper must implement.
type Stopper interface {
    Stop()
    IsStopped() bool
}

// DefaultStopper provides a basic stopper.
type DefaultStopper struct {
    stopped bool
}

// DefaultStopper constructor.
func NewDefaultStopper() *DefaultStopper {
    return &DefaultStopper{
        stopped: false,
    }
}

// Stop sets the stopper in a stopped state.
func (c *DefaultStopper) Stop() {
    c.stopped = true
}

// IsStopped indicates if the Stop method has already been called.
func (c *DefaultStopper) IsStopped() bool {
    return c.stopped
}

// TrapSignal listens os signals SIGINT and SIGTERM in a goroutine.
// If trapped once, the stopper's Stop method is called to cleanup.
// If trapped twice, the program is interrupted.
func TrapSignal(stopper Stopper) {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        for sig := range c {
            if !stopper.IsStopped() {
                fmt.Println(fmt.Sprintf("captured %s, properly exiting...", sig.String()))
                stopper.Stop()
            } else {
                fmt.Println(fmt.Sprintf("captured %s, exiting forced...", sig.String()))
                os.Exit(130)
            }
        }
    }()
}
