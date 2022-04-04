package main

import (
	"fmt"
	"github.com/bitfield/script"
)

func main() {
	fmt.Println("run ping")
	p := script.Exec("ping -c 3 127.0.0.1")
	p.Stdout()
	fmt.Printf("exit status: %v\n", p.ExitStatus())
	fmt.Println("exit..")
}
