package main

import "awesomeProject3/internal/pkg/app"

const cfgPath = "../../config/config.yaml"

func main() {
	app := app.New(cfgPath)
	app.Start()
}
