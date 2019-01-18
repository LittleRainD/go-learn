package testType

import "fmt"

func MathSymbol() {
	var a bool = false
	var b bool = true

	if !a && b {
		c := a && b
		fmt.Println("&&", c)
	}

}
