package main

import "fmt"

type TZ int

func (tz *TZ) Increase(num int) {
	//*tz=*tz+num //报错(mismatched types TZ and int)
	*tz = *tz + TZ(num) //必须把num强转成TZ
}

func main() {
	var a TZ
	a = 1
	a.Increase(100)
	a.Increase(100)
	a.Increase(100)
	fmt.Println(a)
}
