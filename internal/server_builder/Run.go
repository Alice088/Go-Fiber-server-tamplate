package server_builder

func Run() {
	d := &Director{
		Builder: &Builder{
			Server: &Server{},
		},
	}

	d.Construct()
}
