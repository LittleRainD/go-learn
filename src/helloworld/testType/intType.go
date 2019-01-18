package testType

import (
	"fmt"
)

func boolType() {

}

func Numer() {
	println(1 + 1)

	//变量的声明
	var aa int
	aa = 1
	//or
	var bb int = 2
	//or
	cc := 3
	fmt.Println(aa, bb, cc)

	// 如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误，
	// 例如下面这个例子当中的变量 dd声明成功但是不用编译器会报错【src/helloworld/testType/intType.go:21:2: dd declared and not used】

	//dd := 333

	// 多变量同时在一行声明
	var a, b, c int
	a, b, c = 1, 2, 3
	println("多变量同时在同一行声明：a,b,c = ", a, b, c)

	aStr, bInt, cBool := "字符串", 2, true

	println("不同变量在同一行声明：", aStr, bInt, cBool)

	//***********************简单的交换俩个同类型变量的值*****************
	a, b = b, a
	println("交换后的变量值,a,b,c = ", a, b, c)

	// 空白标识符也没用作抛弃值
	// _ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
	//并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。
	_, b = 1, 2

	//************************************常量类型*********************************************

	//demo
	Circle(1.2)

	//可以当成常量池来看,常量是可以声明但是不用的
	const (
		unknow = 0
		man    = 1
		woman  = 2
	)

	fmt.Println("未知性别", unknow)

	//默认是A值
	const (
		A = 1
		B
		C
		D = iota
	)
	fmt.Println("A,B,C,D:", A, B, C, D)
	//iota 类似常量的计数器
	const (
		MONDAY = iota
		Thursday
		hhh  = 'a'
		WUYI = "WUYI"
		TIME
		W = iota
		Z
	)

	fmt.Println("look ", MONDAY, Thursday, hhh, WUYI, TIME, W, Z)
	//iota特殊的常量配参数
}
