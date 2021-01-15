// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package mtproto

import (
	"encoding/binary"
	"fmt"
	"io"
	"mztl/common"
)

type MTProtoFullCodec struct {
	conn *common.BufferedConn
}

func NewMTProtoFullCodec(conn *common.BufferedConn) *MTProtoFullCodec {
	return &MTProtoFullCodec{
		conn: conn,
	}
}

func (c *MTProtoFullCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size = int(binary.LittleEndian.Uint32(b) << 2)
	// Check bufLen
	if size < 12 || size%4 != 0 {
		err = fmt.Errorf("invalid len: %d", size)
		return nil, err
	}

	//buf := make([]byte, size - 4)
	//n, err = io.ReadFull(c.conn, buf)
	//if err != nil {
	//	return nil, err
	//}

	left := size
	buf := make([]byte, size-4)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			fmt.Println("ReadFull2 error: ", err)
			return nil, err
		}
		left -= n
	}

	seqNum := binary.LittleEndian.Uint32(buf[:4])
	// TODO(@benqi): check seqNum, save last seq_num
	_ = seqNum

	crc32 := binary.LittleEndian.Uint32(buf[len(buf)-4:])
	// TODO(@benqi): check crc32
	_ = crc32

	authKeyId := int64(binary.LittleEndian.Uint64(buf[4:]))
	message := common.NewMTPRawMessage(authKeyId, 0, common.TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *MTProtoFullCodec) Send(msg interface{}) error {
	message, ok := msg.(*common.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		fmt.Println(err)
		return err
	}

	b := message.Encode()

	sb := make([]byte, 8)
	// minus padding
	size := len(b) / 4

	//if size < 127 {
	//	sb = []byte{byte(size)}
	//} else {

	binary.LittleEndian.PutUint32(sb, uint32(size))
	// TODO(@benqi): gen seq_num
	var seqNum uint32 = 0
	binary.LittleEndian.PutUint32(sb[4:], seqNum)
	//}
	b = append(sb, b...)
	var crc32Buf []byte = make([]byte, 4)
	var crc32 uint32 = 0
	binary.LittleEndian.PutUint32(crc32Buf, crc32)
	b = append(sb, crc32Buf...)

	_, err := c.conn.Write(b)
	if err != nil {
		fmt.Println("Send msg error: %s", err)
	}

	return err
}

func (c *MTProtoFullCodec) Close() error {
	return c.conn.Close()
}
