package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
	"zuji/common/debug"
	"zuji/common/exec"
)

func createWorker() {
	go func() {
		for {
			inCmd := "sleep " + strconv.Itoa(rand.Intn(300))
			fmt.Println(inCmd)
			err := exec.RunCmdAsyncWithTimeout1(inCmd, 5)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("success:" + inCmd)
			}
		}
	}()
}

func PrintNumGoroutine() {
	for {
		fmt.Println("NumGoroutine:" + strconv.Itoa(runtime.NumGoroutine()))
		time.Sleep(time.Second)
	}
}

func main() {
	debug.IsDebug = true
	rand.Seed(time.Now().UnixNano())

	go PrintNumGoroutine()

	for i := 0; i < 10; i++ {
		createWorker()
	}

	time.Sleep(time.Second * 300)
}
