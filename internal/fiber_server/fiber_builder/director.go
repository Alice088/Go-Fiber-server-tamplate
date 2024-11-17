package fiber_builder

type IDirector interface {
	Construct()
}

type Director struct {
	Builder IBuilder
}

func (d Director) Construct() {
	d.Builder.InitFiber()
	d.Builder.InitLogger()
	d.Builder.InitSession()
	d.Builder.InitMiddleware()
	d.Builder.InitRoutes()
}
