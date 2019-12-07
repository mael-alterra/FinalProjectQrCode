package main

import (
	m "FinalProjectQrCode/middleware"
	"FinalProjectQrCode/routes"
)

func main() {
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
