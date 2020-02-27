/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	VersionCmdName  = "version"
	VersionCmdShort = "Show version info"
)

// GetVersionCmd returns a basic cobra command to show version information.
func GetVersionCmd(version string, commit string) *cobra.Command {
	return &cobra.Command{
		Use:   VersionCmdName,
		Short: VersionCmdShort,
		Run: func(*cobra.Command, []string) {
			if commit != "" {
				fmt.Println(fmt.Sprintf("%s (%s)", version, commit))
			} else {
				fmt.Println(version)
			}
		},
	}
}
