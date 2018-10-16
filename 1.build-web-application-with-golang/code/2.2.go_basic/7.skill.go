package main

import (
	"fmt"
	"os"
)

const (
	i      = 100
	pi     = 3.14
	prefix = "go_"
)

func main() {
	pid := os.Getpid()
	fmt.Printf("%d\n", pid)
	fmt.Printf("%d %f %s\n", pid, pi, prefix)

}
