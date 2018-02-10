// -*- coding: utf-8 -*-

package nccmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	nc "github.com/hiepon/go-netconf/netconf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Command struct {
	Host     string
	Username string
	Password string
	Verbose  bool
}

func (c *Command) SetFlags(cmd *cobra.Command) *cobra.Command {
	cmd.PersistentFlags().StringVarP(&c.Host, "host", "H", "localhost:830", "Host name.")
	cmd.PersistentFlags().StringVarP(&c.Username, "user", "U", os.Getenv("USER"), "Username.")
	cmd.PersistentFlags().StringVarP(&c.Password, "passwd", "P", "", "Password.")
	cmd.PersistentFlags().BoolVarP(&c.Verbose, "verbose", "v", false, "Show detail messages.")
	return cmd
}

func (c *Command) Init() {
	if c.Verbose {
		log.SetLevel(log.DebugLevel)
	}
	nc.SetLog(log.StandardLogger())
}

func (c *Command) Client() (session *nc.Session, err error) {
	config := nc.SSHConfigPassword(c.Username, c.Password)
	session, err = nc.DialSSH(c.Host, config)
	if session != nil {
		log.Debugf("SESSION: %d", session.SessionID)
		for index, capa := range session.ServerCapabilities {
			log.Debugf("CAPA[%d]: %s", index, capa)
		}
	}
	return
}

func outputRPCReply(path string, reply *nc.RPCReply) error {
	switch path {
	case "-", "":
		return writeRPCReply(os.Stdout, reply)
	default:
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		return writeRPCReply(f, reply)
	}
}

func writeRPCReply(w io.Writer, reply *nc.RPCReply) error {
	_, err := fmt.Fprint(w, reply.RawReply)
	return err
}

func inputData(path string) ([]byte, error) {
	switch path {
	case "-", "":
		return readData(os.Stdin)
	default:
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		return readData(f)
	}
}

func readData(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}
