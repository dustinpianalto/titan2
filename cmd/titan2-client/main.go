package main

import (
	"crypto/tls"
	"log"
)

func main() {
	config := &tls.Config{
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp4", "localhost:1965", config)
	if err != nil {
		log.Fatal(err)
	}
	n, err := conn.Write([]byte("Test message"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
}
