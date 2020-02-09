//提供的常用库，有一些常用的方法，方便使用

package mylib

import (
	"fmt"
)

// MyPrint function
// def pro print
func MyPrint(a ...interface{}) {
	fmt.Print(a...)
	// fmt.Print(json.Decoder())
	// fmt.Print(strings.Split("s string", "sep string"))
}
