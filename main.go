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
	hello_w := widget.Label{Text: " Hello Fyne"}
	widget_button := widget.NewButton("Hi!", func() { hello_w.SetText("Welcome :)") })
	vboxCntnr := container.NewVBox(widget_button)
	w.SetContent(vboxCntnr)
	w.SetCloseIntercept(closeOn)
	w.ShowAndRun()

}
func closeOn() {
	fmt.Printf("close button clicked")
}
