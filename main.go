package main

import (
	"github.com/Nivelian/ingrid-task/api"
	"github.com/Nivelian/ingrid-task/helpers"
)

const port = ":8080"

func main() {
	helpers.InitLog()

	api.StartServer(port)
}
