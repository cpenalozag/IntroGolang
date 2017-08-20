package main

import (
	"os"
	"time"
	"util"
)

func main() {
	listenPort := "8080"
	args := os.Args
	if len(args) > 1 {
		hostIP := args[1]
		util.RunGuest(hostIP)
	} else {
		listenIP := util.GetLocalNetworkIP()
		util.RunHost(listenIP + ":" + listenPort)
	}
}

func getListenPort() string {
	hourOfDay := time.Now().Hour()
	if hourOfDay < 12 {
		return "8080"
	} else if hourOfDay < 20 {
		return "8081"
	} else {
		return "8082"
	}
}
