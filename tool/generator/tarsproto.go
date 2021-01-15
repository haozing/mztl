package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func ToTarsProto() {

	Pbproto.WriteString("module apitars{\n")
	Goproto.WriteString("package api\n")
	Goproto.WriteString("import(\n")
	Goproto.WriteString("  \"fmt\"\n")
	Goproto.WriteString("  \"github.com/haozing/mztl/pkg/code\"\n")
	Goproto.WriteString("import )\n")
	//import (
	//	"fmt"
	//	"github.com/haozing/mztl/pkg/code"
	//)

	Baseproto.WriteString("package api\n")
	Baseproto.WriteString("import \"github.com/haozing/mztl/pkg/code\"\n")
	//整理参数
	TarsFieldConversion()

	//结构
	TarsConstructorsToProto()
	//接口
	TarsMethodsToProto()
	Pbproto.WriteString("};\n")
	//参数与小类型解码

	TarsCodec()
}

//RPC接口划分
type rpcmethod struct {
	FieldConParams
	TypeName string
}

var rpcmethods = make(map[string][]rpcmethod)

func TarsFieldConversion() {
	//大类型下的字段
	for k, v := range PbConstructors.Type {
		//字段对应
		FieldMap := make(map[string]string)
		//去掉 Vector
		if k == "Vector t" {
			continue
		}
		pdid := 1
		typename := TarsTypeFormat(k)
		var DeduplicationString = []string{}
		var RepeatedFields = []string{}
		for _, vvv := range PbConstructors.Params[k] {
			//去掉 # flag
			if vvv.Type == "#" {
				continue
			}
			//去掉flags.0?
			vvvType := TarsReplaceFlags(vvv.Type)
			//Vector<Message>换vector<Message>
			vvvType = TarsVectorToRepeated(vvvType)

			//去重复
			dc := TarsDeduplication(DeduplicationString, vvvType+vvv.Name)
			if !dc {
				continue
			}

			DeduplicationString = append(DeduplicationString, vvvType+vvv.Name)

			//重复字段添加id
			vvvName := vvv.Name
			rf := TarsDeduplication(RepeatedFields, vvv.Name)

			if !rf {
				vvvName = vvv.Name + "_" + strconv.Itoa(pdid)
			}
			RepeatedFields = append(RepeatedFields, vvv.Name)

			vvvName = TarskeywordSubstitution(vvvName)
			FieldMap[vvvType+vvv.Name] = vvvName
			FieldParamsType := FieldParamsType{
				Name: vvvName,
				Type: TarsTypeFormat(vvvType),
			}
			FC.Type[typename] = append(FC.Type[typename], FieldParamsType)
			pdid++
		}
		//大类型下面小类型
		for _, mv := range v {
			mvPredicate := TarsTypeFormat(mv.Predicate)
			FC.TypeConstructors[typename] = append(FC.TypeConstructors[typename], mvPredicate)
			FieldConParam := FieldConParams{
				Name: mvPredicate,
			}
			//小类型字段
			isflag := false
			for _, xv := range mv.Params {
				//去掉 # flag
				if xv.Type == "#" {
					isflag = true
					continue
				}
				//去掉flags.0?
				vvvType := TarsReplaceFlags(xv.Type)
				//Vector<Message>换repeated Message
				vvvType = TarsVectorToRepeated(vvvType)

				//flagid
				flagid, _ := GetTarsFlagsId(xv.Type)
				ParamsType := FieldParamsType{
					Name:   FieldMap[vvvType+xv.Name],
					Type:   TarsTypeFormat(vvvType),
					FlagId: flagid,
				}
				FieldConParam.Params = append(FieldConParam.Params, ParamsType)
			}
			FieldConParam.Flag = isflag
			FieldConParam.Id = mv.ID
			FC.ConstructorsParams[typename] = append(FC.ConstructorsParams[typename], FieldConParam)
		}

	}

	//接口
	for PbMk, PbMv := range PbMethods.Type {

		//找Vector<的
		//message Vector_DialogFilterSuggested {
		//    repeated DialogFilterSuggested datas = 1;
		//}

		typename := TarsTypeFormat(PbMk)
		//接口参数
		for _, mv := range PbMv {
			mvMethod := TarsTypeFormat(mv.Method)
			FieldConParam := FieldConParams{
				Name: mvMethod,
			}
			//小类型字段
			isflag := false
			for _, xv := range mv.Params {
				//去掉 # flag
				if xv.Type == "#" {
					isflag = true
					continue
				}
				//去掉flags.0?
				vvvType := TarsReplaceFlags(xv.Type)
				//Vector<Message>换repeated Message
				vvvType = TarsVectorToRepeated(vvvType)
				//flagid
				flagid, _ := GetTarsFlagsId(xv.Type)
				//替换
				ParamsType := FieldParamsType{
					Name:   TarskeywordSubstitution(TarsTypeFormat(xv.Name)),
					Type:   TarsTypeFormat(vvvType),
					FlagId: flagid,
				}
				FieldConParam.Params = append(FieldConParam.Params, ParamsType)
			}

			FieldConParam.Flag = isflag
			FieldConParam.Id = mv.ID
			rpcmv := strings.Split(mv.Method, ".")
			rpcmm := mv.Method
			if len(rpcmv) > 1 {
				rpcmm = rpcmv[0]
				rpcst := rpcmethod{
					TypeName:       mv.Type,
					FieldConParams: FieldConParam,
				}
				rpcmethods[rpcmm] = append(rpcmethods[rpcmm], rpcst)

			}
			FC.MethodParams[typename] = append(FC.MethodParams[typename], FieldConParam)
		}
	}

}
func TarsConstructorsToProto() {
	for k, v := range FC.TypeConstructors {
		//注释
		Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
		Pbproto.WriteString("// Constructors:" + k + " \n")
		Pbproto.WriteString("// " + k + " <--\n")
		for _, vv := range v {
			Pbproto.WriteString("// + TL" + vv + "\n")
		}
		Pbproto.WriteString("//\n")

		//打类型
		/*		type Photo struct {
				Constructor          int32                  `protobuf:"varint,1,opt,name=constructor,proto3" json:"constructor,omitempty"`
				Data                 *ChatBannedRights_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
				XXX_NoUnkeyedLiteral struct{}               `json:"-"`
				XXX_unrecognized     []byte                 `json:"-"`
				XXX_sizecache        int32                  `json:"-"`
			}*/
		Pbproto.WriteString("// Constructors:Type:" + k + " \n")
		Pbproto.WriteString("struct " + k + " {\n")
		Pbproto.WriteString("	0 optional int constructor; \n")
		Pbproto.WriteString("	1 require  " + k + "_Data data ;\n")
		Pbproto.WriteString("};\n")

		//类型数据
		/*	type ChatBannedRights_Data struct {
			ViewMessages         bool */
		Pbproto.WriteString("// Constructors:Type:Data:" + k + " \n")
		Pbproto.WriteString("struct " + k + "_Data {\n")
		for kkk, vvv := range FC.Type[k] {
			Pbproto.WriteString("	" + strconv.Itoa(kkk) + "	optional " + vvv.Type + " " + vvv.Name + "; \n")
		}
		Pbproto.WriteString("};\n")
		for _, vv := range v {
			Pbproto.WriteString("// Constructors:Predicate:" + vv + " \n")
			Pbproto.WriteString("struct TL" + vv + " {\n")
			Pbproto.WriteString("	0	require " + k + "_Data data ;\n")
			Pbproto.WriteString("};\n")

			//func NewTLInputEncryptedChat() *TLInputEncryptedChat {
			//	return &TLInputEncryptedChat{Data2: &InputEncryptedChat_Data{}}
			//}

		}

		for _, cvv := range FC.ConstructorsParams[k] {

			cname := cvv.Name
			for _, ccvv := range cvv.Params {
				Goproto.WriteString("func (m *TL" + cname + ")  Set" + ccvv.Name + "(v " + TarsJudgeType(ccvv.Type) + "){m.Data." + strings.Title(ccvv.Name) + "=v}\n")
				Goproto.WriteString("func (m *TL" + cname + ")  Get" + ccvv.Name + "() " + TarsJudgeType(ccvv.Type) + "{return m.Data." + strings.Title(ccvv.Name) + "}\n")

			}
		}

	}
}
func TarsMethodsToProto() {
	//接口参数
	for _, mv := range FC.MethodParams {
		for _, mmv := range mv {
			Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
			Pbproto.WriteString("// Methods:Params:" + mmv.Name + " \n")
			Pbproto.WriteString("//\n")
			Pbproto.WriteString("struct TL" + mmv.Name + " {\n")
			for kkk, mvv := range mmv.Params {
				Pbproto.WriteString("	" + strconv.Itoa(kkk) + "	optional " + mvv.Type + " " + mvv.Name + ";\n")
			}
			Pbproto.WriteString("};\n")
			for _, mvv := range mmv.Params {
				Goproto.WriteString("func (m *TL" + mmv.Name + ")  Set" + mvv.Name + "(v " + TarsJudgeType(mvv.Type) + "){m." + strings.Title(mvv.Name) + "=v}\n")
				Goproto.WriteString("func (m *TL" + mmv.Name + ")  Get" + mvv.Name + "() " + TarsJudgeType(mvv.Type) + "{return m." + strings.Title(mvv.Name) + "}\n")

			}
		}

	}

	for rpck, rpcv := range rpcmethods {
		//接口rpc
		//	    interface ApiService
		//        {
		//            vector<CodeSettings> auth_sendCode(vector<CodeSettings> sendCode, bool allow_app_hash);
		//
		//        };
		Pbproto.WriteString("interface Api" + rpck + " {\n")
		for _, mmv := range rpcv {

			Pbproto.WriteString("	" + TarsTypeFormat(mmv.TypeName) + " " + mmv.Name + "(TL" + mmv.Name + " params);\n")

		}
		Pbproto.WriteString("};\n")
	}

}
func TarsCodec() {
	for _, mv := range FC.MethodParams {

		//func NewTLMessagesGetMessages() *TLMessagesGetMessages {
		//	return &TLMessagesGetMessages{}
		//}
		for _, mmv := range mv {
			Goproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
			Goproto.WriteString("// MethodCodec:New:" + mmv.Name + " \n")
			Goproto.WriteString("//\n")
			Goproto.WriteString("func NewTL" + mmv.Name + "() *TL" + mmv.Name + " {\n")
			Goproto.WriteString("return &TL" + mmv.Name + "{}\n")
			Goproto.WriteString("}\n")
			TarsCodecEncode(mmv)
			TarsCodecDecode(mmv)
		}
	}
	for _, mv := range FC.ConstructorsParams {

		//func NewTLMessagesGetMessages() *TLMessagesGetMessages {
		//	return &TLMessagesGetMessages{}
		//}
		for _, mmv := range mv {
			Goproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
			Goproto.WriteString("// ConstructorsCodec:New:" + mmv.Name + " \n")
			Goproto.WriteString("//\n")
			Goproto.WriteString("func NewTL" + mmv.Name + "() *TL" + mmv.Name + " {\n")
			Goproto.WriteString("return &TL" + mmv.Name + "{}\n")
			Goproto.WriteString("}\n")
			TarsCodecEncode(mmv)
			TarsCodecDecode(mmv)
		}
	}

	//大类型
	for dk, dv := range FC.ConstructorsParams {
		Goproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
		Goproto.WriteString("// ConstructorsCodec:Type:" + dk + " \n")
		Goproto.WriteString("//\n")
		//func (m *Bool) Encode() []byte {
		//	switch m.GetConstructor() {
		//	case TLConstructor_CRC32_BoolFalse:
		//		t := m.To_BoolFalse()
		//		return t.Encode()
		//	case TLConstructor_CRC32_BoolTrue:
		//		t := m.To_BoolTrue()
		//		return t.Encode()
		//	default:
		//		return nil
		//	}
		//
		//}
		Goproto.WriteString("func (m " + strings.Title(dk) + ") Encode() []byte {\n")
		Goproto.WriteString("    switch m.Constructor {\n")
		for _, dvv := range dv {
			Goproto.WriteString("    case " + dvv.Id + ":\n")
			Goproto.WriteString("        t := m.To_" + dvv.Name + "()\n")
			Goproto.WriteString("        return t.Encode()\n")
		}

		Goproto.WriteString("	default:\n")
		Goproto.WriteString("		return nil\n")
		Goproto.WriteString("	}\n")
		Goproto.WriteString("}\n")

		Goproto.WriteString("func (m " + strings.Title(dk) + ") Decode(dbuf *DecodeBuf) error {\n")
		Goproto.WriteString("    m.Constructor = dbuf.Int()\n")
		Goproto.WriteString("    switch m.Constructor {\n")
		for _, dvv := range dv {
			//case TLConstructor_CRC32_InputCheckPasswordEmpty:
			//		m2 := &TLInputCheckPasswordEmpty{Data2: &InputCheckPasswordSRP_Data{}}
			//		m2.Decode(dbuf)
			//		m.Data2 = m2.Data2
			Goproto.WriteString("    case " + dvv.Id + ":\n")
			Goproto.WriteString("    	m2 := &TL" + dvv.Name + "{Data: " + strings.Title(dk) + "_Data{}}\n")
			Goproto.WriteString("        m2.Decode(dbuf)\n")
			Goproto.WriteString("        m.Data = m2.Data\n")
		}

		Goproto.WriteString("	default:\n")
		Goproto.WriteString("		return  fmt.Errorf(\"Invalid constructorId: %d\", m.Constructor)\n")
		Goproto.WriteString("	}\n")
		Goproto.WriteString("	return dbuf.Err\n")

		Goproto.WriteString("}\n")

		//func (m *Bool) To_BoolFalse() *TLBoolFalse {
		//	return &TLBoolFalse{
		//		Data2: m.Data2,
		//	}
		//}
		for _, dvv := range dv {
			Goproto.WriteString("func (m " + strings.Title(dk) + ") To_" + dvv.Name + "() *TL" + dvv.Name + " {\n")
			Goproto.WriteString("    return &TL" + dvv.Name + "{\n")
			Goproto.WriteString("        Data: m.Data,\n")
			Goproto.WriteString("    }\n")
			Goproto.WriteString("}\n")
		}

	}

	//Baseproto

	Baseproto.WriteString("var ApiRegisters = map[int32]common.TLObject{\n")
	for _, mv := range FC.MethodParams {

		//func NewTLMessagesGetMessages() *TLMessagesGetMessages {
		//	return &TLMessagesGetMessages{}
		//}
		for _, mmv := range mv {
			//	33373783:  NewTLchannels_exportMessageLink(),

			Baseproto.WriteString(mmv.Id + ":  NewTL" + mmv.Name + "(),\n")

		}
	}
	for _, mv := range FC.ConstructorsParams {

		//func NewTLMessagesGetMessages() *TLMessagesGetMessages {
		//	return &TLMessagesGetMessages{}
		//}
		for _, mmv := range mv {
			Baseproto.WriteString(mmv.Id + ":  NewTL" + mmv.Name + "(),\n")
		}
	}
	Baseproto.WriteString("}\n")

}
func TarsCodecDecode(fcp FieldConParams) {
	//func (m *TLMessagesGetMessages) Encode() []byte {
	//	x := NewEncodeBuf(512)
	//	x.Int(int32(TLConstructor_CRC32_MessagesGetMessages))
	//	x.Int(int32(TLConstructor_CRC32_vector))
	//	x.Int(int32(len(m.Id)))
	//	for _, v := range m.Id {
	//		x.Buf = append(x.Buf, (*v).Encode()...)
	//	}
	//	return x.Buf
	//}
	Goproto.WriteString("// CodecDecode:" + fcp.Name + " \n")
	Goproto.WriteString("func (m *TL" + fcp.Name + ") Decode(dbuf *code.DecodeBuf) error {\n")
	//flags
	if fcp.Flag {
		//var flags uint32 = 0
		Goproto.WriteString("    flags := dbuf.UInt()\n")
		Goproto.WriteString("    _ = flags\n")
	}

	//params
	for kkk, mvv := range fcp.Params {
		if mvv.FlagId != "" {

			//	if m.Entities != nil {
			//		x.Int(int32(TLConstructor_CRC32_vector))
			//		x.Int(int32(len(m.Entities)))
			//		for _, v := range m.Entities {
			//			x.Buf = append(x.Buf, (*v).Encode()...)
			//		}
			//	}

			switch mvv.Type {
			case "double":
				//if (flags & (1 << 0)) != 0 {
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				//m.RequestedInfoId = dbuf.String()
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.Double())\n")
			case "bool":
				//if (flags & (1 << 0)) != 0 {
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				//m.RequestedInfoId = dbuf.String()
				Goproto.WriteString("    m.Set" + mvv.Name + "(true)\n")
			case "int":
				//if (flags & (1 << 0)) != 0 {
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				//m.RequestedInfoId = dbuf.String()
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.Int())\n")
			case "long":
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.Long())\n")
			case "vector<long>":
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.VectorLong())\n")
			case "vector<int>":
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.VectorInt())\n")
			case "vector<string>":
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.VectorString())\n")
			case "string":
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.String())\n")
			case "vector<unsigned byte>":
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.StringBytes())\n")
			case "vector<vector<unsigned byte>>":
				Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
				Goproto.WriteString("        c" + strconv.Itoa(kkk) + " := dbuf.Int()\n")
				Goproto.WriteString("        if c" + strconv.Itoa(kkk) + " != 481674261 {\n")
				Goproto.WriteString("       	return dbuf.Err\n")
				Goproto.WriteString("        }\n")
				Goproto.WriteString("        l" + strconv.Itoa(kkk) + " := dbuf.Int()\n")
				Goproto.WriteString("        v" + strconv.Itoa(kkk) + " := make([]byte, l" + strconv.Itoa(kkk) + ")\n")
				Goproto.WriteString("        for i := int32(0); i < l" + strconv.Itoa(kkk) + "; i++ {\n")
				Goproto.WriteString("            v" + strconv.Itoa(kkk) + "[i].Bytes(dbuf)\n")
				Goproto.WriteString("        }\n")
				Goproto.WriteString("    m.Set" + mvv.Name + "(v" + strconv.Itoa(kkk) + ")\n")
			default:
				r := mvv.Type[:3]
				if r == "vec" {
					Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
					Goproto.WriteString("        c" + strconv.Itoa(kkk) + " := dbuf.Int()\n")
					Goproto.WriteString("        if c" + strconv.Itoa(kkk) + " != 481674261 {\n")
					Goproto.WriteString("       	return dbuf.Err\n")
					Goproto.WriteString("        }\n")
					Goproto.WriteString("        l" + strconv.Itoa(kkk) + " := dbuf.Int()\n")
					Goproto.WriteString("        v" + strconv.Itoa(kkk) + " := make(" + TarsJudgeType(mvv.Type) + ", l" + strconv.Itoa(kkk) + ")\n")
					Goproto.WriteString("        for i := int32(0); i < l" + strconv.Itoa(kkk) + "; i++ {\n")
					Goproto.WriteString("        v" + strconv.Itoa(kkk) + "[i] = " + TarsJudgeType(mvv.Type)[2:] + "{}\n")
					Goproto.WriteString("            v" + strconv.Itoa(kkk) + "[i].Decode(dbuf)\n")
					Goproto.WriteString("        }\n")
					Goproto.WriteString("    m.Set" + mvv.Name + "(v" + strconv.Itoa(kkk) + ")\n")

				} else {
					Goproto.WriteString("    if (flags & (1 << " + strconv.Itoa(kkk) + ")) != 0 {\n")
					//	m5 := &InputPaymentCredentials{}
					//	m5.Decode(dbuf)
					//	m.Credentials = m5

					Goproto.WriteString("    m" + strconv.Itoa(kkk) + " := &" + TarsJudgeType(mvv.Type) + "{}\n")
					Goproto.WriteString("    m" + strconv.Itoa(kkk) + ".Decode(dbuf)\n")
					Goproto.WriteString("    m.Set" + mvv.Name + "(m" + strconv.Itoa(kkk) + ")\n")
				}
			}
			Goproto.WriteString("    }\n")
		} else {
			switch mvv.Type {
			case "double":
				//m.RequestedInfoId = dbuf.String()
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.Double())\n")
			case "bool":
				//m.RequestedInfoId = dbuf.String()
				Goproto.WriteString("    m.Set" + mvv.Name + "(true)\n")
			case "int":
				//if (flags & (1 << 0)) != 0 {

				//m.RequestedInfoId = dbuf.String()
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.Int())\n")
			case "long":
				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.Long())\n")
			case "vector<long>":

				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.VectorLong())\n")
			case "vector<int>":

				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.VectorInt())\n")
			case "vector<string>":

				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.VectorString())\n")

			case "string":

				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.String())\n")
			case "vector<unsigned byte>":

				Goproto.WriteString("    m.Set" + mvv.Name + "(dbuf.StringBytes())\n")
			case "vector<vector<unsigned byte>>":

				Goproto.WriteString("        c" + strconv.Itoa(kkk) + " := dbuf.Int()\n")
				Goproto.WriteString("        if c" + strconv.Itoa(kkk) + " != 481674261 {\n")
				Goproto.WriteString("       	return dbuf.Err\n")
				Goproto.WriteString("        }\n")
				Goproto.WriteString("        l" + strconv.Itoa(kkk) + " := dbuf.Int()\n")
				Goproto.WriteString("        v" + strconv.Itoa(kkk) + " := make([]byte, l" + strconv.Itoa(kkk) + ")\n")
				Goproto.WriteString("        for i := int32(0); i < l" + strconv.Itoa(kkk) + "; i++ {\n")
				Goproto.WriteString("            v" + strconv.Itoa(kkk) + "[i].Bytes(dbuf)\n")
				Goproto.WriteString("        }\n")
				Goproto.WriteString("    m.Set" + mvv.Name + "(v" + strconv.Itoa(kkk) + ")\n")
			default:
				r := mvv.Type[:3]
				if r == "vec" {

					Goproto.WriteString("        c" + strconv.Itoa(kkk) + " := dbuf.Int()\n")
					Goproto.WriteString("        if c" + strconv.Itoa(kkk) + " != 481674261 {\n")
					Goproto.WriteString("       	return dbuf.Err\n")
					Goproto.WriteString("        }\n")
					Goproto.WriteString("        l" + strconv.Itoa(kkk) + " := dbuf.Int()\n")
					Goproto.WriteString("        v" + strconv.Itoa(kkk) + " := make(" + TarsJudgeType(mvv.Type) + ", l" + strconv.Itoa(kkk) + ")\n")
					Goproto.WriteString("        for i := int32(0); i < l" + strconv.Itoa(kkk) + "; i++ {\n")
					Goproto.WriteString("        v" + strconv.Itoa(kkk) + "[i] = " + TarsJudgeType(mvv.Type)[2:] + "{}\n")
					Goproto.WriteString("            v" + strconv.Itoa(kkk) + "[i].Decode(dbuf)\n")
					Goproto.WriteString("        }\n")
					Goproto.WriteString("    m.Set" + mvv.Name + "(v" + strconv.Itoa(kkk) + ")\n")

				} else {
					//	m5 := &InputPaymentCredentials{}
					//	m5.Decode(dbuf)
					//	m.Credentials = m5

					Goproto.WriteString("    m" + strconv.Itoa(kkk) + " := &" + TarsJudgeType(mvv.Type) + "{}\n")
					Goproto.WriteString("    m" + strconv.Itoa(kkk) + ".Decode(dbuf)\n")
					Goproto.WriteString("    m.Set" + mvv.Name + "(m" + strconv.Itoa(kkk) + ")\n")
				}
			}
		}

	}
	Goproto.WriteString("    return dbuf.Err\n")
	Goproto.WriteString("}\n")
}
func TarsCodecEncode(fcp FieldConParams) {
	//func (m *TLMessagesGetMessages) Encode() []byte {
	//	x := NewEncodeBuf(512)
	//	x.Int(int32(TLConstructor_CRC32_MessagesGetMessages))
	//	x.Int(int32(TLConstructor_CRC32_vector))
	//	x.Int(int32(len(m.Id)))
	//	for _, v := range m.Id {
	//		x.Buf = append(x.Buf, (*v).Encode()...)
	//	}
	//	return x.Buf
	//}
	Goproto.WriteString("// CodecEncode:" + fcp.Name + " \n")
	Goproto.WriteString("func (m *TL" + fcp.Name + ") Encode() []byte {\n")
	Goproto.WriteString("    x := code.NewEncodeBuf(512)\n")
	Goproto.WriteString("    x.Int(" + fcp.Id + ")\n")
	//flags
	if fcp.Flag {
		//var flags uint32 = 0
		Goproto.WriteString("    var flags uint32 = 0\n")
	}
	for kkk, mvv := range fcp.Params {
		if mvv.FlagId != "" {
			switch mvv.Type {
			case "bool":
				Goproto.WriteString("    if m.Get" + mvv.Name + "()!=false {\n")

			case "int", "Long", "vector<int>", "vector<long>", "vector<string>":
				Goproto.WriteString("    if m.Get" + mvv.Name + "()!=0 {\n")
			case "bytes":
				Goproto.WriteString("    if m.Get" + mvv.Name + "()!=nil {\n")
			case "string":
				Goproto.WriteString("    if m.Get" + mvv.Name + "()!=\"\" {\n")
				//Vector<int>
			default:
				r := mvv.Type[:3]
				if r == "vec" {
					Goproto.WriteString("    if len(m.Get" + mvv.Name + "())>0 {\n")
				} else {
					Goproto.WriteString("    if m.Get" + mvv.Name + "().Constructor!=0 {\n")
				}
			}

			Goproto.WriteString("        flags |= " + mvv.FlagId + " << " + strconv.Itoa(kkk) + "\n")
			Goproto.WriteString("    }\n")

		}
	}

	if fcp.Flag {
		Goproto.WriteString("    x.UInt(flags)\n")
	}
	//params
	for _, mvv := range fcp.Params {
		if mvv.FlagId != "" {

			//	if m.Entities != nil {
			//		x.Int(int32(TLConstructor_CRC32_vector))
			//		x.Int(int32(len(m.Entities)))
			//		for _, v := range m.Entities {
			//			x.Buf = append(x.Buf, (*v).Encode()...)
			//		}
			//	}

			//bool类型屏蔽
			if mvv.Type == "bool" {
				continue
			}

			switch mvv.Type {
			case "bool":
				//ss ="Get"+ s + "()!=false"
				Goproto.WriteString("    if m.Get" + mvv.Name + "() != false {\n")
				Goproto.WriteString("    x.Bytes(m.Get" + mvv.Name + "())\n")
			case "int":
				//ss ="Get"+ s + "()!=false"
				Goproto.WriteString("    if m.Get" + mvv.Name + "() != 0 {\n")
				Goproto.WriteString("    x.Int(m.Get" + mvv.Name + "())\n")
			case "long":
				Goproto.WriteString("    if m.Get" + mvv.Name + "() != 0 {\n")
				Goproto.WriteString("    x.Long(m.Get" + mvv.Name + "())\n")
			case "string":
				Goproto.WriteString("    if m.Get" + mvv.Name + "() != \"\" {\n")
				Goproto.WriteString("    x.String(m.Get" + mvv.Name + "())\n")
			case "vector<unsigned byte>":
				Goproto.WriteString("    if m.Get" + mvv.Name + "() != nil {\n")
				Goproto.WriteString("    x.Bytes(m.Get" + mvv.Name + "())\n")
			case "vector<vector<unsigned byte>>":
				Goproto.WriteString("    if m.Get" + mvv.Name + "()!= nil {\n")
				Goproto.WriteString("        x.Int(481674261)\n")
				Goproto.WriteString("        x.Int(int32(len(m.Get" + mvv.Name + "())))\n")
				Goproto.WriteString("        for _, v := range m.Get" + mvv.Name + "() {\n")
				Goproto.WriteString("            x.Buf = append(x.Buf, v...)\n")

				Goproto.WriteString("        }\n")
			default:
				r := mvv.Type[:3]
				if r == "vec" {
					Goproto.WriteString("    if len(m.Get" + mvv.Name + "()) >0  {\n")
					Goproto.WriteString("        x.Int(481674261)\n")
					Goproto.WriteString("        x.Int(int32(len(m.Get" + mvv.Name + "())))\n")
					Goproto.WriteString("        for _, v := range m.Get" + mvv.Name + "() {\n")
					Goproto.WriteString("            x.Buf = append(x.Buf, v.Encode()...)\n")

					Goproto.WriteString("        }\n")
				} else {
					Goproto.WriteString("    if m.Get" + mvv.Name + "().Constructor!=0 {\n")
					Goproto.WriteString("    x.Bytes(m.Get" + mvv.Name + "().Encode())\n")
				}
			}
			Goproto.WriteString("    }\n")
		} else {

			switch mvv.Type {
			case "bool":
				//ss ="Get"+ s + "()!=false"
				Goproto.WriteString("    x.Bytes(m.Get" + mvv.Name + "())\n")
			case "int":
				//ss ="Get"+ s + "()!=false"
				Goproto.WriteString("    x.Int(m.Get" + mvv.Name + "())\n")
			case "long":
				Goproto.WriteString("    x.Long(m.Get" + mvv.Name + "())\n")
			case "string":
				Goproto.WriteString("    x.String(m.Get" + mvv.Name + "())\n")
			case "vector<unsigned byte>":
				Goproto.WriteString("    x.Bytes(m.Get" + mvv.Name + "())\n")
			case "vector<vector<unsigned byte>>":
				Goproto.WriteString("        x.Int(481674261)\n")
				Goproto.WriteString("        x.Int(int32(len(m.Get" + mvv.Name + "())))\n")
				Goproto.WriteString("        for _, v := range m.Get" + mvv.Name + "() {\n")
				Goproto.WriteString("            x.Buf = append(x.Buf, v...)\n")

				Goproto.WriteString("        }\n")
			default:
				r := mvv.Type[:3]
				if r == "vec" {
					Goproto.WriteString("        x.Int(481674261)\n")
					Goproto.WriteString("        x.Int(int32(len(m.Get" + mvv.Name + "())))\n")
					Goproto.WriteString("        for _, v := range m.Get" + mvv.Name + "() {\n")
					Goproto.WriteString("            x.Buf = append(x.Buf, v.Encode()...)\n")

					Goproto.WriteString("        }\n")
				} else {
					Goproto.WriteString("    x.Bytes(m.Get" + mvv.Name + "().Encode())\n")
				}
			}
		}

	}
	Goproto.WriteString("    return x.Buf\n")
	Goproto.WriteString("}\n")
}

//mtproto转go
func TarsTransferType(s string) string {

	var ss string
	switch s {
	case "bool":
		ss = "Bytes"
	case "uint32":
		ss = "UInt"
	case "int":
		ss = "Int"
	case "long":
		ss = "Long"
	case "double":
		ss = "Double"
	case "string":
		ss = "String"
	case "vector<unsigned byte>":
		ss = "Bytes"
		//Vector<int>
	case "vector<int>":
		ss = "VectorInt"
	case "vector<long>":
		ss = "VectorLong"
	case "vector<string>":
		ss = "VectorString"
	default:
		ss = "Bytes"
	}
	return ss
}
func TarsEncodeTypere(t, s string) string {

	var ss string
	switch t {
	case "bool":
		ss = "Get" + s + "()!=false"
	case "int32", "int64", "float32", "float64":
		ss = "Get" + s + "()!=0"
	case "bytes":
		ss = "Get" + s + "() != nil"
	case "string":
		ss = "Get" + s + "()!=\"\""
		//Vector<int>
	case "Vector<int>":
		ss = "Get" + s + "()!=0"
	case "Vector<long>":
		ss = "Get" + s + "()!=0"
	case "Vector<string>":
		ss = "Get" + s + "()!=0"
	case "Vector<bytes>":
		ss = "Get" + s + "()!=0"
	default:
		ss = "Get" + s + "().Constructor!=0"
	}
	return ss
}
func TarsTransferName(t, s string) string {

	var ss string
	switch t {
	case "int", "long", "double", "bool", "string":
		ss = s
	default:
		ss = s + "().Encode"
	}
	return ss
}

func TarsJudgeType(s string) string {
	var ss string

	switch s {
	case "int":
		ss = "int32"
	case "long":
		ss = "int64"
	case "bool":
		ss = "bool"
	case "string":
		ss = "string"
	case "vector<unsigned byte>":
		ss = "[]uint8"
	case "vector<int>":
		ss = "[]int32"
	case "vector<vector<unsigned byte>>":
		ss = "[][]uint8"
	case "double":
		ss = "float64"

	default:
		//[]*MessageEntity
		r := regexp.MustCompile("^vector<(.*)>$")
		if r.MatchString(s) {
			sss := r.FindStringSubmatch(s)
			ss = "[]" + strings.Title(sss[1])
		} else {
			ss = strings.Title(s)
		}

	}
	return ss
}

//去除Vector<,去除.
//Vector<messages.SearchCounter>
func TarsTypeFormat(s string) string {
	//reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
	//s = TarsChangeTypeS(s)
	s = strings.Replace(s, "Vector", "vector", -1)
	s = strings.Replace(s, ".", "_", -1)
	return s

}

//取flagsid
func GetTarsFlagsId(s string) (string, error) {
	flags := strings.Split(s, "?")
	if len(flags) > 1 {
		flagids := strings.Split(flags[0], ".")
		return flagids[1], nil
	} else {
		return "", errors.New("wu")
	}
}

//去掉flags.0?
func TarsReplaceFlags(s string) string {
	flags := strings.Split(s, "?")
	if len(flags) > 1 {
		return flags[1]
	} else {
		return s
	}
}

func TarsVectorToRepeated(s string) string {
	s = TarsChangeTypeS(s)
	s = strings.Replace(s, "Vector", "vector", -1)
	return s
}

func TarsChangeTypeS(s string) string {
	var ss string
	switch s {
	case "true":
		ss = "bool"
	case "!X":
		ss = "vector<unsigned byte>"
	case "bytes":
		ss = "vector<unsigned byte>"
		//Vector<int>
	case "Vector<bytes>":
		ss = "vector<vector<unsigned byte>>"
	default:
		ss = s
	}
	return ss
}

//去重复
func TarsDeduplication(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return false
		}
	}
	return true
}
func TarskeywordSubstitution(s string) string {
	//
	ssarr := []string{"void", "struct", "bool", "byte", "short", "int", "double", "float", "long", "string", "vector", "map", "key", "routekey", "module", "interface", "out", "require", "optional", "false", "true", "enum", "const"}

	ss := s
	for _, v := range ssarr {
		if v == s {
			ss = s + "_key"
		}
	}
	return ss
}
