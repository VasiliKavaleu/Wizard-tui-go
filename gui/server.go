package gui

import (
	"configurator/utils"
	"github.com/rivo/tview"
	"strconv"
	"strings"
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

func (g *Gui) drawServerConfForm() {
	initialCnf := utils.ServerConfig{}

	if g.editServer {
		// utils.ReadConfigFile(serverFilePath, initialCnf)
		yfile, _ := ioutil.ReadFile(serverFilePath)
		_ = yaml.Unmarshal(yfile, &initialCnf)
		
	} 
	
	// else {
	// 	initialCnf = utils.ServerConfig{}
	// }
	cnf := &initialCnf

	//g.drawNotifyMsgOkForm(utils.ListToStr(cnf.Server.Media.Storages), "")
	


	form := tview.NewForm()

	form.AddInputField(ServerInputLabel["mediaStorages"], utils.ListToStr(cnf.Server.Media.Storages), inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["mediaThreads"], strconv.Itoa(cnf.Server.Media.Threads), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["mediaStreams"], utils.ListToStr(cnf.Server.Media.Streams), inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["tokenTtl"], strconv.Itoa(cnf.Server.Token.Ttl), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["tokenSecret"], cnf.Server.Token.Secret, inputWidth, nil, nil)

	form.AddDropDown(ServerInputLabel["broadcastSsl"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Server.Broadcast.Ssl), 
					func(option string, optionIndex int){
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

	form.AddInputField(ServerInputLabel["controllerEvents"], utils.ListToStr(cnf.Controller.Events), inputWidth, nil, func(text string){
		g.drawcontrollerEventsForm(form, cnf)
	})


	form.AddDropDown(ServerInputLabel["controllerDsn"], controllerDsnOptions, utils.GetIndexFromVal(controllerDsnOptions, cnf.Controller.Dsn), 
					func(option string, optionIndex int){
							cnf.Controller.Dsn = option
					})

	form.AddDropDown(ServerInputLabel["apiUsage"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Api.Enable), 
					func(option string, optionIndex int){
						if optionIndex != 0 {
							cnf.Api.Enable = true
							g.drawApiConfForm(cnf, form)
						} else {
							cnf.Api.Enable = false
						}
					})
	
	form.AddDropDown(ServerInputLabel["cpanelUsage"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Cpanel.Enable), 
					func(option string, optionIndex int){
						if optionIndex != 0 {
							cnf.Cpanel.Enable = true
							g.drawCpanelConfForm(cnf, form)
						} else {
							cnf.Cpanel.Enable = false
						}
					})

	form.AddDropDown(ServerInputLabel["clusterUsage"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Cluster.Enable), 
					func(option string, optionIndex int){
						if optionIndex != 0 {
							cnf.Cluster.Enable = true
							g.drawClusterConfForm(cnf, form)
						} else {
							cnf.Cluster.Enable = false
						}
					})

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

	createSereverConfigPage := func(form tview.Primitive) tview.Primitive {
		createPage := tview.NewFlex().SetDirection(tview.FlexRow)
		createPage.AddItem(form, 0, 2, true)
		return createPage
	}(form)

	g.pages.AddAndSwitchToPage(serverFormId, createSereverConfigPage, true)
}


func (g *Gui) validateSaveServerConf(form *tview.Form, cnf *utils.ServerConfig) bool {
	mediaStorages := form.GetFormItemByLabel(ServerInputLabel["mediaStorages"]).(*tview.InputField).GetText()
	mediaStorages = strings.TrimSpace(mediaStorages)
	if !utils.ValidReqField(mediaStorages) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["mediaStorages"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	}

	mediaThreads := form.GetFormItemByLabel(ServerInputLabel["mediaThreads"]).(*tview.InputField).GetText()
	mediaThreads = strings.TrimSpace(mediaThreads)
	if !utils.ValidReqField(mediaThreads) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["mediaThreads"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	} 

	mediaStreams := form.GetFormItemByLabel(ServerInputLabel["mediaStreams"]).(*tview.InputField).GetText()
	mediaStreams = strings.TrimSpace(mediaStreams)
	if !utils.ValidReqField(mediaStreams) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["mediaStreams"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	} 

	tokenTtl := form.GetFormItemByLabel(ServerInputLabel["tokenTtl"]).(*tview.InputField).GetText()
	tokenTtl = strings.TrimSpace(tokenTtl)
	if !utils.ValidReqField(tokenTtl) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["tokenTtl"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	}

	tokenSecret := form.GetFormItemByLabel(ServerInputLabel["tokenSecret"]).(*tview.InputField).GetText()
	tokenSecret = strings.TrimSpace(tokenSecret)
	if !utils.ValidReqField(tokenTtl) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["tokenSecret"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	}

	broadcastWhitelist := form.GetFormItemByLabel(ServerInputLabel["broadcastWhitelist"]).(*tview.InputField).GetText()
	broadcastWhitelist = strings.TrimSpace(broadcastWhitelist)
	if !utils.ValidReqField(broadcastWhitelist) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["broadcastWhitelist"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	}

	webThreads := form.GetFormItemByLabel(ServerInputLabel["webThreads"]).(*tview.InputField).GetText()
	webThreads = strings.TrimSpace(webThreads)
	if !utils.ValidReqField(webThreads) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["webThreads"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false 
	}

	webListen := form.GetFormItemByLabel(ServerInputLabel["webListen"]).(*tview.InputField).GetText()
	webListen = strings.TrimSpace(webListen)
	if !utils.ValidReqField(webListen) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["webListen"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false 
	}

	rtspThreads := form.GetFormItemByLabel(ServerInputLabel["rtspThreads"]).(*tview.InputField).GetText()
	rtspThreads = strings.TrimSpace(rtspThreads)
	if !utils.ValidReqField(rtspThreads) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["rtspThreads"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	}

	rtspListen := form.GetFormItemByLabel(ServerInputLabel["rtspListen"]).(*tview.InputField).GetText()
	rtspListen = strings.TrimSpace(rtspListen)
	if !utils.ValidReqField(rtspListen) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["rtspListen"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	}

	publishThreads := form.GetFormItemByLabel(ServerInputLabel["publishThreads"]).(*tview.InputField).GetText()
	publishThreads = strings.TrimSpace(publishThreads)
	if !utils.ValidReqField(publishThreads) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["publishThreads"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
		return false
	}

	publishListen := form.GetFormItemByLabel(ServerInputLabel["publishListen"]).(*tview.InputField).GetText()
	publishListen = strings.TrimSpace(publishListen)

	if !utils.ValidReqField(publishListen) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["publishListen"])
		g.drawNotifyMsgOkForm(msg, serverFormId)
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

	return true
}


func (g *Gui) drawcontrollerEventsForm(parentForm *tview.Form, cnf *utils.ServerConfig) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle(" Choose controller events ")
	events := map[string]bool{}

	form.AddCheckbox(controllerEventsOptions[0], false, func(checked bool){
		if checked {
			events[controllerEventsOptions[0]] = checked
		} else {
			delete(events, controllerEventsOptions[0])
		}
	})		
	form.AddCheckbox(controllerEventsOptions[1], false, func(checked bool){
		if checked {
			events[controllerEventsOptions[1]] = checked
		} else {
			delete(events, controllerEventsOptions[1])
		}
	})
	form.AddCheckbox(controllerEventsOptions[2], false, func(checked bool){
		if checked {
			events[controllerEventsOptions[2]] = checked
		} else {
			delete(events, controllerEventsOptions[2])
		}
	})
	form.AddCheckbox(controllerEventsOptions[3], false, func(checked bool){
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


func (g *Gui) drawSaveServerNotifyForm(cnf *utils.ServerConfig) {
	modal:= tview.NewModal().
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
