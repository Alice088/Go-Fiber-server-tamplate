package fiber_builder

func New() {
	d := &Director{
		Builder: &Builder{
			Server: &Server{},
		},
	}
	
	d.Construct()
}
