// -*- coding: utf-8 -*-

package nccmd

import (
	"os"

	"github.com/spf13/cobra"
)

func RootCmd(name string) *cobra.Command {
	var showCompletion bool

	c := &cobra.Command{
		Use:   name,
		Short: "netconf shell,",
		Run: func(cmd *cobra.Command, args []string) {
			if showCompletion {
				cmd.GenBashCompletion(os.Stdout)
			} else {
				cmd.Usage()
			}
		},
	}

	c.PersistentFlags().BoolVar(&showCompletion, "show-completion", false, "Show bash-comnpletion")

	c.AddCommand(
		GetCmd(),
		GetConfigCmd(),
	)
	return c
}
