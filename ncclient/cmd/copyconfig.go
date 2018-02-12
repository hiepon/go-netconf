// -*- coding: utf-8 -*-

package nccmd

import (
	nc "github.com/hiepon/go-netconf/netconf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CopyConfigCommand struct {
	Command
}

func (c *CopyConfigCommand) CopyConfig(source string, target string) error {
	c.Command.Init()

	client, err := c.Client()
	if err != nil {
		log.Errorf("Client initialize error. %s", err)
		return err
	}
	defer client.Close()

	if _, err := client.Exec(nc.MethodCopyConfig(target, source)); err != nil {
		log.Errorf("Exec error. %s", err)
		return err
	}

	return nil
}

func CopyConfigCmd() *cobra.Command {
	copy := CopyConfigCommand{}
	c := copy.SetFlags(&cobra.Command{
		Use:   "copy-config [source] [target]",
		Short: "copy-config command.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return copy.CopyConfig(args[0], args[1])
		},
	})

	return c
}
