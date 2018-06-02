package main

import (
	"log"
	"net/rpc"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// General Config
	kingpin.Version("0.0.1")
	kingpin.Command("reload", "reload something")
	kingpin.Command("do", "do something")

	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// Parse Commands
	switch kingpin.Parse() {

	case "reload":
		var reply string
		client.Call("CliSrv.ReloadSomething", "something", &reply)
		log.Println(reply)

	case "do":
		var reply bool
		client.Call("CliSrv.DoSomething", true, &reply)
		if reply {
			log.Println("done")
			return
		}
		log.Println("Command failed")
	}
}
