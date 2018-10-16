package main

import (
	"fmt"
	"unsafe"
)

//(u)int:
//和操作系统走,64系统就是64位

//uintptr:
//指针类型, 和操作系统走,64系统就是64位

//byte:1字节/8位

//rune:4字节/32位
//因为union code是2字节,utf8是3字节,所以rune用了4字节
//gloang中的char, 不用char是为了避免躲过语言,char是1个字节,在多国语言存在很大的问题

//复数:
//complex64=实部float32 虚部float32
//complex128=实部float64 虚部float64
func main() {
	var a int = 123
	fmt.Println(unsafe.Sizeof(a)) //8字节
}
