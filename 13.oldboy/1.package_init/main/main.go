package main

import (
	"fmt"
	_ "go_study/13.oldboy/1.package_init/mylib"
	_ "go_study/13.oldboy/1.package_init/mylib2"
)

func init() {
	fmt.Println("init main")
}

/*
输出:
init a
init b
init c
init main

可以看到main/c.go的init没有被调用
*/
func main() {

}
