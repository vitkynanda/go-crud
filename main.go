package main

import (
	"go-api/connection"
	"go-api/routes"
)

func main() {
	connection.Connect()
	routes.HandlerRequest()
}