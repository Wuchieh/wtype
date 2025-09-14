package wtype_test

import (
	"fmt"

	"github.com/wuchieh/wtype"
)

func ExampleFileSize() {
	fmt.Println(wtype.B)
	fmt.Println(wtype.KB)
	fmt.Println(wtype.MB)
	fmt.Println(wtype.GB)
	fmt.Println(wtype.TB)
	fmt.Println(wtype.TB + wtype.GB*100)

	//output:
	//1B
	//1.0KB
	//1.0MB
	//1.0GB
	//1.0TB
	//1.1TB
}
