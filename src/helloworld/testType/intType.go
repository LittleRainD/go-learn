package testType

import "fmt"

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
}
