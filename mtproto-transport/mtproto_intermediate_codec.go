package mtproto

import (
	"encoding/binary"
	"fmt"
	"github.com/haozing/mztl/common"
	"io"
)

type MTProtoIntermediateCodec struct {
	conn *common.BufferedConn
}

func NewMTProtoIntermediateCodec(conn *common.BufferedConn) *MTProtoIntermediateCodec {
	return &MTProtoIntermediateCodec{
		conn: conn,
	}
}

func (c *MTProtoIntermediateCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size = int(binary.LittleEndian.Uint32(b) << 2)

	// glog.Info("first_byte: ", hex.EncodeToString(b[:1]))
	// needAck := bool(b[0] >> 7 == 1)
	// _ = needAck

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			fmt.Println("ReadFull2 error: ", err)
			return nil, err
		}
		left -= n
	}
	//if size > 10240 {
	//	glog.Info("ReadFull2: ", hex.EncodeToString(buf[:256]))
	//}

	// TODO(@benqi): process report ack and quickack
	// 截断QuickAck消息，客户端有问题
	if size == 4 {
		fmt.Println("Server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := common.NewMTPRawMessage(authKeyId, 0, common.TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *MTProtoIntermediateCodec) Send(msg interface{}) error {
	message, ok := msg.(*common.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		fmt.Println(err)
		return err
	}

	b := message.Encode()

	sb := make([]byte, 4)
	// minus padding
	size := len(b) / 4

	//if size < 127 {
	//	sb = []byte{byte(size)}
	//} else {
	binary.LittleEndian.PutUint32(sb, uint32(size))
	//}

	b = append(sb, b...)
	_, err := c.conn.Write(b)

	if err != nil {
		fmt.Println("Send msg error: %s", err)
	}

	return err
}

func (c *MTProtoIntermediateCodec) Close() error {
	return c.conn.Close()
}
