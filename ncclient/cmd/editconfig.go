// -*- coding: utf-8 -*-

package nccmd

import (
	nc "github.com/hiepon/go-netconf/netconf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type EditConfigCommand struct {
	Command
	Ope    string
	Commit bool
}

func (c *EditConfigCommand) SetFlags(cmd *cobra.Command) *cobra.Command {
	cmd.PersistentFlags().StringVarP(&c.Ope, "operation", "d", "replace", "default-operation(replace/none/merge).")
	cmd.PersistentFlags().BoolVarP(&c.Commit, "commit", "c", false, "Commit on success.")
	return c.Command.SetFlags(cmd)
}

func (c *EditConfigCommand) EditConfig(target string, input string) error {
	c.Command.Init()

	data, err := inputData(input)
	if err != nil {
		return err
	}

	client, err := c.Client()
	if err != nil {
		log.Errorf("Client initialize error. %s", err)
		return err
	}
	defer client.Close()

	if _, err := client.Exec(nc.MethodEditConfig(target, c.Ope, string(data))); err != nil {
		log.Errorf("Exec error. %s", err)
		return err
	}

	if c.Commit {
		if _, err := client.Exec(nc.MethodCommit()); err != nil {
			return err
		}
	}

	return nil
}

func EditConfigCmd() *cobra.Command {
	edit := EditConfigCommand{}
	c := edit.SetFlags(&cobra.Command{
		Use:   "edit-config [running/candidate] [filename or -]",
		Short: "edit-config command.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return edit.EditConfig(args[0], args[1])
		},
	})

	return c
}
