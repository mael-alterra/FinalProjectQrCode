package main

import (
	m "FinalProjectQrCode/middleware"
	"FinalProjectQrCode/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":" + port))
}
