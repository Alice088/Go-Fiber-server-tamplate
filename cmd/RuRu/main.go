package main

import (
	"RuRu/internal/fiber_server/fiber_builder"
	"RuRu/pkg/env"
)

func main() {
	env.Load()
	fiber_builder.Run()
}
