package main

import (
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
)

func main() {
	// setup cors plugin
	plugin.Register(cors.NewPlugin())

	// init command
	cmd.Init()
}
