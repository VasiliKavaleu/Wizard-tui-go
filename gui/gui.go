package gui

import (
	"github.com/rivo/tview"
	"os"
)


type Gui struct {
	app   *tview.Application
	pages *tview.Pages
	editServer bool
}

func New() *Gui {
	return &Gui{
		app:   tview.NewApplication(),
	}
}

func (g *Gui) Start() error {
	g.initMenu()
	if err := g.app.Run(); err != nil {
		g.app.Stop()
		return err
	}

	return nil
}

func (g *Gui) initMenu() {
	commandList := createCommandList()

	commandList.AddItem("Create server configuration", "", '1', func(){
		g.editServer = false
		g.drawServerConfForm()
	})
	
	commandList.AddItem("Quit", "", 'q', g.drawQuitNotifyForm)
	if _, err := os.Stat(serverFilePath); err == nil {
		commandList.InsertItem(1, "Change server configuration", "", '2', func(){
			g.editServer = true
			g.drawServerConfForm()
		})
	  }

	layout := createMainLayout(commandList)

	g.pages = tview.NewPages().AddAndSwitchToPage("main", layout, true)
	g.app.SetRoot(g.pages, true)
}

func (g *Gui) drawQuitNotifyForm() {
	modal:= tview.NewModal().
			SetText(quitMsg).
			AddButtons([]string{"Quit", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Quit" {
					g.app.Stop()
				} else {
					g.pages.RemovePage("quitNotify").ShowPage("main")
				}
			})
	g.pages.AddAndSwitchToPage("quitNotify", modal, true)
}

func (g *Gui) drawOkCancelNotifyForm(notifyMsg, okTitleBtn, cancelTitleBtn, showCurrentForm, showNextForm string) {
	modal:= tview.NewModal().
			SetText(notifyMsg).
			AddButtons([]string{okTitleBtn, cancelTitleBtn}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == cancelTitleBtn {
					g.pages.RemovePage("okCancelNotify").ShowPage(showCurrentForm)
				} else {
					g.pages.RemovePage("okCancelNotify").SwitchToPage(showNextForm)
				}
			})
	g.pages.AddAndSwitchToPage("okCancelNotify", modal, true)
}

func (g *Gui) drawNotifyMsgOkForm(msg string, showCurrentForm string) {
	modal:= tview.NewModal().
			SetText(msg).
			AddButtons([]string{"Ok"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Ok" {
					g.pages.RemovePage("notifyMsgOk")
				}
			})
	g.pages.AddAndSwitchToPage("notifyMsgOk", modal, true).ShowPage(showCurrentForm)
}

func (g *Gui) modal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}


func createMainLayout(commandList tview.Primitive) (layout *tview.Frame) {
	///// Main Layout /////
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

	clayout := tview.NewFrame(flex).SetBorders(4, 2, 8, 8, 0, 0)

	return clayout
}

func createCommandList() (commandList *tview.List) {
	///// Commands /////
	commandList = tview.NewList()                         
	commandList.SetBorder(true).SetTitle(mainMenuTitle) 
	commandList.ShowSecondaryText(false)
	return commandList
}
