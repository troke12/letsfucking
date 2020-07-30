package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Lets Fucking Gooo")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Lets Fucking Go"),
	))

	w.ShowAndRun()
}
