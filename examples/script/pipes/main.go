package main

import (
	"fmt"
	"github.com/bitfield/script"
)

func main() {
	echoPipe()
	yqPipe()
}

func echoPipe() {
	p := script.Echo("one\ntwo\nthree").Exec("grep one")
	p.Stdout()
}

func yqPipe() {
	yml := `
applications:
  - name: mcs-devops-2
    version: v0.12.0
    repo: finance/marketplace/experiments/mcs-devops-2
    envs: [ dev ]
    services:
      - type: mongo
        name: mcs-devops-2
      - postgre
      - rabbit
      - redis
`
	currentVersion, _ := script.Echo(yml).Exec("yq  '(.applications.[] | select(.name == \"mcs-devops-2\")).version'").String()
	fmt.Println("current version: ", currentVersion)
	s, _ := script.Echo(yml).Exec("yq  '(.applications.[] | select(.name == \"mcs-devops-2\")).version = \"v0.13.0\"'").String()
	fmt.Println(">>> result:\n" + s)
}
