package mtproto

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/haozing/mztl/common"
	"io"
)

// https://core.telegram.org/mtproto#tcp-transport
//
// There is an abridged version of the same protocol:
// 同一协议有一个删节版本
// if the client sends 0xef as the first byte (**important:** only prior to the very first data packet),
// then packet length is encoded by a single byte (0x01..0x7e = data length divided by 4;
// or 0x7f followed by 3 length bytes (little endian) divided by 4) followed
// by the data themselves (sequence number and CRC32 not added).
// In this case, server responses look the same (the server does not send 0xefas the first byte).
// 如果客户机发送0xef作为第一个字节（**important:**仅在第一个数据包之前）,则数据包长度由单个字节（0x01..0x7e=数据长度除以4；或0x7f后跟3个长度字节（小尾数）除以4）编码后跟数据本身（序列号和CRC32未添加）。
// 在这种情况下，服务器响应看起来是一样的（服务器不发送0xefas第一个字节）。
type MTProtoAbridgedCodec struct {
	conn *common.BufferedConn
}

func NewMTProtoAbridgedCodec(conn *common.BufferedConn) *MTProtoAbridgedCodec {
	return &MTProtoAbridgedCodec{
		conn: conn,
	}
}

func (c *MTProtoAbridgedCodec) Receive() (interface{}, error) {

	var size int
	var n int
	var err error

	b := make([]byte, 1)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	// glog.Info("first_byte: ", hex.EncodeToString(b[:1]))
	needAck := bool(b[0]>>7 == 1)
	_ = needAck

	b[0] = b[0] & 0x7f
	// glog.Info("first_byte2: ", hex.EncodeToString(b[:1]))

	if b[0] < 0x7f {
		size = int(b[0]) << 2
		if size == 0 {
			return nil, nil
		}
	} else {

		b2 := make([]byte, 3)
		n, err = io.ReadFull(c.conn, b2)
		if err != nil {
			return nil, err
		}
		size = (int(b2[0]) | int(b2[1])<<8 | int(b2[2])<<16) << 2

	}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {

			return nil, err
		}
		left -= n
	}
	if size > 10240 {
		fmt.Println("ReadFull2: ", hex.EncodeToString(buf[:256]))
	}

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

func (c *MTProtoAbridgedCodec) Send(msg interface{}) error {
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

	if size < 127 {
		sb = []byte{byte(size)}
	} else {
		binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
	}

	b = append(sb, b...)
	_, err := c.conn.Write(b)

	if err != nil {
		fmt.Println("Send msg error: %s", err)
	}

	return err
}

func (c *MTProtoAbridgedCodec) Close() error {
	return c.conn.Close()
}
