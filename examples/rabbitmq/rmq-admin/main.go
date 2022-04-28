package main

import (
	"flag"
	"github.com/michaelklishin/rabbit-hole/v2"
	"log"
)

func main() {

	url := flag.String("url", "", "rabbit server url")
	username := flag.String("username", "guest", "")
	password := flag.String("password", "guest", "")

	flag.Parse()

	//rmqExamples(*url, *username, *password)
	createVhostAndUser(*url, *username, *password)
}

func rmqExamples(url string, username string, password string) {
	rmqc, err := rabbithole.NewClient(url, username, password)
	if err != nil {
		log.Fatalf("Error create client %v", err)
	}

	resp, err := rmqc.Overview()
	if err != nil {
		log.Fatalf("Error overview %v", err)
	}

	log.Printf("rmq overview: %v\n", resp)

	vhosts, err := rmqc.ListVhosts()
	if err != nil {
		log.Fatalf("Error get vhosts %v", err)
	}
	log.Printf("vhosts: %#v/n", vhosts)

	vh, err := rmqc.GetVhost("/")
	if err != nil {
		log.Fatalf("Error get vhost %v", err)
	}
	log.Printf("vhost: %#v/n", vh)

	ul, err := rmqc.ListUsers()
	// => []UserInfo, err
	log.Printf("users: %#v/n", ul)

	// information about individual user
	//x, err := rmqc.GetUser("my.user")
}

func createVhostAndUser(url string, username string, password string) {
	rmqc, err := rabbithole.NewClient(url, username, password)
	if err != nil {
		log.Fatalf("Error create client %v", err)
	}
	vhost := "prod"
	ensureVhost(rmqc, vhost)

	newUsername := "devops-mcs"
	newPassword := "1q2w3e4r"
	ensureUser(rmqc, vhost, newUsername, newPassword)
}

func ensureUser(rmqc *rabbithole.Client, vhost string, username string, password string) {
	ul, err := rmqc.ListUsers()
	if err != nil {
		log.Fatalf("Error get users %v", err)
	}
	// => []UserInfo, err
	for _, ui := range ul {
		if ui.Name == username {
			log.Printf("user %s already exists!\n", username)
			return
		}
	}
	userTags := rabbithole.UserTags{"management", "policymaker"}
	resp, err := rmqc.PutUser(username, rabbithole.UserSettings{Password: password, Tags: userTags})
	if err != nil {
		log.Fatalf("Error create user %v", err)
	}
	log.Printf("resp: %#v\n", resp)

	resp, err = rmqc.UpdatePermissionsIn(vhost, username, rabbithole.Permissions{Configure: ".*", Write: ".*", Read: ".*"})
	if err != nil {
		log.Fatalf("Error update permissions for user %v", err)
	}
	log.Printf("resp: %#v\n", resp)

}

func ensureVhost(rmqc *rabbithole.Client, vhost string) {
	vhosts, err := rmqc.ListVhosts()
	if err != nil {
		log.Fatalf("Error get vhosts %v", err)
	}
	for _, vh := range vhosts {
		if vh.Name == vhost {
			log.Printf("vhost %s already exists\n", vhost)
			return
		}
	}
	resp, err := rmqc.PutVhost(vhost, rabbithole.VhostSettings{Tracing: false})
	if err != nil {
		log.Fatalf("Error vhost creation: %v", err)
	}
	log.Printf("resp: %#v\n", resp)

}
