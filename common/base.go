package common

const (
	MTPROTO_VERSION = 2
)

// Transport类型，不支持UDP
const (
	TRANSPORT_TCP  = 1 // TCP
	TRANSPORT_HTTP = 2 // HTTP
	TRANSPORT_UDP  = 3 // UDP, @benqi: 未发现有支持UDP的客户端
)

type Codec interface {
	Receive() (interface{}, error)
	Send(interface{}) error
	Close() error
}
