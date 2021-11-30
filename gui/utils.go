package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func createMainLayout(commandList tview.Primitive) (layout *tview.Frame) {
	info := tview.NewTextView()
	info.SetBorder(true)
	info.SetText("Mediaserver Wizard v1.0")
	info.SetTextAlign(tview.AlignCenter)

	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(commandList, 10, 1, true)

	flex := tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetTitle("Top"), 0, 1, false).
			AddItem(mainLayout, 0, 3, true).
			AddItem(info, 3, 1, false), 0, 2, true).
		AddItem(tview.NewBox(), 0, 1, false)

	layout = tview.NewFrame(flex).SetBorders(4, 2, 8, 8, 0, 0)
	return
}

func createCommandList() (commandList *tview.List) {
	commandList = tview.NewList()
	commandList.SetBorder(true).SetTitle(mainMenuTitle)
	commandList.ShowSecondaryText(false)
	return commandList
}

func createFormLayout(form tview.Primitive) tview.Primitive {
	layout := tview.NewFlex().SetDirection(tview.FlexRow)
	layout.AddItem(form, 0, 2, true)
	layout.AddItem(tview.NewTextView().
		SetText(navigate).
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorSpringGreen), 2, 1, false)
	return layout
}
