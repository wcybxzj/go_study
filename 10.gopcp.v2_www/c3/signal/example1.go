package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("pid", syscall.Getpid())

	sigRecv1 := make(chan os.Signal, 10)
	sigS1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRecv1, sigS1 ...)

	sigRecv2 := make(chan os.Signal, 10)
	sigS2 := []os.Signal{syscall.SIGQUIT}
	signal.Notify(sigRecv2, sigS2 ...)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for s :=range sigRecv1 {
			fmt.Println("recv sigRecv1 sig is:", s)
		}
		fmt.Println("End sigRecv1")
		wg.Done()
	}()

	go func() {
		for s :=range sigRecv2 {
			fmt.Println("recv sigRecv2 sig is:", s)
		}
		fmt.Println("End sigRecv2")
		wg.Done()
	}()

	fmt.Println("sleep 60 seconds")
	time.Sleep(time.Second*60)
	fmt.Println("before signal.Stop sigRecv1")
	signal.Stop(sigRecv1)
	close(sigRecv1)
	fmt.Println("finish signal.Stop sigRecv1")
}
