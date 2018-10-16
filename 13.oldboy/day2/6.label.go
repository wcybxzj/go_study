package main

import (
	"fmt"
	"time"
)

//break+label
//continue+label
//goto+label

/*
i is 0, and j is:0
i is 1, and j is:0
i is 2, and j is:0
ok!
*/
func continueTest() {
ABC:
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if j == 1 {
				continue ABC
			}

			fmt.Printf("i is %d, and j is:%d\n", i, j)
		}
	}
	fmt.Println("ok!") //可以执行
}

/*
i is 0, and j is:0
ok!
*/
func breakTest() {
ABC:
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if j == 1 {
				break ABC
			}

			fmt.Printf("i is %d, and j is:%d\n", i, j)
		}
	}
	fmt.Println("ok!") //可以执行
}

/*
i is 0, and j is:0
i is 0, and j is:0
i is 0, and j is:0
i is 0, and j is:0
i is 0, and j is:0
.......永远重复.....
*/
func gotoTest() {
ABC:
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if j == 1 {
				goto ABC
			}
			fmt.Printf("i is %d, and j is:%d\n", i, j)
			time.Sleep(time.Second)
		}
	}
	fmt.Println("ok!") //执行不到
}

func main() {
	//continueTest()
	//breakTest()
	gotoTest()
}
