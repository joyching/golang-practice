package main

import (
	"github.com/joyching/golang-practice/config"
	"github.com/joyching/golang-practice/server"
)

func main() {
	config.Init()

	server.Init()
}
