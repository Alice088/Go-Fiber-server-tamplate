package main

import (
	"app/pkg/env"
	"app/startup"
)

func main() {
	env.Load()
	startup.Server()
}
