package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	js := mewn.String("./frontend/build/main.js")
	css := mewn.String("./frontend/build/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  512,
		Height: 512,
		Title:  "Hidrate-se!",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	//app.Bind()
	app.Run()
}
