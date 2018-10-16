package main

import "fmt"

//标签名是大小写敏感的。
func myfunc() {
	i := 0
Here:
	fmt.Println(i)
	i++
	goto Here

}

func main() {
	myfunc()
}
