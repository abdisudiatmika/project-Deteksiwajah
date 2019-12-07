package main

import (
	"go-training-restful/routes"
)

func main() {
	e := routes.New()

	e.Logger.Fatal(e.Start(":80000"))
}
