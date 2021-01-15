package common

import "github.com/haozing/mztl/pkg/code"

type TLObjectHelper struct {
	newTLObjectFunc func() TLObject
}
type TLObject interface {
	Encode() []byte
	//EncodeToLayer(layer int) []byte
	Decode(dbuf *code.DecodeBuf) error
}
