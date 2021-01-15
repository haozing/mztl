package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func ToProto() {
	//syntax = "proto3";
	//
	//package mtproto;
	//
	//option go_package = ".;mtproto";
	//Pbproto.WriteString(`syntax = "proto3";\n`)
	Pbproto.WriteString("syntax = \"proto3\";\n")
	Pbproto.WriteString("package main;\n")
	Pbproto.WriteString("option go_package = \".;main\";\n")
	///////////////////////////////////////////////////////////////////////////////
	// ChatBannedRights <--
	//  + TL_chatBannedRights
	//
	ConstructorsToProto()
	MethodsToProto()

}

//接口
func MethodsToProto() {
	servicerpc := make(map[string][]Methods)
	//接口参数
	for _, mv := range TLjson.Methods {
		//制作rpc归类
		servicerpcgl := strings.Split(mv.Method, ".")
		if len(servicerpcgl) > 1 {
			servicerpc[servicerpcgl[0]] = append(servicerpc[servicerpcgl[0]], mv)
		} else {
			servicerpc["other"] = append(servicerpc["other"], mv)
		}

		mpdid := 1
		mvMethod := TypeFormat(mv.Method)
		Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
		Pbproto.WriteString("// Methods:Params:" + mv.Method + " \n")
		Pbproto.WriteString("//\n")
		Pbproto.WriteString("message TL_" + mvMethod + " {\n")
		for _, mvv := range mv.Params {
			//去掉 # flag
			if mvv.Type == "#" {
				continue
			}
			//去掉flags.0?
			vvvType := ReplaceFlags(mvv.Type)
			//Vector<Message>换repeated Message
			vvvType = VectorToRepeated(vvvType)
			//换类型
			//message TL_invokeWithTakeout {
			//    int64 takeout_id = 1;
			//    bytes query = 2;
			//}
			vvvType = ChangeType(vvvType)

			Pbproto.WriteString("	" + TypeFormat(vvvType) + " " + TypeFormat(mvv.Name) + " = " + strconv.Itoa(mpdid) + ";\n")
			mpdid++
		}
		Pbproto.WriteString("}\n")
	}
	//接口返回值
	//Vector<
	for vk, _ := range PbMethods.Type {

		//找Vector<的
		//message Vector_DialogFilterSuggested {
		//    repeated DialogFilterSuggested datas = 1;
		//}

		r := regexp.MustCompile("^Vector<(.*)>$")
		if r.MatchString(vk) {
			ss := r.FindStringSubmatch(vk)
			Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
			Pbproto.WriteString("// Methods:VectorType:" + vk + " \n")
			Pbproto.WriteString("//\n")
			Pbproto.WriteString("message V_" + ChangeType(TypeFormat(ss[1])) + " {\n")
			Pbproto.WriteString("  repeated " + ChangeType(TypeFormat(ss[1])) + " data = 1;\n")
			Pbproto.WriteString("}\n")
		}
	}

	//接口rpc
	//service RPCLangpack {
	//// langpack.getStrings#efea3803 lang_pack:string lang_code:string keys:Vector<string> = Vector<LangPackString>;
	//    rpc langpack_getStrings(TL_langpack_getStrings) returns (Vector_LangPackString) {}
	// }

	for rpck, rpcv := range servicerpc {

		//过滤other
		if rpck == "other" {
			continue
		}
		Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
		Pbproto.WriteString("// RPC:" + rpck + " \n")
		Pbproto.WriteString("//\n")
		Pbproto.WriteString("service R_" + rpck + " {\n")

		for _, rpcvv := range rpcv {
			r := regexp.MustCompile("^Vector<(.*)>$")
			rpcvvType := ChangeType(rpcvv.Type)
			if r.MatchString(rpcvv.Type) {
				ss := r.FindStringSubmatch(rpcvv.Type)
				rpcvvType = "V_" + ChangeType(ss[1])

			}

			Pbproto.WriteString("  rpc " + TypeFormat(rpcvv.Method) + "(TL_" + TypeFormat(rpcvv.Method) + ") returns (" + TypeFormat(rpcvvType) + "){}\n")

		}
		Pbproto.WriteString("}\n")
	}
}

//类型
func ConstructorsToProto() {
	//结构
	for k, v := range PbConstructors.Type {

		//去掉 Vector
		if k == "Vector t" {
			continue
		}
		pdid := 1

		typename := TypeFormat(k)
		//注释
		Pbproto.WriteString("///////////////////////////////////////////////////////////////////////////////\n")
		Pbproto.WriteString("// Constructors:" + k + " \n")
		Pbproto.WriteString("// " + typename + " <--\n")
		for _, vv := range v {
			Pbproto.WriteString("// + TL_" + TypeFormat(vv.Predicate) + "\n")
		}
		Pbproto.WriteString("//\n")
		//接口返回值类型
		/*	message ChatBannedRights {
			TLConstructor constructor = 1;
			ChatBannedRights_Data data2 = 2;
		}*/
		Pbproto.WriteString("// Constructors:Type:" + k + " \n")
		Pbproto.WriteString("message " + typename + " {\n")
		Pbproto.WriteString("	int32 constructor = 1;\n")
		Pbproto.WriteString("	" + typename + "_Data data = 2;\n")
		Pbproto.WriteString("}\n")

		//接口返回值类型数据
		/*		message ChatBannedRights_Data {
				bool view_messagesbool = 1;
				bool send_messagesbool = 2;
			}*/
		Pbproto.WriteString("// Constructors:Type:Data:" + k + " \n")
		Pbproto.WriteString("message " + typename + "_Data {\n")
		var DeduplicationString = []string{}
		var RepeatedFields = []string{}
		for _, vvv := range PbConstructors.Params[k] {

			//去掉 # flag
			if vvv.Type == "#" {
				continue
			}
			//去掉flags.0?
			vvvType := ReplaceFlags(vvv.Type)
			//Vector<Message>换repeated Message
			vvvType = VectorToRepeated(vvvType)

			//去重复
			dc := Deduplication(DeduplicationString, vvvType+vvv.Name)
			if !dc {
				continue
			}

			DeduplicationString = append(DeduplicationString, vvvType+vvv.Name)

			//重复字段添加id
			vvvName := vvv.Name
			rf := Deduplication(RepeatedFields, vvv.Name)

			if !rf {
				vvvName = vvv.Name + "_" + strconv.Itoa(pdid)
			}
			RepeatedFields = append(RepeatedFields, vvv.Name)
			//换类型
			vvvType = ChangeType(vvvType)
			Pbproto.WriteString("	" + TypeFormat(vvvType) + " " + TypeFormat(vvvName) + " = " + strconv.Itoa(pdid) + ";\n")
			pdid++
		}
		fmt.Println(DeduplicationString)
		Pbproto.WriteString("}\n")

		//接口
		/*		message TL_chatBannedRights {
				ChatBannedRights_Data data2 = 2;
			}*/

		for _, vv := range v {
			Pbproto.WriteString("// Constructors:Predicate:" + vv.Predicate + " \n")
			Pbproto.WriteString("message TL_" + TypeFormat(vv.Predicate) + " {\n")
			Pbproto.WriteString("	" + typename + "_Data data = 1;\n")
			Pbproto.WriteString("}\n")
		}
	}
}

//换类型

func ChangeType(s string) string {

	//repeated MessageEntityentities
	xType := strings.Split(s, " ")
	if len(xType) > 1 {
		return xType[0] + " " + ChangeTypeS(xType[1])
	} else {
		return ChangeTypeS(s)
	}
}

func ChangeTypeS(s string) string {
	var ss string
	switch s {
	case "true":
		ss = "bool"
	case "int":
		ss = "int32"
	case "long":
		ss = "int64"
	case "!X":
		ss = "bytes"
	default:
		ss = s
	}
	return ss
}

//去除Vector<,去除.
//Vector<messages.SearchCounter>
func TypeFormat(s string) string {
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
func ReplaceFlags(s string) string {
	flags := strings.Split(s, "?")
	if len(flags) > 1 {
		return flags[1]
	} else {
		return s
	}
}

//Vector转换repeated
func VectorToRepeated(s string) string {
	r := regexp.MustCompile("^Vector<(.*)>$")
	if r.MatchString(s) {
		ss := r.FindStringSubmatch(s)
		s = "repeated " + ss[1]
	}

	return s
}

//去重复
func Deduplication(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return false
		}
	}
	return true
}
func SaveFile(filename string) {
	beauty := Pbproto.Bytes()
	err := ioutil.WriteFile(filename, beauty, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
