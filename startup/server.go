package startup

import (
	fiber_builder2 "RuRu/internal/server_builder"
)

func Server() {
	d := &fiber_builder2.Director{
		Builder: &fiber_builder2.Builder{
			Server: &fiber_builder2.Server{},
		},
	}
	d.Construct()
}
