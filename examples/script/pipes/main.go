package main

import (
	"fmt"
	"github.com/bitfield/script"
	"log"
)

func main() {
	echoPipe()
	yqPipe()
	pipeExitCode()
}

func echoPipe() {
	p := script.Echo("one\ntwo\nthree").Exec("grep one")
	p.Stdout()
}

func pipeExitCode() {
	p := script.Exec("ls").Exec("cat manifest.yml")
	fmt.Printf("ExitStatus: %d\n", p.ExitStatus())
	s, err := p.String()
	fmt.Printf("s: %s err: %v\n", s, err)
}

func updateVersion(manifest string, appName string, newVersion string) string {
	cmd := fmt.Sprintf("yq  '(.applications.[] | select(.name == \"%s\")).version'", appName)
	currentVersion, err := script.Echo(manifest).Exec(cmd).String()
	if err != nil {
		log.Fatalf("Error get current version: %v\n", err)
	}
	fmt.Println("current version: ", currentVersion)

	cmd = fmt.Sprintf("yq  '(.applications.[] | select(.name == \"%s\")).version =%s'", appName, newVersion)
	p := script.Echo(manifest)
	p = p.Exec(cmd)
	log.Printf("Error: %v ExitStatus: %v\n", p.Error(), p.ExitStatus())
	s, err := p.String()
	if err != nil {
		log.Printf("s: %s\n", s)
		log.Fatalf("Error update version: %v\n", err)
	}
	return s
}

func yqPipe() {
	yml := `
applications:
  - name: mcs-devops
    version: 1.6
    repo: finance/marketplace/experiments/mcs-devops
    envs:
      - dev
    services: []
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
  - name: dev-nginx
    version: "22"
    repo: finance/marketplace/authorization/dev-nginx
    envs:
      - dev
    services: []`

	r1 := updateVersion(yml, "mcs-devops", "1.7")
	fmt.Println(">>> result:\n" + r1)

	r2 := updateVersion(yml, "mcs-devops-2", "v0.13.0")
	fmt.Println(">>> result:\n" + r2)

	r3 := updateVersion(yml, "dev-nginx", "23")
	fmt.Println(">>> result:\n" + r3)
}
