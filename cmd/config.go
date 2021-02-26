// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/clivern/goenv/core/module"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the goenv application.",
	Run: func(cmd *cobra.Command, args []string) {
		golang := module.NewGolangEnvironment(HOME)

		err := golang.Configure()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Goenv configured successfully!")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}