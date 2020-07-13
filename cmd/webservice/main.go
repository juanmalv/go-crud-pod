package main

import (
	"github.com/juanmalv/go-crud-pod/cmd/webservice/routes"
)

func main() {
	apiRouter := routes.GetRouter(routes.GinRouting)

	apiRouter.Run()
}
