package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
)

var port = flag.Int("p", 1965, "Port to bind the server to. Defaults to 1965")
var certFile = flag.String("cert", "cert.pem", "TLS certificate to run the server with. Can be self signed.")
var keyFile = flag.String("key", "key.pem", "TLS private key to run the server with.")

func main() {
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("Could not load certificate: %s\n", err)
	}
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}
	// Listen on the port passed in (default 1965) for IPv4 connections
	// on any IP address of the local machine
	listener, err := tls.Listen("tcp4", fmt.Sprintf(":%d", *port), &config)
	if err != nil {
		log.Fatalf("Error opening listener: %s\n", err)
	}
	defer listener.Close()
	for {
		// Wait for a connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error opening connection: %s\n", err)
			break
		}
		log.Printf("Connection established from %s\n", conn.RemoteAddr())
		// Handle the connection in a new goroutine so the loop can go back to
		// accepting connections
		go func(c net.Conn) {
			bytes := make([]byte, 1024, 1024)
			n, err := c.Read(bytes)
			if err != nil {
				log.Println(err)
			}
			log.Printf("Bytes Read: %d\n", n)
			log.Println(string(bytes))
			c.Close()
		}(conn)
	}
}
