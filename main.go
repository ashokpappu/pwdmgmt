package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	fmt.Printf("Hello World")
	a := app.New()
	w := a.NewWindow("Password Management Tool")
	var hello_w widget.Label
	hello_w.SetText("Welcome:")
	var widgetButton = widget.NewButton("Hi!", widgetButtonHiClicked)
	vboxCntnr := container.NewVBox(widgetButton)
	w.SetContent(vboxCntnr)
	w.SetCloseIntercept(closeOn)
	w.ShowAndRun()

}
func closeOn() {
	fmt.Printf("close button clicked")
}

func widgetButtonHiClicked() {

}
