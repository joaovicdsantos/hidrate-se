package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	js := mewn.String("./frontend/build/main.js")
	css := mewn.String("./frontend/build/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  500,
		Height: 400,
		Title:  "Hidrate-se!",
		JS:     js,
		CSS:    css,
		Colour: "#2980b9",
	})
	app.Bind(&Contador{})
	app.Run()
}
