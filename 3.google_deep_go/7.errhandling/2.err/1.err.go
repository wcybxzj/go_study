package main

import (
	"bufio"
	"fmt"
	"os"
	"go_study/3.google_deep_go/7.errhandling/fib"
)

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	//自定义error
	//err = errors.New("this is a custom error")
	if err != nil {
		//fmt.Println("file exists")
		//fmt.Println("Error:", err)
		//fmt.Println("Error:", err.Error())
		//fmt.Println("======================")
		if pathError, ok := err.(*os.PathError); !ok {
			fmt.Println("说明err是未知类型")
			panic(err)
		} else {
			fmt.Println("说明error是已知道类型:os.PathError")
			fmt.Printf("打印真实的错误情况:%s,%s,%s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonnaci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}
