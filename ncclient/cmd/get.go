// -*- coding: utf-8 -*-

package nccmd

import (
	nc "github.com/hiepon/go-netconf/netconf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type GetCommand struct {
	Command
	Output string
}

func (c *GetCommand) SetFlags(cmd *cobra.Command) *cobra.Command {
	cmd.PersistentFlags().StringVarP(&c.Output, "output", "o", "-", "Output filename.")
	return c.Command.SetFlags(cmd)
}

func (c *GetCommand) Get() error {
	c.Command.Init()

	client, err := c.Client()
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	defer client.Close()

	reply, err := client.Exec(nc.MethodGet())
	if err != nil {
		log.Errorf("%s", err)
		return err
	}

	return outputRPCReply(c.Output, reply)
}

func GetCmd() *cobra.Command {
	get := GetCommand{}
	c := get.SetFlags(&cobra.Command{
		Use:   "get",
		Short: "get command.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return get.Get()
		},
	})

	return c
}
