package main

import (
	"honnef.co/go/js/dom"
)

func main() {
	d := dom.GetWindow().Document()

	clickme := d.GetElementByID("clickme").(*dom.HTMLButtonElement)
	display := d.GetElementByID("display").(*dom.HTMLDivElement)

	clickme.AddEventListener("click", false, func(event dom.Event) {
		display.SetInnerHTML("nandakumar")
	})
}
