package main

import (
	"github.com/Bios-Marcel/tview"
	"strings"
)

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		}).
		SetScrollable(true).
		SetText(strings.Repeat("OwO\n", 100)).
		SetBorderSides(true, true, true, false).
		SetBorder(true)
	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
