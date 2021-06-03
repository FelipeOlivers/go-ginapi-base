package main

import "go-ginapi-base/server"

// Main
func main() {
	// Start Server
	s := server.StartServer()
	s.Run()
}
