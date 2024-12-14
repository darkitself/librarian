package main

import "librarian/network"

func main() {
	network.NewServer(9999).Start()
}
