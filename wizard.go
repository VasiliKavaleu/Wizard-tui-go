package main

import (
	"configurator/gui"
	"log"
	"os"

	"runtime"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func run() int {
	if runtime.GOOS == "windows" {
		tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack
		tview.Styles.ContrastBackgroundColor = tcell.ColorLightGrey
		tview.Styles.MoreContrastBackgroundColor = tcell.ColorYellow
		tview.Styles.BorderColor = tcell.ColorYellow
		tview.Styles.TitleColor = tcell.ColorYellow
	} else {
		tview.Styles.PrimitiveBackgroundColor = tcell.ColorDarkSlateGrey
		tview.Styles.ContrastBackgroundColor = tcell.ColorDimGrey
		tview.Styles.MoreContrastBackgroundColor = tcell.ColorTan
		tview.Styles.BorderColor = tcell.ColorTan
		tview.Styles.TitleColor = tcell.ColorLightGoldenrodYellow
	}

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
