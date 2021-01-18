package api

type TLObjectHelper struct {
	newTLObjectFunc func() TLObject
}
type TLObject interface {
	Encode() []byte
	//EncodeToLayer(layer int) []byte
	Decode(dbuf *DecodeBuf) error
}
