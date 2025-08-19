package models

import (
	"log"
	"net"
	"time"
)

func checkInternetConnection() bool {
	timeout := 3 * time.Second
	// Attempt to dial a connection to a reliable server. Google's DNS is a good choice.
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", timeout)
	if err != nil {
		log.Printf("No internet connection: %v\n", err)
		return false
	}
	defer conn.Close()
	return true
}
