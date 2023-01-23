package main

import (
	"github.com/ysrckr/simpleblog_api/routes"
)

func main() {
	

	app := routes.CreateServer()
	routes.ServeRoutes(app)
}
