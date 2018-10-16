package main

import "fmt"

/*
输出:
4
3
2
1
0
9
8
7
6
5
*/

func main() {
	for i := 5; i < 10; i++ {
		defer fmt.Printf("%d\n", i)
	}

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}
