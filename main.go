package main

import "Vectorquest/server"

func main() {
	err := server.Start()
	if err != nil {
		return
	}
}
