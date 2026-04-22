package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My simple application")

	label := widget.NewLabel("Hello")
	button := widget.NewButton("Hit", func() {
		label.SetText("Good")
	})

	myWindow.SetContent(container.NewVBox(
		label,
		button,
	))

	myWindow.ShowAndRun()
}

