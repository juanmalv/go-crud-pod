package main

import "github.com/juanmalv/go-crud-pod/cmd/webservice/routes"

func main() {
	var ginRouter routes.GinRouter

	ginRouter.Run()
}
