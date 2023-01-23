package main

import "github.com/ysrckr/simpleblog_api/routes"

const PORT = ":4000"

func main() {

	app := routes.CreateServer()
	routes.ServeRoutes(app, PORT)
}
