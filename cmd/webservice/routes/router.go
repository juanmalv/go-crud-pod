package routes

type Router interface {
	Configure()
	Register()
	Run() error
}

//GetRouter allows us to implement a different routing approach.
func GetRouter(routingMethod string) Router {
	if routingMethod != GinRouting {
		return nil
	}

	//Para tener en cuenta, este ginRouter se moverá al heap para que no se pierda en el limpiado del stack.
	var ginRouter GinRouter
	return &ginRouter
}