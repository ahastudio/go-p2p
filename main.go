package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ahastudio/go-p2p/upnp"
)

const (
	Port = 8080
)

func main() {
	nat, error := upnp.Discover()
	if error != nil {
		fmt.Printf("Cannot discover UPnP: %s\n", error)
		os.Exit(3)
	}
	ip, error := nat.GetExternalAddress()
	if error != nil {
		fmt.Printf("Cannot get external address: %s\n", error)
		os.Exit(3)
	}
	port, error := nat.AddPortMapping("tcp", Port, Port, "SAESUL listen port", 0)
	if error != nil {
		fmt.Printf("Cannot add port mapping: %s\n", error)
		os.Exit(3)
	}
	fmt.Printf("External IP: %s, Port: %d\n", ip, port)
	fmt.Printf("http://%s:%d\n", ip, port)
	time.Sleep(5 * time.Second)
	error = nat.DeletePortMapping("tcp", Port, Port)
	if error != nil {
		fmt.Printf("Cannot delete port mapping: %s\n", error)
		os.Exit(3)
	}
}
