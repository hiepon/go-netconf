// -*- coding: utf-8 -*-

package nccmd

import (
	nc "github.com/hiepon/go-netconf/netconf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type GetConfigCommand struct {
	Command
	Output string
}

func (c *GetConfigCommand) SetFlags(cmd *cobra.Command) *cobra.Command {
	cmd.PersistentFlags().StringVarP(&c.Output, "output", "o", "-", "Output filename.")
	return c.Command.SetFlags(cmd)
}

func (c *GetConfigCommand) GetConfig(source string) error {
	c.Command.Init()

	client, err := c.Client()
	if err != nil {
		log.Errorf("Client initialize error. %s", err)
		return err
	}
	defer client.Close()

	reply, err := client.Exec(nc.MethodGetConfig(source))
	if err != nil {
		log.Errorf("Exec error. %s", err)
		return err
	}

	return outputRPCReply(c.Output, reply)
}

func GetConfigCmd() *cobra.Command {
	get := GetConfigCommand{}
	c := get.SetFlags(&cobra.Command{
		Use:   "get-config [running/startup/candidate]",
		Short: "get-config command.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return get.GetConfig(args[0])
		},
	})

	return c
}
