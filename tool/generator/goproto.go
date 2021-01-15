package main

import (
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func ToGoProto() {

	Pbproto.WriteString("package mtproto\n")

	//整理参数
	FieldConversion()

	//结构
	GoConstructorsToProto()
	//接口
	GoMethodsToProto()
	//参数与小类型解码
	GoCodec()

}

type FieldParamsType struct {
	Name   string
	Type   string
	FlagId string
}
type FieldConParams struct {
	Name   string
	Id     string
	Flag   bool
	Params []FieldParamsType
}
type FieldCon struct {
	Type               map[string][]FieldParamsType //大类型下面的字段
	MethodParams       map[string][]FieldConParams  //接口参数下面的字段
	ConstructorsParams map[string][]FieldConParams  //小类型下面的字段
	TypeConstructors   map[string][]string          //大类型下面的小类型
	TypeMethod         []string                     //接口返回类型下面的小类型【只有Vector这种，其他的包含在TypeConstructors里】

}

var FC = FieldCon{
	Type:               make(map[string][]FieldParamsType),
	MethodParams:       make(map[string][]FieldConParams),
	ConstructorsParams: make(map[string][]FieldConParams),
	TypeConstructors:   make(map[string][]string),
	TypeMethod:         nil,
}

func CodecEncode(fcp FieldConParams) {
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
	Pbproto.WriteString("// CodecEncode:" + fcp.Name + " \n")
	Pbproto.WriteString("func (m *TL" + fcp.Name + ") Encode() []byte {\n")
	Pbproto.WriteString("    x := NewEncodeBuf(512)\n")
	Pbproto.WriteString("    x.Int(" + fcp.Id + ")\n")
	//flags
	if fcp.Flag {
		//var flags uint32 = 0
		Pbproto.WriteString("    var flags uint32 = 0\n")
	}
	for _, mvv := range fcp.Params {
		if mvv.FlagId != "" {
			Pbproto.WriteString("    if m.Get" + mvv.Name + "() != " + TransferTypere(mvv.Type) + " {\n")
			Pbproto.WriteString("        flags |= " + mvv.FlagId + " << 0\n")
			Pbproto.WriteString("    }\n")

		}
	}

	if fcp.Flag {
		Pbproto.WriteString("    x.UInt(flags)\n")
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

			r := mvv.Type[:3]

			if r == "[]*" || r == "[][" {
				Pbproto.WriteString("    if m.Get" + mvv.Name + "() != " + TransferTypere(mvv.Type) + " {\n")
				Pbproto.WriteString("        x.Int(481674261)\n")
				Pbproto.WriteString("        x.Int(int32(len(m.Get" + mvv.Name + "())))\n")
				Pbproto.WriteString("        for _, v := range m.Get" + mvv.Name + "() {\n")
				if r == "[][" {
					Pbproto.WriteString("            x.Buf = append(x.Buf, v...)\n")
				} else {

					Pbproto.WriteString("            x.Buf = append(x.Buf, (*v).Encode()...)\n")
				}

				Pbproto.WriteString("        }\n")
			} else {
				Pbproto.WriteString("    if m.Get" + mvv.Name + "() != " + TransferTypere(mvv.Type) + " {\n")
				Pbproto.WriteString("    x." + GoTransferType(mvv.Type) + "(m.Get" + TransferName(mvv.Type, mvv.Name) + "())\n")
			}

			Pbproto.WriteString("    }\n")
		} else {
			r := mvv.Type[:3]
			if r == "[]*" || r == "[][" {
				Pbproto.WriteString("    x.Int(481674261)\n")
				Pbproto.WriteString("    x.Int(int32(len(m.Get" + mvv.Name + "())))\n")
				Pbproto.WriteString("    for _, v := range m.Get" + mvv.Name + "() {\n")
				if r == "[][" {
					Pbproto.WriteString("            x.Buf = append(x.Buf, v...)\n")
				} else {

					Pbproto.WriteString("            x.Buf = append(x.Buf, (*v).Encode()...)\n")
				}
				Pbproto.WriteString("       }\n")
			} else {

				Pbproto.WriteString("    x." + GoTransferType(mvv.Type) + "(m.Get" + TransferName(mvv.Type, mvv.Name) + "())\n")
			}
		}

	}
	Pbproto.WriteString("    return x.Buf\n")
	Pbproto.WriteString("}\n")
}
func FieldConversion() {

	//大类型下的字段
	for k, v := range PbConstructors.Type {
		//字段对应
		FieldMap := make(map[string]string)
		//去掉 Vector
		if k == "Vector t" {
			continue
		}
		pdid := 1
		typename := GoTypeFormat(k)
		var DeduplicationString = []string{}
		var RepeatedFields = []string{}
		for _, vvv := range PbConstructors.Params[k] {
			//去掉 # flag
			if vvv.Type == "#" {
				continue
			}
			//去掉flags.0?
			vvvType := GoReplaceFlags(vvv.Type)
			//Vector<Message>换repeated Message
			vvvType = GoVectorToRepeated(vvvType)

			//去重复
			dc := GoDeduplication(DeduplicationString, vvvType+vvv.Name)
			if !dc {
				continue
			}

			DeduplicationString = append(DeduplicationString, vvvType+vvv.Name)

			//重复字段添加id
			vvvName := vvv.Name
			rf := GoDeduplication(RepeatedFields, vvv.Name)

			if !rf {
				vvvName = vvv.Name + "_" + strconv.Itoa(pdid)
			}
			RepeatedFields = append(RepeatedFields, vvv.Name)

			vvvName = keywordSubstitution(vvvName)
			FieldMap[vvvType+vvv.Name] = vvvName
			FieldParamsType := FieldParamsType{
				Name: vvvName,
				Type: GoTypeFormat(vvvType),
			}
			FC.Type[typename] = append(FC.Type[typename], FieldParamsType)
			pdid++
		}
		//大类型下面小类型
		for _, mv := range v {
			mvPredicate := GoTypeFormat(mv.Predicate)
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
				vvvType := GoReplaceFlags(xv.Type)
				//Vector<Message>换repeated Message
				vvvType = GoVectorToRepeated(vvvType)

				//flagid
				flagid, _ := GetFlagsId(xv.Type)
				ParamsType := FieldParamsType{
					Name:   FieldMap[vvvType+xv.Name],
					Type:   GoTypeFormat(vvvType),
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

		//接口返回的数组
		r := regexp.MustCompile("^Vector<(.*)>$")
		if r.MatchString(PbMk) {
			ss := r.FindStringSubmatch(PbMk)
			FC.TypeMethod = append(FC.TypeMethod, GoChangeType(GoTypeFormat(ss[1])))
		}
		typename := GoTypeFormat(PbMk)
		//接口参数
		for _, mv := range PbMv {
			mvMethod := GoTypeFormat(mv.Method)
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
				vvvType := GoReplaceFlags(xv.Type)
				//Vector<Message>换repeated Message
				vvvType = GoVectorToRepeated(vvvType)
				//flagid
				flagid, _ := GetFlagsId(xv.Type)
				//替换
				ParamsType := FieldParamsType{
					Name:   keywordSubstitution(GoTypeFormat(xv.Name)),
					Type:   GoTypeFormat(vvvType),
					FlagId: flagid,
				}
				FieldConParam.Params = append(FieldConParam.Params, ParamsType)
			}
			FieldConParam.Flag = isflag
			FieldConParam.Id = mv.ID
			FC.MethodParams[typename] = append(FC.MethodParams[typename], FieldConParam)
		}
	}
}
func GoCodec() {
	for _, mv := range FC.MethodParams {

		//func NewTLMessagesGetMessages() *TLMessagesGetMessages {
		//	return &TLMessagesGetMessages{}
		//}
		for _, mmv := range mv {
			Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
			Pbproto.WriteString("// MethodCodec:New:" + mmv.Name + " \n")
			Pbproto.WriteString("//\n")

			CodecEncode(mmv)
		}
	}
	for _, mv := range FC.ConstructorsParams {

		//func NewTLMessagesGetMessages() *TLMessagesGetMessages {
		//	return &TLMessagesGetMessages{}
		//}
		for _, mmv := range mv {
			Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
			Pbproto.WriteString("// ConstructorsCodec:New:" + mmv.Name + " \n")
			Pbproto.WriteString("//\n")

			CodecEncode(mmv)
		}
	}

	//大类型
	for dk, dv := range FC.ConstructorsParams {
		Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
		Pbproto.WriteString("// ConstructorsCodec:Type:" + dk + " \n")
		Pbproto.WriteString("//\n")
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
		Pbproto.WriteString("func (m *" + dk + ") Encode() []byte {\n")
		Pbproto.WriteString("    switch m.constructor {\n")
		for _, dvv := range dv {
			Pbproto.WriteString("    case " + dvv.Id + ":\n")
			Pbproto.WriteString("        t := m.To_" + dvv.Name + "()\n")
			Pbproto.WriteString("        return t.Encode()\n")
		}

		Pbproto.WriteString("	default:\n")
		Pbproto.WriteString("		return nil\n")
		Pbproto.WriteString("	}\n")
		Pbproto.WriteString("}\n")
		//func (m *Bool) To_BoolFalse() *TLBoolFalse {
		//	return &TLBoolFalse{
		//		Data2: m.Data2,
		//	}
		//}
		for _, dvv := range dv {
			Pbproto.WriteString("func (m *" + dk + ") To_" + dvv.Name + "() *TL" + dvv.Name + " {\n")
			Pbproto.WriteString("    return &TL" + dvv.Name + "{\n")
			Pbproto.WriteString("        data: m.data,\n")
			Pbproto.WriteString("    }\n")
			Pbproto.WriteString("}\n")
		}

	}

}
func TransferTypere(s string) string {

	var ss string
	switch s {
	case "bool":
		ss = "false"
	case "int32", "int64", "float32", "float64":
		ss = "0"
	case "bytes":
		ss = "nil"
	case "string":
		ss = "\"\""
		//Vector<int>
	case "Vector<int>":
		ss = "VectorInt"
	case "Vector<long>":
		ss = "0"
	case "Vector<string>":
		ss = "0"
	case "Vector<bytes>":
		ss = "0"
	default:
		ss = "nil"
	}
	return ss
}
func TransferName(t, s string) string {
	var ss string
	switch t {
	case "uint32", "int32", "int64", "float64", "string", "[]byte", "[]int32", "[]int64", "[]string":
		ss = s
	default:
		ss = s + "().Encode"
	}
	return ss
}

//mtproto转go
func GoTransferType(s string) string {

	var ss string
	switch s {
	case "bool":
		ss = "Bytes"
	case "uint32":
		ss = "UInt"
	case "int32":
		ss = "Int"
	case "int64":
		ss = "Long"
	case "float64":
		ss = "Double"
	case "string":
		ss = "String"
	case "[]byte":
		ss = "Bytes"
		//Vector<int>
	case "[]int32":
		ss = "VectorInt"
	case "[]int64":
		ss = "VectorLong"
	case "[]string":
		ss = "VectorString"
	default:
		ss = "Bytes"
	}
	return ss
}

//类型转对应空状态
func TransferType(s string) string {
	var ss string
	switch s {
	case "true":
		ss = "bool"
	case "int":
		ss = "0"
	case "long":
		ss = "0"
	case "!X":
		ss = "nil"
	case "double":
		ss = "0"
	case "bytes":
		ss = "nil"
		//Vector<int>
	case "Vector<int>":
		ss = "0"
	case "Vector<long>":
		ss = "0"
	case "Vector<string>":
		ss = "0"
	case "Vector<bytes>":
		ss = "0"
	default:
		ss = "nil"
	}
	return ss
}

//接口
func GoMethodsToProto() {

	//接口参数
	for _, mv := range FC.MethodParams {
		for _, mmv := range mv {
			Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
			Pbproto.WriteString("// Methods:Params:" + mmv.Name + " \n")
			Pbproto.WriteString("//\n")
			Pbproto.WriteString("type TL" + mmv.Name + " struct {\n")
			for _, mvv := range mmv.Params {
				Pbproto.WriteString("	" + mvv.Name + " " + GoJudgeType(mvv.Type) + "\n")
			}
			Pbproto.WriteString("}\n")
			for _, mvv := range mmv.Params {
				Pbproto.WriteString("func (m *TL" + mmv.Name + ")  Set" + mvv.Name + "(v " + GoJudgeType(mvv.Type) + "){m." + mvv.Name + "=v}\n")
				Pbproto.WriteString("func (m *TL" + mmv.Name + ")  Get" + mvv.Name + "() " + GoJudgeType(mvv.Type) + "{return m." + mvv.Name + "}\n")

			}

		}
	}
	//接口返回值
	//Vector<
	for _, vv := range FC.TypeMethod {

		//找Vector<的
		//message Vector_DialogFilterSuggested {
		//    repeated DialogFilterSuggested datas = 1;
		//}
		Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
		Pbproto.WriteString("// Methods:VectorType:" + vv + " \n")
		Pbproto.WriteString("//\n")
		Pbproto.WriteString("type V" + GoChangeType(GoTypeFormat(vv)) + " struct {\n")
		Pbproto.WriteString("  data []*" + GoChangeType(GoTypeFormat(vv)) + "\n")
		Pbproto.WriteString("}\n")
	}

	//接口rpc
	//service RPCLangpack {
	//// langpack.getStrings#efea3803 lang_pack:string lang_code:string keys:Vector<string> = Vector<LangPackString>;
	//    rpc langpack_getStrings(TL_langpack_getStrings) returns (Vector_LangPackString) {}
	// }

}

//类型
func GoConstructorsToProto() {
	//结构
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
		Pbproto.WriteString("type " + k + " struct {\n")
		Pbproto.WriteString("	constructor int32\n")
		Pbproto.WriteString("	data *" + k + "_Data\n")
		Pbproto.WriteString("}\n")

		//类型数据
		/*	type ChatBannedRights_Data struct {
			ViewMessages         bool */
		Pbproto.WriteString("// Constructors:Type:Data:" + k + " \n")
		Pbproto.WriteString("type " + k + "_Data struct {\n")
		for _, vvv := range FC.Type[k] {
			Pbproto.WriteString("	" + vvv.Name + " " + GoJudgeType(vvv.Type) + " \n")
		}
		Pbproto.WriteString("}\n")
		for _, vv := range v {
			Pbproto.WriteString("// Constructors:Predicate:" + vv + " \n")
			Pbproto.WriteString("type TL" + vv + " struct {\n")
			Pbproto.WriteString("	data *" + k + "_Data\n")
			Pbproto.WriteString("}\n")

			//func NewTLInputEncryptedChat() *TLInputEncryptedChat {
			//	return &TLInputEncryptedChat{Data2: &InputEncryptedChat_Data{}}
			//}

			Pbproto.WriteString("// Constructors:New:" + vv + " \n")
			Pbproto.WriteString("func NewTL" + vv + "() *TL" + vv + " {\n")
			Pbproto.WriteString("	return &TL" + vv + "{data: &" + k + "_Data{}}\n")
			Pbproto.WriteString("}\n")
		}
		for _, cvv := range FC.ConstructorsParams[k] {

			cname := cvv.Name
			for _, ccvv := range cvv.Params {
				Pbproto.WriteString("func (m *TL" + cname + ")  Set" + ccvv.Name + "(v " + GoJudgeType(ccvv.Type) + "){m.data." + ccvv.Name + "=v}\n")
				Pbproto.WriteString("func (m *TL" + cname + ")  Get" + ccvv.Name + "() " + GoJudgeType(ccvv.Type) + "{return m.data." + ccvv.Name + "}\n")

			}
		}

	}
}

//换类型

func GoChangeType(s string) string {
	s = GoChangeTypeS(s)
	//repeated MessageEntityentities
	xType := strings.Split(s, "]")
	if len(xType) > 1 {
		return xType[0] + "]" + GoChangeTypeS(xType[1])
	} else {
		return GoChangeTypeS(s)
	}
}
func keywordSubstitution(s string) string {
	var ss string
	switch s {
	case "default":
		ss = "default_key"
	case "type":
		ss = "type_key"
	case "range":
		ss = "range_key"
	default:
		ss = s
	}
	return ss
}
func GoChangeTypeS(s string) string {
	var ss string
	switch s {
	case "true":
		ss = "bool"
	case "int":
		ss = "int32"
	case "long":
		ss = "int64"
	case "!X":
		ss = "[]byte"
	case "double":
		ss = "float64"
	case "bytes":
		ss = "[]byte"
	case "string":
		ss = "string"
		//Vector<int>
	case "Vector<int>":
		ss = "[]int32"
	case "Vector<long>":
		ss = "[]int64"
	case "Vector<string>":
		ss = "[]string"
	case "Vector<bytes>":
		ss = "[][]byte"
	default:
		ss = s
	}
	return ss
}
func GoJudgeType(s string) string {
	var ss string

	switch s {
	case "bool", "[][]byte", "string", "int32", "int64", "[]byte", "float64", "[]int32", "[]int64", "[]string":
		ss = s
	default:
		//[]*MessageEntity

		r := s[:3]
		if r == "[]*" {
			ss = s
		} else {
			ss = "*" + s
		}

	}
	return ss
}

//去除Vector<,去除.
//Vector<messages.SearchCounter>
func GoTypeFormat(s string) string {
	//reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
	r := regexp.MustCompile("^Vector<(.*)>$")
	if r.MatchString(s) {
		ss := r.FindStringSubmatch(s)
		s = ss[1]
	}
	s = strings.Replace(s, ".", "_", -1)
	return s

}

//去掉flags.0?
func GoReplaceFlags(s string) string {
	flags := strings.Split(s, "?")
	if len(flags) > 1 {
		return flags[1]
	} else {
		return s
	}
}

//取flagsid
func GetFlagsId(s string) (string, error) {
	flags := strings.Split(s, "?")
	if len(flags) > 1 {
		flagids := strings.Split(flags[0], ".")
		return flagids[1], nil
	} else {
		return "", errors.New("wu")
	}
}

//Vector转换repeated
func GoVectorToRepeated(s string) string {
	s = GoChangeTypeS(s)
	r := regexp.MustCompile("^Vector<(.*)>$")
	if r.MatchString(s) {
		ss := r.FindStringSubmatch(s)
		s = "[]*" + ss[1]
	}

	return s
}

//去重复
func GoDeduplication(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return false
		}
	}
	return true
}
func GoSaveFile(filename string) {
	beauty := Pbproto.Bytes()
	err := ioutil.WriteFile(filename, beauty, 0666)
	if err != nil {
		fmt.Println(err)
	}
	beauty, err = format.Source(Goproto.Bytes())
	if err != nil {
		fmt.Println("go fmt fail. " + filename + " " + err.Error())
	}
	err = ioutil.WriteFile("tarstl/api/codec.go", beauty, 0666)
	if err != nil {
		fmt.Println(err)
	}
	//Baseproto
	Basebeauty, err := format.Source(Baseproto.Bytes())
	if err != nil {
		fmt.Println("go fmt fail. " + filename + " " + err.Error())
	}
	err = ioutil.WriteFile("tarstl/api/registers.go", Basebeauty, 0666)
}
