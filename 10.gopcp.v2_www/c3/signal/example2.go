package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("example2 pid", syscall.Getpid())
	sigRecv := make(chan os.Signal, 10)
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT, syscall.SIGUSR1}
	signal.Notify(sigRecv, sigs ...)
	for s :=range sigRecv {
		fmt.Println("recv signal is:", s)
		time.Sleep(time.Second*3)
	}
}
