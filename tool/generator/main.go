package main

import "bytes"

var Pbproto bytes.Buffer
var Goproto bytes.Buffer
var Baseproto bytes.Buffer

func main() {
	//获取json
	err := JsonToStruct("json.json")

	if err != nil {
		panic("JsonToStruct err")
	}
	//返回类型结构

	MethodsExtractionType()
	ConstructorsExtractionType()

	//接口结构
	ToGoProto()
	GoSaveFile("go/api.go")
	//类型结构

}
