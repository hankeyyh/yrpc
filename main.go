package main

import "github.com/hankeyyh/yrpc/pkg/application"

func main() {
	app := application.Application{}

	if err := app.Init(); err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
