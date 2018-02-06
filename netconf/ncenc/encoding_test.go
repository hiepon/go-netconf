// -*- coding: utf-8 -*-

package ncenc

import (
	"testing"
)

func TestDecodeRpcReplyData(t *testing.T) {
	b := []byte(
		"<rpc-reply message-id=\"a5857ca1-a3b8-4dec-82f8-6dd54afc995f\" " +
			"xmlns=\"urn:ietf:params:xml:ns:netconf:base:1.0\">" +
			"<data xmlns=\"urn:ietf:params:xml:ns:netconf:base:1.0\">" +
			"<test>test</test>" +
			"</data>" +
			"</rpc-reply>",
	)
	var reply RpcReply
	if err := DecodeRpcReply(b, &reply); err != nil {
		t.Errorf("DecodeReply error. %s", err)
	}

	data := reply.Data
	if data == nil {
		t.Errorf("DecodeReply error. reply.data=%v", data)
	}
	if s := string(data.Xml); s != "<test>test</test>" {
		t.Errorf("DecodeReply error. reply.data.xml=%s", s)
	}
	if ok := reply.Ok; ok != nil {
		t.Errorf("DecodeReply error. reply.ok=%v", ok)
	}
	if err := reply.Error; err != nil {
		t.Errorf("DecodeReply error. reply.error=%v", err)
	}
}

func TestDecodeRpcReplyOk(t *testing.T) {
	b := []byte(
		"<rpc-reply message-id=\"a5857ca1-a3b8-4dec-82f8-6dd54afc995f\" " +
			"xmlns=\"urn:ietf:params:xml:ns:netconf:base:1.0\">" +
			"<ok/>" +
			"</rpc-reply>",
	)
	var reply RpcReply
	if err := DecodeRpcReply(b, &reply); err != nil {
		t.Errorf("DecodeReply error. %s", err)
	}

	if data := reply.Data; data != nil {
		t.Errorf("DecodeReply error. reply.data=%v", data)
	}
	if ok := reply.Ok; ok == nil {
		t.Errorf("DecodeReply error. reply.ok=%v", ok)
	}
	if err := reply.Error; err != nil {
		t.Errorf("DecodeReply error. reply.error=%v", err)
	}
}

func TestDecodeRpcReplyError(t *testing.T) {
	b := []byte(
		"<rpc-reply message-id=\"a5857ca1-a3b8-4dec-82f8-6dd54afc995f\" " +
			"xmlns=\"urn:ietf:params:xml:ns:netconf:base:1.0\">" +
			"<rpc-error>" +
			"<error-severity>test-error-severity</error-severity>" +
			"<error-path>test-error-path</error-path>" +
			"<error-message>test-error-message</error-message>" +
			"<error-info>test-error-info</error-info>" +
			"</rpc-error>" +
			"</rpc-reply>",
	)
	var reply RpcReply
	if err := DecodeRpcReply(b, &reply); err != nil {
		t.Errorf("DecodeReply error. %s", err)
	}

	if data := reply.Data; data != nil {
		t.Errorf("DecodeReply error. reply.data=%v", data)
	}
	if ok := reply.Ok; ok != nil {
		t.Errorf("DecodeReply error. reply.ok=%v", ok)
	}
	err := reply.Error
	if err == nil {
		t.Errorf("DecodeReply error. reply.error=%v", err)
	}
	if v := err.Severity; v != "test-error-severity" {
		t.Errorf("DecodeReply error. reply.error.severity=%s", v)
	}
	if v := err.Path; v != "test-error-path" {
		t.Errorf("DecodeReply error. reply.error.path=%s", v)
	}
	if v := string(err.Info.Content); v != "test-error-info" {
		t.Errorf("DecodeReply error. reply.error.info=%s", v)
	}
}
