package main

import (
	"gizumon.com/life-manager-api/routes"
)

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}

func hello() string {
	return "Hello, World!"
}
