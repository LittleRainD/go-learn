package testType

import (
	"fmt"
	"strings"
)
//在同一个文件package下是不允许有重复的fun名称的,区分大小写，如果是小写可以当做不重复

var a = "str"
func number() {
	println(1 + 1)
	var aa int
	aa = 1
	//or
	var bb int = 2
	//or
	cc :=3
	fmt.Println(aa,bb,cc)
	var b =strings.Replace("keith","keith","keithorbit",-1);
	println(b)
}
func NamePutOutStringHandler()  {
	var str string = "I'm keith"
	var nStr string  = "not keith"
	var b bool = 0==0

	//这种不带声明格式的只能在函数体中出现
	aString,bTypeInt,cBool := "hi",2,true
	if b {
		fmt.Println(str)
	}else {
		fmt.Println(nStr)
	};

	fmt.Println(aString,bTypeInt,cBool)

	fmt.Println(a)
}
