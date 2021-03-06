// Package api comment
// This file was generated by tars2go 1.1.4
// Generated from api.tars
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/protocol/tup"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8
var _ = unsafe.Pointer(nil)

//Apiusers struct
type Apiusers struct {
	s m.Servant
}

//Users_setSecureValueErrors is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiusers) Users_setSecureValueErrors(params *TLusers_setSecureValueErrors, _opt ...map[string]string) (ret Bool, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = params.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "users_setSecureValueErrors", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = ret.ReadBlock(_is, 0, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Users_setSecureValueErrorsWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiusers) Users_setSecureValueErrorsWithContext(tarsCtx context.Context, params *TLusers_setSecureValueErrors, _opt ...map[string]string) (ret Bool, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = params.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "users_setSecureValueErrors", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = ret.ReadBlock(_is, 0, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Users_setSecureValueErrorsOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiusers) Users_setSecureValueErrorsOneWayWithContext(tarsCtx context.Context, params *TLusers_setSecureValueErrors, _opt ...map[string]string) (ret Bool, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = params.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "users_setSecureValueErrors", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Users_getUsers is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiusers) Users_getUsers(params *TLusers_getUsers, _opt ...map[string]string) (ret []User, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = params.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "users_getUsers", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err, have, ty = _is.SkipToNoCheck(0, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}

		ret = make([]User, length)
		for i12, e12 := int32(0), length; i12 < e12; i12++ {

			err = ret[i12].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}

		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}

	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Users_getUsersWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiusers) Users_getUsersWithContext(tarsCtx context.Context, params *TLusers_getUsers, _opt ...map[string]string) (ret []User, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = params.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "users_getUsers", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err, have, ty = _is.SkipToNoCheck(0, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}

		ret = make([]User, length)
		for i13, e13 := int32(0), length; i13 < e13; i13++ {

			err = ret[i13].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}

		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}

	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Users_getUsersOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiusers) Users_getUsersOneWayWithContext(tarsCtx context.Context, params *TLusers_getUsers, _opt ...map[string]string) (ret []User, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = params.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "users_getUsers", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Users_getFullUser is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiusers) Users_getFullUser(params *TLusers_getFullUser, _opt ...map[string]string) (ret UserFull, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = params.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "users_getFullUser", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = ret.ReadBlock(_is, 0, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Users_getFullUserWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiusers) Users_getFullUserWithContext(tarsCtx context.Context, params *TLusers_getFullUser, _opt ...map[string]string) (ret UserFull, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = params.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "users_getFullUser", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = ret.ReadBlock(_is, 0, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Users_getFullUserOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiusers) Users_getFullUserOneWayWithContext(tarsCtx context.Context, params *TLusers_getFullUser, _opt ...map[string]string) (ret UserFull, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = params.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "users_getFullUser", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *Apiusers) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *Apiusers) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol sets the protocol for the servant.
func (_obj *Apiusers) TarsSetProtocol(p m.Protocol) {
	_obj.s.TarsSetProtocol(p)
}

//AddServant adds servant  for the service.
func (_obj *Apiusers) AddServant(imp _impApiusers, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *Apiusers) AddServantWithContext(imp _impApiusersWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impApiusers interface {
	Users_setSecureValueErrors(params *TLusers_setSecureValueErrors) (ret Bool, err error)
	Users_getUsers(params *TLusers_getUsers) (ret []User, err error)
	Users_getFullUser(params *TLusers_getFullUser) (ret UserFull, err error)
}
type _impApiusersWithContext interface {
	Users_setSecureValueErrors(tarsCtx context.Context, params *TLusers_setSecureValueErrors) (ret Bool, err error)
	Users_getUsers(tarsCtx context.Context, params *TLusers_getUsers) (ret []User, err error)
	Users_getFullUser(tarsCtx context.Context, params *TLusers_getFullUser) (ret UserFull, err error)
}

// Dispatch is used to call the server side implemnet for the method defined in the tars file. _withContext shows using context or not.
func (_obj *Apiusers) Dispatch(tarsCtx context.Context, _val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, _withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	_os := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "users_setSecureValueErrors":
		var params TLusers_setSecureValueErrors

		if tarsReq.IVersion == basef.TARSVERSION {

			err = params.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("params", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = params.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["params"])
				if err = json.Unmarshal([]byte(_jsonStr_), &params); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		var _funRet_ Bool
		if _withContext == false {
			_imp := _val.(_impApiusers)
			_funRet_, err = _imp.Users_setSecureValueErrors(&params)
		} else {
			_imp := _val.(_impApiusersWithContext)
			_funRet_, err = _imp.Users_setSecureValueErrors(tarsCtx, &params)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _funRet_.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _funRet_.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "users_getUsers":
		var params TLusers_getUsers

		if tarsReq.IVersion == basef.TARSVERSION {

			err = params.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("params", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = params.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["params"])
				if err = json.Unmarshal([]byte(_jsonStr_), &params); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		var _funRet_ []User
		if _withContext == false {
			_imp := _val.(_impApiusers)
			_funRet_, err = _imp.Users_getUsers(&params)
		} else {
			_imp := _val.(_impApiusersWithContext)
			_funRet_, err = _imp.Users_getUsers(tarsCtx, &params)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.WriteHead(codec.LIST, 0)
			if err != nil {
				return err
			}

			err = _os.Write_int32(int32(len(_funRet_)), 0)
			if err != nil {
				return err
			}

			for _, v := range _funRet_ {

				err = v.WriteBlock(_os, 0)
				if err != nil {
					return err
				}

			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.WriteHead(codec.LIST, 0)
			if err != nil {
				return err
			}

			err = _os.Write_int32(int32(len(_funRet_)), 0)
			if err != nil {
				return err
			}

			for _, v := range _funRet_ {

				err = v.WriteBlock(_os, 0)
				if err != nil {
					return err
				}

			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "users_getFullUser":
		var params TLusers_getFullUser

		if tarsReq.IVersion == basef.TARSVERSION {

			err = params.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("params", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = params.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["params"])
				if err = json.Unmarshal([]byte(_jsonStr_), &params); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		var _funRet_ UserFull
		if _withContext == false {
			_imp := _val.(_impApiusers)
			_funRet_, err = _imp.Users_getFullUser(&params)
		} else {
			_imp := _val.(_impApiusersWithContext)
			_funRet_, err = _imp.Users_getFullUser(tarsCtx, &params)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _funRet_.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _funRet_.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(tarsCtx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(tarsCtx)
	if ok && c != nil {
		_context = c
	}
	*tarsResp = requestf.ResponsePacket{
		IVersion:     tarsReq.IVersion,
		CPacketType:  0,
		IRequestId:   tarsReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}

	_ = _is
	_ = _os
	_ = length
	_ = have
	_ = ty
	return nil
}
