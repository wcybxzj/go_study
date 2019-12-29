package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("pid", syscall.Getpid())
	sigRecv := make(chan os.Signal, 10)
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRecv, sigs ...)
	for s :=range sigRecv {
		fmt.Println("recv signal is:", s)
		time.Sleep(time.Second*3)
	}
}
