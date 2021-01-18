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
	"errors"
	"fmt"
	"math"
	"math/big"
)

type DecodeBuf struct {
	Buf  []byte
	off  int
	Size int
	Err  error
}

func NewDecodeBuf(b []byte) *DecodeBuf {
	return &DecodeBuf{b, 0, len(b), nil}
}

func (m *DecodeBuf) GetError() error {
	return m.Err
}

func (m *DecodeBuf) Long() int64 {
	if m.Err != nil {
		return 0
	}
	if m.off+8 > m.Size {
		m.Err = errors.New("DecodeLong")
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.Buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Double() float64 {
	if m.Err != nil {
		return 0
	}
	if m.off+8 > m.Size {
		m.Err = errors.New("DecodeDouble")
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.Buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Int() int32 {
	if m.Err != nil {
		return 0
	}
	if m.off+4 > m.Size {
		m.Err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.Buf[m.off : m.off+4])
	m.off += 4
	return int32(x)
}

func (m *DecodeBuf) UInt() uint32 {
	if m.Err != nil {
		return 0
	}
	if m.off+4 > m.Size {
		m.Err = errors.New("DecodeUInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.Buf[m.off : m.off+4])
	m.off += 4
	return x
}

func (m *DecodeBuf) Bytes(size int) []byte {
	if m.Err != nil {
		return nil
	}
	if m.off+size > m.Size {
		m.Err = errors.New("DecodeBytes")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.Buf[m.off:m.off+size])
	m.off += size
	return x
}

func (m *DecodeBuf) StringBytes() []byte {
	if m.Err != nil {
		return nil
	}
	var size, padding int

	if m.off+1 > m.Size {
		m.Err = errors.New("DecodeStringBytes")
		return nil
	}
	size = int(m.Buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.Size {
			m.Err = errors.New("DecodeStringBytes")
			return nil
		}
		size = int(m.Buf[m.off]) | int(m.Buf[m.off+1])<<8 | int(m.Buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.Size {
		m.Err = errors.New("DecodeStringBytes: Wrong size")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.Buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.Size {
		m.Err = errors.New("DecodeStringBytes: Wrong padding")
		return nil
	}
	m.off += padding

	return x
}

func (m *DecodeBuf) String() string {
	b := m.StringBytes()
	if m.Err != nil {
		return ""
	}
	x := string(b)
	return x
}

func (m *DecodeBuf) BigInt() *big.Int {
	b := m.StringBytes()
	if m.Err != nil {
		return nil
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	return x
}

func (m *DecodeBuf) VectorInt() []int32 {
	constructor := m.Int()
	if m.Err != nil {
		return nil
	}
	if constructor != 481674261 {
		m.Err = fmt.Errorf("DecodeVectorInt: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.Err != nil {
		return nil
	}
	if size < 0 {
		m.Err = errors.New("DecodeVectorInt: Wrong size")
		return nil
	}
	x := make([]int32, size)
	i := int32(0)
	for i < size {
		y := m.Int()
		if m.Err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorLong() []int64 {
	constructor := m.Int()
	if m.Err != nil {
		return nil
	}
	if constructor != 481674261 {
		m.Err = fmt.Errorf("DecodeVectorLong: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.Err != nil {
		return nil
	}
	if size < 0 {
		m.Err = errors.New("DecodeVectorLong: Wrong size")
		return nil
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := m.Long()
		if m.Err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorString() []string {
	constructor := m.Int()
	if m.Err != nil {
		return nil
	}
	if constructor != 481674261 {
		m.Err = fmt.Errorf("DecodeVectorString: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.Err != nil {
		return nil
	}
	if size < 0 {
		m.Err = errors.New("DecodeVectorString: Wrong size")
		return nil
	}
	x := make([]string, size)
	i := int32(0)
	for i < size {
		y := m.String()
		if m.Err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) Bool() bool {
	constructor := m.Int()
	if m.Err != nil {
		return false
	}
	switch constructor {
	case -1720552011:
		return true
	case -1132882121:
		return false
	}
	return false
}

func (m *DecodeBuf) Object() (r TLObject) {

	classID := m.Int()

	if m.Err != nil {
		fmt.Println("classID m.Int():%s", m.Err)
		return nil
	}

	re, ok := ApiRegisters[classID]
	fmt.Println("decode classid %s", classID)
	if !ok {
		fmt.Println("registers2[classID] !ok")
		return nil
	}

	if r == nil {
		fmt.Println("re.newTLObjectFunc() !ok")
		m.Err = fmt.Errorf("can't find registed classId: 0x%x", uint32(classID))

		return nil
	}

	err := re.Decode(m)

	if err != nil {
		fmt.Println("r.(TLObject).Decode(m) !ok")
		m.Err = err
	}
	return
}
