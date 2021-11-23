package gui

import (
	"github.com/rivo/tview"
	"os"
)


var inputWidth = 50

var mainMenuTitle = " Select an option "
var ValidationTitle = "Check input value"
var quitMsg = "Exit? Server configuration will be completed."
var saveSuccessConfMsg = "Configuration file saved successfully!"
var labelServerForm = " Fill out server configuration form "
var serverFormId = "serverForm"
var apiFormId = "apiForm"
var cpanelFormId = "cpanelForm"
var usageDropDownOptions = []string{"Disable", "Enable"}
var controllerDsnOptions = []string{1: "http://", 2: "ws://", 3: "mysql://", 4: "rabbitmq://", 5: "tarantool://"}
var apiAuthDropDownOptions = []string{"none", "basic", "digest", "token"}
var cpanelAuthDropDownOptions = []string{"none", "basic", "digest"}
var controllerEventsOptions = []string{"up", "state", "stream", "cluster"}
var serverFilePath = "server.yaml"
var passwordMask = '*'

var ServerInputLabel = map[string]string{
	"mediaThreads": "Number of mediaserver threads",
	"mediaStreams": "List paths or file streams configuration",
	"mediaStorages": "List mediaserver storages",

	"rtspThreads": "Number of RTSP threads",
	"rtspListen": "RTSP socket",
	"publishThreads": "Number of publish threads",
	"publishListen": "Publish socket",
	"webThreads": "Number of web threads",
	"webListen": "Web socket",

	"broadcastSsl": "Broadcast SSL",
	"broadcastWhitelist": "Broadcast whitelist",

	"tokenSecret": "Token secret",
	"tokenTtl": "Token TTL",
	"controllerDsn": "Сontroller DSN",
	"controllerEvents": "Сontroller events",
	
	"apiUsage": "API usage",
	"apiListen": "Socket",
	"apiModule": "Module",
	"apiSsl": "SSL",
	"apiWhitelist": "Whitelist",
	"apiAuth": "Authorization",
	"apiUsersAdmin": "Password for admin",
	"apiUsersRoot": "Password for root",

	"cpanelUsage": "Cpanel usage",
	"cpanelListen": "Socket",
	"cpanelModule": "Module",
	"cpanelSsl": "SSL",
	"cpanelWhitelist": "Whitelist",
	"cpanelAuth": "Authorization",
	"cpanelUsersAdmin": "Password for admin",
	"cpanelUsersRoot": "Password for root",
	"cpanelUsersUser": "Password for user",
	"cpanelUsersGuest": "Password for guest",

	"clusterUsage": "Cluster usage",
	"clusterId": "Cluster usage",
	"clusterNode": "Cluster node",
	"clusterPool": "Cluster pool",
	"clusterWarmingUp": "Warming up time",
	"clusterRetries": "Number of retries",
	"clusterInterval": "Interval for reconnect",

}


type Gui struct {
	app   *tview.Application
	pages *tview.Pages
	editServer bool
}

// New create new gui
func New() *Gui {
	return &Gui{
		app:   tview.NewApplication(),
	}
}

// Start application
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

func (g *Gui) drawNotifyMsgOkForm(msg string, showPage string) {
	modal:= tview.NewModal().
			SetText(msg).
			AddButtons([]string{"Ok"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Ok" {
					g.pages.RemovePage("notifyMsgOk").ShowPage(showPage)
				}
			})
	g.pages.AddAndSwitchToPage("notifyMsgOk", modal, true).ShowPage(showPage)
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
	info.SetText("Mediaserver Wizard v1.0 - Copyright 2021 NavekSoft") 
	info.SetTextAlign(tview.AlignCenter)
	
	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow). 
	AddItem(commandList, 10, 1, true) // 10, 1

	// inner_layout := tview.NewFlex().SetDirection(tview.FlexRow).
	// 	AddItem(mainLayout, 0, 1, true).
	// 	AddItem(info, 3, 1, false)

	flex := tview.NewFlex().
	AddItem(tview.NewBox(), 0, 1, false).
	AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetTitle("Top"), 0, 1, false).
		AddItem(mainLayout, 0, 3, true).
		AddItem(info, 3, 1, false), 0, 2, true).
	AddItem(tview.NewBox(), 0, 1, false)

	clayout := tview.NewFrame(flex).SetBorders(4, 2, 8, 8, 12, 12)

	return clayout
}

func createCommandList() (commandList *tview.List) {
	///// Commands /////
	commandList = tview.NewList()                         
	commandList.SetBorder(true).SetTitle(mainMenuTitle) 
	commandList.ShowSecondaryText(false)
	return commandList
}
