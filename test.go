package main

import (
	"github.com/golang/glog"
	"github.com/haozing/mztl/transport"
	"io"
)

func main() {
	var conn io.ReadWriter
	_, err := transport.NewMTProtoProxy().NewCodec(conn)
	if err != nil {
		glog.Error(err)
		return
	}
}
