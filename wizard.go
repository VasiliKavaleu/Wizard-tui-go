package main

import (
	"configurator/gui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"os"
	"runtime"
)

func run() int {
	if runtime.GOOS == "windows" {
		tview.Styles.PrimitiveBackgroundColor = tcell.ColorSlateGrey
		tview.Styles.ContrastBackgroundColor = tcell.ColorDimGrey
	} else {
		tview.Styles.ContrastBackgroundColor = tcell.ColorSlateGrey
	}
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorTan
	tview.Styles.BorderColor = tcell.ColorTan
	tview.Styles.TitleColor = tcell.ColorLightGoldenrodYellow

	gui := gui.New()

	if err := gui.Start(); err != nil {
		log.Fatal("Cannot start Wizard: %s", err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run())
}
