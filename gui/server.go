package gui

import (
	"configurator/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
	"strings"
)

// Main form consisting all input field and DropDown filds for moving other forms
func (g *Gui) drawServerConfForm() {
	initialCnf := utils.ServerConfig{}

	if g.editServer {
		yfile, _ := ioutil.ReadFile(serverFilePath)
		_ = yaml.Unmarshal(yfile, &initialCnf)

	}

	cnf := &initialCnf

	form := tview.NewForm()

	form.AddInputField(ServerInputLabel["mediaStorages"], utils.ListToStr(cnf.Server.Media.Storages), inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["mediaThreads"], strconv.Itoa(cnf.Server.Media.Threads), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["mediaStreams"], utils.ListToStr(cnf.Server.Media.Streams), inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["tokenTtl"], strconv.Itoa(cnf.Server.Token.Ttl), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["tokenSecret"], cnf.Server.Token.Secret, inputWidth, nil, nil)

	form.AddDropDown(ServerInputLabel["broadcastSsl"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Server.Broadcast.Ssl),
		func(option string, optionIndex int) {
			if optionIndex != 0 {
				cnf.Server.Broadcast.Ssl = true
			} else {
				cnf.Server.Broadcast.Ssl = false
			}
		})

	form.AddInputField(ServerInputLabel["broadcastWhitelist"], utils.ListToStr(cnf.Server.Broadcast.Whitelist), inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["webThreads"], strconv.Itoa(cnf.Server.Broadcast.Web.Threads), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["webListen"], cnf.Server.Broadcast.Web.Listen, inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["rtspThreads"], strconv.Itoa(cnf.Server.Broadcast.Rtsp.Threads), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["rtspListen"], cnf.Server.Broadcast.Rtsp.Listen, inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["publishThreads"], strconv.Itoa(cnf.Server.Broadcast.Publish.Threads), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["publishListen"], cnf.Server.Broadcast.Publish.Listen, inputWidth, nil, nil)

	form.AddInputField(ServerInputLabel["controllerEvents"], utils.ListToStr(cnf.Controller.Events), inputWidth, nil, func(text string) {
		g.drawcontrollerEventsForm(form, cnf)
	})

	form.AddDropDown(ServerInputLabel["controllerDsn"], controllerDsnOptions, utils.GetIndexFromVal(controllerDsnOptions, cnf.Controller.Dsn),
		func(option string, optionIndex int) {
			cnf.Controller.Dsn = option
		})

	form.AddDropDown(ServerInputLabel["apiUsage"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Api.Enable), nil)
	form.AddDropDown(ServerInputLabel["cpanelUsage"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Cpanel.Enable), nil)
	form.AddDropDown(ServerInputLabel["clusterUsage"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Cluster.Enable), nil)

	exit := func() {
		g.drawOkCancelNotifyForm(notSaveMsg, "Ok", "Cancel", serverFormId, "main")
	}

	saveServerConf := func() {
		if g.validateSaveServerConf(form, cnf) {
			g.drawSaveServerNotifyForm(cnf)
		}
	}

	form.AddButton("Save", saveServerConf)
	form.AddButton("Cancel", exit)
	form.SetCancelFunc(exit)
	form.SetButtonsAlign(tview.AlignRight)
	form.SetBorder(true).SetTitle(labelServerForm)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlS:
			saveServerConf()
		}
		return event
	})

	g.addPopUpServerHandler(cnf, form)
	createSereverConfigPage := createFormLayout(form)
	g.pages.AddAndSwitchToPage(serverFormId, createSereverConfigPage, true)
}

// addPopUpHandler adds handler to some DropDown field which show forms depending on input value
func (g *Gui) addPopUpServerHandler(cnf *utils.ServerConfig, form *tview.Form) {
	apiUsageDropDown := form.GetFormItemByLabel(ServerInputLabel["apiUsage"]).(*tview.DropDown)
	apiUsageDropDown.SetSelectedFunc(func(text string, index int) {
		if index != 0 {
			cnf.Api.Enable = true
			g.drawApiConfForm(cnf, form)
		} else {
			cnf.Api.Enable = false
		}
	})

	cpanelUsageDropDown := form.GetFormItemByLabel(ServerInputLabel["cpanelUsage"]).(*tview.DropDown)
	cpanelUsageDropDown.SetSelectedFunc(func(text string, index int) {
		if index != 0 {
			cnf.Cpanel.Enable = true
			g.drawCpanelConfForm(cnf, form)
		} else {
			cnf.Cpanel.Enable = false
		}
	})

	clusterUsageDropDown := form.GetFormItemByLabel(ServerInputLabel["clusterUsage"]).(*tview.DropDown)
	clusterUsageDropDown.SetSelectedFunc(func(text string, index int) {
		if index != 0 {
			cnf.Cluster.Enable = true
			g.drawClusterConfForm(cnf, form)
		} else {
			cnf.Cluster.Enable = false
		}
	})
}

func (g *Gui) validateSaveServerConf(form *tview.Form, cnf *utils.ServerConfig) bool {

	mediaStorages := form.GetFormItemByLabel(ServerInputLabel["mediaStorages"]).(*tview.InputField).GetText()
	mediaThreads := form.GetFormItemByLabel(ServerInputLabel["mediaThreads"]).(*tview.InputField).GetText()
	mediaStreams := form.GetFormItemByLabel(ServerInputLabel["mediaStreams"]).(*tview.InputField).GetText()
	tokenTtl := form.GetFormItemByLabel(ServerInputLabel["tokenTtl"]).(*tview.InputField).GetText()
	tokenSecret := form.GetFormItemByLabel(ServerInputLabel["tokenSecret"]).(*tview.InputField).GetText()
	broadcastWhitelist := form.GetFormItemByLabel(ServerInputLabel["broadcastWhitelist"]).(*tview.InputField).GetText()
	webThreads := form.GetFormItemByLabel(ServerInputLabel["webThreads"]).(*tview.InputField).GetText()
	webListen := form.GetFormItemByLabel(ServerInputLabel["webListen"]).(*tview.InputField).GetText()
	rtspThreads := form.GetFormItemByLabel(ServerInputLabel["rtspThreads"]).(*tview.InputField).GetText()
	rtspListen := form.GetFormItemByLabel(ServerInputLabel["rtspListen"]).(*tview.InputField).GetText()
	publishThreads := form.GetFormItemByLabel(ServerInputLabel["publishThreads"]).(*tview.InputField).GetText()
	publishListen := form.GetFormItemByLabel(ServerInputLabel["publishListen"]).(*tview.InputField).GetText()

	mediaStorages = strings.TrimSpace(mediaStorages)
	mediaThreads = strings.TrimSpace(mediaThreads)
	mediaStreams = strings.TrimSpace(mediaStreams)
	tokenTtl = strings.TrimSpace(tokenTtl)
	tokenSecret = strings.TrimSpace(tokenSecret)
	broadcastWhitelist = strings.TrimSpace(broadcastWhitelist)
	webThreads = strings.TrimSpace(webThreads)
	webListen = strings.TrimSpace(webListen)
	rtspThreads = strings.TrimSpace(rtspThreads)
	rtspListen = strings.TrimSpace(rtspListen)
	publishThreads = strings.TrimSpace(publishThreads)
	publishListen = strings.TrimSpace(publishListen)

	if g.checkAndNotifyRequiredField(mediaThreads, ServerInputLabel["mediaThreads"], serverFormId) &&
		g.checkAndNotifyRequiredField(mediaStreams, ServerInputLabel["mediaStreams"], serverFormId) &&
		g.checkAndNotifyRequiredField(tokenTtl, ServerInputLabel["tokenTtl"], serverFormId) &&
		g.checkAndNotifyRequiredField(tokenSecret, ServerInputLabel["tokenSecret"], serverFormId) &&
		g.checkAndNotifyRequiredField(broadcastWhitelist, ServerInputLabel["broadcastWhitelist"], serverFormId) &&
		g.checkAndNotifyRequiredField(webThreads, ServerInputLabel["webThreads"], serverFormId) &&
		g.checkAndNotifyRequiredField(webListen, ServerInputLabel["webListen"], serverFormId) &&
		g.checkAndNotifyRequiredField(rtspThreads, ServerInputLabel["rtspThreads"], serverFormId) &&
		g.checkAndNotifyRequiredField(rtspListen, ServerInputLabel["rtspListen"], serverFormId) &&
		g.checkAndNotifyRequiredField(publishThreads, ServerInputLabel["publishThreads"], serverFormId) &&
		g.checkAndNotifyRequiredField(publishListen, ServerInputLabel["publishListen"], serverFormId) {

		cnf.Server.Media.Storages = utils.StrToList(mediaStorages)
		cnf.Server.Media.Threads, _ = strconv.Atoi(mediaThreads)
		cnf.Server.Media.Streams = utils.StrToList(mediaStreams)
		cnf.Server.Token.Ttl, _ = strconv.Atoi(tokenTtl)
		cnf.Server.Token.Secret = tokenSecret
		cnf.Server.Broadcast.Whitelist = utils.StrToList(broadcastWhitelist)
		cnf.Server.Broadcast.Web.Listen = webThreads
		cnf.Server.Broadcast.Web.Listen = webListen
		cnf.Server.Broadcast.Rtsp.Threads, _ = strconv.Atoi(rtspThreads)
		cnf.Server.Broadcast.Rtsp.Listen = rtspListen
		cnf.Server.Broadcast.Publish.Threads, _ = strconv.Atoi(publishThreads)
		cnf.Server.Broadcast.Publish.Listen = publishListen

	} else {
		return false
	}

	if cnf.Controller.Dsn == "" {
		msg := utils.GetReqFieldMsg(ServerInputLabel["controllerDsn"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	}

	if !cnf.Api.Enable {
		cnf.Api = utils.Api{}
	}

	if !cnf.Cpanel.Enable {
		cnf.Cpanel = utils.Cpanel{}
	}

	if !cnf.Cluster.Enable {
		cnf.Cluster = utils.Cluster{}
	}

	return true
}

// form reflecting checkboxes
func (g *Gui) drawcontrollerEventsForm(parentForm *tview.Form, cnf *utils.ServerConfig) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle(" Choose controller events ")
	events := map[string]bool{}

	form.AddCheckbox(controllerEventsOptions[0], false, func(checked bool) {
		if checked {
			events[controllerEventsOptions[0]] = checked
		} else {
			delete(events, controllerEventsOptions[0])
		}
	})
	form.AddCheckbox(controllerEventsOptions[1], false, func(checked bool) {
		if checked {
			events[controllerEventsOptions[1]] = checked
		} else {
			delete(events, controllerEventsOptions[1])
		}
	})
	form.AddCheckbox(controllerEventsOptions[2], false, func(checked bool) {
		if checked {
			events[controllerEventsOptions[2]] = checked
		} else {
			delete(events, controllerEventsOptions[2])
		}
	})
	form.AddCheckbox(controllerEventsOptions[3], false, func(checked bool) {
		if checked {
			events[controllerEventsOptions[3]] = checked
		} else {
			delete(events, controllerEventsOptions[3])
		}
	})

	exit := func() {
		eventsList := []string{}
		for key := range events {
			eventsList = append(eventsList, key)
		}
		cnf.Controller.Events = eventsList

		controllerEventsInput := parentForm.GetFormItemByLabel(ServerInputLabel["controllerEvents"]).(*tview.InputField)
		controllerEventsInput.SetText(utils.ListToStr(eventsList))
		g.app.SetFocus(controllerEventsInput)

		g.pages.RemovePage("controllerEvents").ShowPage("main")
	}

	form.AddButton("Ok", exit)
	form.SetButtonsAlign(tview.AlignRight)
	form.SetCancelFunc(exit)

	g.pages.AddAndSwitchToPage("controllerEvents", g.modal(form, 40, 13), true).ShowPage("serverForm")
}

// Form for confirmation of saving the main form
func (g *Gui) drawSaveServerNotifyForm(cnf *utils.ServerConfig) {
	modal := tview.NewModal().
		SetText(saveMsg).
		AddButtons([]string{"Save", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
				g.pages.RemovePage("saveServerNotify")
			} else {
				g.pages.RemovePage("saveServerNotify")
				utils.SaveConfigToFile(cnf, serverFilePath)
				g.addChangeServerMenuItem()
				g.drawNotifyMsgOkForm(saveSuccessConfMsg, "main")
			}
		})
	g.pages.AddAndSwitchToPage("saveServerNotify", modal, true).ShowPage(serverFormId)
}
