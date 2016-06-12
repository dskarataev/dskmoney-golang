package main

import (
	"dskmoney-golang/dskmoney"
)

func main() {
	app := dskmoney.NewApp()

	if err := app.Init(); err != nil {
		panic("Init error: " + err.Error())
	}

	if err := app.Run(); err != nil {
		panic("Run error: " + err.Error())
	}
}