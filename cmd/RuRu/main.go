package main

import (
	"RuRu/pkg/env"
	"RuRu/startup"
)

func main() {
	env.Load()
	startup.Server()
}
