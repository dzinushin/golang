package main

import (
	"fmt"
	"github.com/bitfield/script"
)

func main() {
	fmt.Println("run ping")
	script.Exec("ping -c 3 127.0.0.1").Stdout()

	fmt.Println("exit..")
}
