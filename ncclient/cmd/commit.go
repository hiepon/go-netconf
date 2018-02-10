// -*- coding: utf-8 -*-

package nccmd

import (
	nc "github.com/hiepon/go-netconf/netconf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CommitCommand struct {
	Command
}

func (c *CommitCommand) Commit() error {
	c.Command.Init()

	client, err := c.Client()
	if err != nil {
		log.Errorf("Client initialize error. %s", err)
		return err
	}
	defer client.Close()

	if _, err := client.Exec(nc.MethodCommit()); err != nil {
		log.Errorf("Exec error. %s", err)
		return err
	}

	return nil
}

func CommitCmd() *cobra.Command {
	commit := CommitCommand{}
	c := commit.SetFlags(&cobra.Command{
		Use:   "commit",
		Short: "commit command.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return commit.Commit()
		},
	})

	return c
}
