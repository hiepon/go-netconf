// -*- coding: utf-8 -*-

package ncenc

import (
	"encoding/xml"
)

type RpcReplyData struct {
	XMLName xml.Name `xml:"data"`
	Xml     []byte   `xml:",innerxml"`
}

type RpcReplyOk struct {
	XMLName xml.Name `xml:"ok"`
	Xml     []byte   `xml:",innerxml"`
}

type RpcReplyErrorInfo struct {
	XMLName xml.Name `xml:"error-info"`
	Content []byte   `xml:",innerxml"`
}

type RpcReplyError struct {
	XMLName  xml.Name          `xml:"rpc-error"`
	Severity string            `xml:"error-severity"`
	Path     string            `xml:"error-path"`
	Info     RpcReplyErrorInfo `xml:"error-info"`
}

type RpcReply struct {
	XMLName xml.Name       `xml:"rpc-reply"`
	Data    *RpcReplyData  `xml:"data"`
	Ok      *RpcReplyOk    `xml:"ok"`
	Error   *RpcReplyError `xml:"rpc-error"`
}

func (r *RpcReply) IsOk() bool {
	return r.Ok != nil
}

func DecodeRpcReply(b []byte, reply *RpcReply) error {
	return xml.Unmarshal(b, reply)
}
