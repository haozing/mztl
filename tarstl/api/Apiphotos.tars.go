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

//Apiphotos struct
type Apiphotos struct {
	s m.Servant
}

//Photos_getUserPhotos is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_getUserPhotos(params *TLphotos_getUserPhotos, _opt ...map[string]string) (ret Photos_Photos, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "photos_getUserPhotos", _os.ToBytes(), _status, _context, _resp)
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

//Photos_getUserPhotosWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_getUserPhotosWithContext(tarsCtx context.Context, params *TLphotos_getUserPhotos, _opt ...map[string]string) (ret Photos_Photos, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "photos_getUserPhotos", _os.ToBytes(), _status, _context, _resp)
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

//Photos_getUserPhotosOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_getUserPhotosOneWayWithContext(tarsCtx context.Context, params *TLphotos_getUserPhotos, _opt ...map[string]string) (ret Photos_Photos, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 1, "photos_getUserPhotos", _os.ToBytes(), _status, _context, _resp)
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

//Photos_updateProfilePhoto is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_updateProfilePhoto(params *TLphotos_updateProfilePhoto, _opt ...map[string]string) (ret Photos_Photo, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "photos_updateProfilePhoto", _os.ToBytes(), _status, _context, _resp)
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

//Photos_updateProfilePhotoWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_updateProfilePhotoWithContext(tarsCtx context.Context, params *TLphotos_updateProfilePhoto, _opt ...map[string]string) (ret Photos_Photo, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "photos_updateProfilePhoto", _os.ToBytes(), _status, _context, _resp)
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

//Photos_updateProfilePhotoOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_updateProfilePhotoOneWayWithContext(tarsCtx context.Context, params *TLphotos_updateProfilePhoto, _opt ...map[string]string) (ret Photos_Photo, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 1, "photos_updateProfilePhoto", _os.ToBytes(), _status, _context, _resp)
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

//Photos_uploadProfilePhoto is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_uploadProfilePhoto(params *TLphotos_uploadProfilePhoto, _opt ...map[string]string) (ret Photos_Photo, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "photos_uploadProfilePhoto", _os.ToBytes(), _status, _context, _resp)
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

//Photos_uploadProfilePhotoWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_uploadProfilePhotoWithContext(tarsCtx context.Context, params *TLphotos_uploadProfilePhoto, _opt ...map[string]string) (ret Photos_Photo, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "photos_uploadProfilePhoto", _os.ToBytes(), _status, _context, _resp)
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

//Photos_uploadProfilePhotoOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_uploadProfilePhotoOneWayWithContext(tarsCtx context.Context, params *TLphotos_uploadProfilePhoto, _opt ...map[string]string) (ret Photos_Photo, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 1, "photos_uploadProfilePhoto", _os.ToBytes(), _status, _context, _resp)
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

//Photos_deletePhotos is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_deletePhotos(params *TLphotos_deletePhotos, _opt ...map[string]string) (ret []int64, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "photos_deletePhotos", _os.ToBytes(), _status, _context, _resp)
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

		ret = make([]int64, length)
		for i6, e6 := int32(0), length; i6 < e6; i6++ {

			err = _is.Read_int64(&ret[i6], 0, false)
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

//Photos_deletePhotosWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_deletePhotosWithContext(tarsCtx context.Context, params *TLphotos_deletePhotos, _opt ...map[string]string) (ret []int64, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "photos_deletePhotos", _os.ToBytes(), _status, _context, _resp)
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

		ret = make([]int64, length)
		for i7, e7 := int32(0), length; i7 < e7; i7++ {

			err = _is.Read_int64(&ret[i7], 0, false)
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

//Photos_deletePhotosOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Apiphotos) Photos_deletePhotosOneWayWithContext(tarsCtx context.Context, params *TLphotos_deletePhotos, _opt ...map[string]string) (ret []int64, err error) {

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

	err = _obj.s.Tars_invoke(tarsCtx, 1, "photos_deletePhotos", _os.ToBytes(), _status, _context, _resp)
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
func (_obj *Apiphotos) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *Apiphotos) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol sets the protocol for the servant.
func (_obj *Apiphotos) TarsSetProtocol(p m.Protocol) {
	_obj.s.TarsSetProtocol(p)
}

//AddServant adds servant  for the service.
func (_obj *Apiphotos) AddServant(imp _impApiphotos, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *Apiphotos) AddServantWithContext(imp _impApiphotosWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impApiphotos interface {
	Photos_getUserPhotos(params *TLphotos_getUserPhotos) (ret Photos_Photos, err error)
	Photos_updateProfilePhoto(params *TLphotos_updateProfilePhoto) (ret Photos_Photo, err error)
	Photos_uploadProfilePhoto(params *TLphotos_uploadProfilePhoto) (ret Photos_Photo, err error)
	Photos_deletePhotos(params *TLphotos_deletePhotos) (ret []int64, err error)
}
type _impApiphotosWithContext interface {
	Photos_getUserPhotos(tarsCtx context.Context, params *TLphotos_getUserPhotos) (ret Photos_Photos, err error)
	Photos_updateProfilePhoto(tarsCtx context.Context, params *TLphotos_updateProfilePhoto) (ret Photos_Photo, err error)
	Photos_uploadProfilePhoto(tarsCtx context.Context, params *TLphotos_uploadProfilePhoto) (ret Photos_Photo, err error)
	Photos_deletePhotos(tarsCtx context.Context, params *TLphotos_deletePhotos) (ret []int64, err error)
}

// Dispatch is used to call the server side implemnet for the method defined in the tars file. _withContext shows using context or not.
func (_obj *Apiphotos) Dispatch(tarsCtx context.Context, _val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, _withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	_os := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "photos_getUserPhotos":
		var params TLphotos_getUserPhotos

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

		var _funRet_ Photos_Photos
		if _withContext == false {
			_imp := _val.(_impApiphotos)
			_funRet_, err = _imp.Photos_getUserPhotos(&params)
		} else {
			_imp := _val.(_impApiphotosWithContext)
			_funRet_, err = _imp.Photos_getUserPhotos(tarsCtx, &params)
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
	case "photos_updateProfilePhoto":
		var params TLphotos_updateProfilePhoto

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

		var _funRet_ Photos_Photo
		if _withContext == false {
			_imp := _val.(_impApiphotos)
			_funRet_, err = _imp.Photos_updateProfilePhoto(&params)
		} else {
			_imp := _val.(_impApiphotosWithContext)
			_funRet_, err = _imp.Photos_updateProfilePhoto(tarsCtx, &params)
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
	case "photos_uploadProfilePhoto":
		var params TLphotos_uploadProfilePhoto

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

		var _funRet_ Photos_Photo
		if _withContext == false {
			_imp := _val.(_impApiphotos)
			_funRet_, err = _imp.Photos_uploadProfilePhoto(&params)
		} else {
			_imp := _val.(_impApiphotosWithContext)
			_funRet_, err = _imp.Photos_uploadProfilePhoto(tarsCtx, &params)
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
	case "photos_deletePhotos":
		var params TLphotos_deletePhotos

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

		var _funRet_ []int64
		if _withContext == false {
			_imp := _val.(_impApiphotos)
			_funRet_, err = _imp.Photos_deletePhotos(&params)
		} else {
			_imp := _val.(_impApiphotosWithContext)
			_funRet_, err = _imp.Photos_deletePhotos(tarsCtx, &params)
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

				err = _os.Write_int64(v, 0)
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

				err = _os.Write_int64(v, 0)
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
