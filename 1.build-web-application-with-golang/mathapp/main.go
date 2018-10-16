package main

import (
	"fmt"
	"mymath"
)

//这个的package是main，import里面调用的包是mymath,这个就是相对于$GOPATH/src的路径.
//如果是多级目录，就在import里面引入多级目录，如果你有多个GOPATH，也是一样，Go会自动在多个$GOPATH/src中寻找。





func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
}
