package gui

import (
	"configurator/utils"
	"github.com/rivo/tview"
	"os"
)

type Gui struct {
	app        *tview.Application
	pages      *tview.Pages
	editServer bool
	menuList   *tview.List
}

func New() *Gui {
	return &Gui{
		app: tview.NewApplication(),
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
	g.menuList = createCommandList()

	g.menuList.AddItem(menuLabes["createServer"], "", '1', func() {
		g.editServer = false
		g.drawServerConfForm()
	})

	g.menuList.AddItem(menuLabes["createStream"], "", '2', func() {
		g.drawStreamConfForm()
	})

	g.menuList.AddItem(menuLabes["exit"], "", 'q', g.drawQuitNotifyForm)
	g.addChangeServerMenuItem()

	layout := createMainLayout(g.menuList)

	g.pages = tview.NewPages().AddAndSwitchToPage("main", layout, true)
	g.app.SetRoot(g.pages, true)
}

func (g *Gui) addChangeServerMenuItem() {
	if _, err := os.Stat(serverFilePath); err == nil {
		if g.menuList.GetItemCount() == 3 {
			g.menuList.InsertItem(2, menuLabes["changeServer"], "", '3', func() {
				g.editServer = true
				g.drawServerConfForm()
			})
		}
	}
}

func (g *Gui) drawQuitNotifyForm() {
	modal := tview.NewModal().
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
	modal := tview.NewModal().
		SetText(notifyMsg).
		AddButtons([]string{okTitleBtn, cancelTitleBtn}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == cancelTitleBtn {
				g.pages.RemovePage("okCancelNotify").SwitchToPage(showCurrentForm)
			} else {
				g.pages.RemovePage("okCancelNotify").SwitchToPage(showNextForm)
			}
		})
	g.pages.AddAndSwitchToPage("okCancelNotify", modal, true)
}

func (g *Gui) drawNotifyMsgOkForm(msg string, whichFormShow string) {
	modal := tview.NewModal().
		SetText(msg).
		SetTextColor(requiredMSGColor).
		AddButtons([]string{"Ok"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Ok" {
				g.pages.RemovePage("notifyMsgOk")
			}
		})
	g.pages.AddAndSwitchToPage("notifyMsgOk", modal, true).ShowPage(whichFormShow)
}

// Verification notification form of mandatory field
func (g *Gui) checkAndNotifyRequiredField(inputValue, inputLabel, whichFormShow string) bool {
	if !utils.ValidReqField(inputValue) {
		msg := utils.GetReqFieldMsg(inputLabel)
		g.drawNotifyMsgOkForm(msg, whichFormShow)
		return false
	}
	return true
}

func (g *Gui) modal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}
