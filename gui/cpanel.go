package gui

import (
	"github.com/rivo/tview"
	"configurator/utils"
	"strings"
)

func (g *Gui) drawCpanelConfForm(cnf *utils.ServerConfig, parentForm *tview.Form) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle(labelCpanelForm)

	form.AddInputField(ServerInputLabel["cpanelListen"], cnf.Cpanel.Listen, inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["cpanelModule"], cnf.Cpanel.Module, inputWidth, nil, nil)

	form.AddDropDown(ServerInputLabel["cpanelSsl"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Cpanel.Ssl), 
					func(option string, optionIndex int){
						if optionIndex != 0 {
							cnf.Cpanel.Ssl = true
						} else {
							cnf.Cpanel.Ssl = false
						}
					})

	form.AddInputField(ServerInputLabel["cpanelWhitelist"], utils.ListToStr(cnf.Api.Whitelist), inputWidth, nil, nil)
	form.AddDropDown(ServerInputLabel["cpanelAuth"], apiAuthDropDownOptions, utils.GetIndexFromVal(cpanelAuthDropDownOptions, cnf.Cpanel.Auth), 
					func(option string, optionIndex int){
						cnf.Cpanel.Auth = option
					})

	form.AddPasswordField(ServerInputLabel["cpanelUsersAdmin"], cnf.Cpanel.Users.Admin, inputWidth, passwordMask, nil)
	form.AddPasswordField(ServerInputLabel["cpanelUsersRoot"], cnf.Cpanel.Users.Root, inputWidth, passwordMask, nil)
	form.AddPasswordField(ServerInputLabel["cpanelUsersUser"], cnf.Cpanel.Users.User, inputWidth, passwordMask, nil)
	form.AddPasswordField(ServerInputLabel["cpanelUsersGuest"], cnf.Cpanel.Users.Guest, inputWidth, passwordMask, nil)

	saveCpanelConf := func() {
		if g.validateSaveCpanelConf(form, cnf) {
			g.pages.RemovePage(cpanelFormId)
		}
	}
	
	exit := func() {
		cpanelUsage := parentForm.GetFormItemByLabel(ServerInputLabel["cpanelUsage"]).(*tview.DropDown)
		cpanelUsage.SetCurrentOption(0)
		g.app.SetFocus(cpanelUsage)
		g.pages.RemovePage(cpanelFormId)
	}

	form.AddButton("Ok", saveCpanelConf)
	form.AddButton("Cancel", exit)
	form.SetButtonsAlign(tview.AlignRight)
	form.SetCancelFunc(exit)

	g.pages.AddAndSwitchToPage(cpanelFormId, g.modal(form, 80, 23), true)

}

func (g *Gui) validateSaveCpanelConf(form *tview.Form, cnf *utils.ServerConfig) bool {
	cpanelListen := form.GetFormItemByLabel(ServerInputLabel["cpanelListen"]).(*tview.InputField).GetText()
	cpanelListen = strings.TrimSpace(cpanelListen)
	if !utils.ValidReqField(cpanelListen) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["cpanelListen"])
		g.drawNotifyMsgOkForm(msg, apiFormId)
		return false 
	}

	cpanelModule := form.GetFormItemByLabel(ServerInputLabel["cpanelModule"]).(*tview.InputField).GetText()
	cpanelModule = strings.TrimSpace(cpanelModule)
	if utils.ValidReqField(cpanelModule) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["cpanelModule"])
		g.drawNotifyMsgOkForm(msg, apiFormId)
		return false
	}

	cpanelWhitelist := form.GetFormItemByLabel(ServerInputLabel["cpanelWhitelist"]).(*tview.InputField).GetText()
	cpanelWhitelist = strings.TrimSpace(cpanelWhitelist)
	if utils.ValidReqField(cpanelWhitelist) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["cpanelWhitelist"])
		g.drawNotifyMsgOkForm(msg, cpanelFormId)
		return false	 
	}

	if cnf.Cpanel.Auth == "" {
		msg := utils.GetReqFieldMsg(ServerInputLabel["cpanelAuth"])
		g.drawNotifyMsgOkForm(msg, cpanelFormId)
		return false
	}  else if cnf.Cpanel.Auth != "none" {
		cpanelUsersAdmin := form.GetFormItemByLabel(ServerInputLabel["cpanelUsersAdmin"]).(*tview.InputField).GetText()
		cpanelUsersAdmin = strings.TrimSpace(cpanelUsersAdmin)
		if utils.ValidReqField(cpanelUsersAdmin) {
			cnf.Cpanel.Users.Admin = cpanelUsersAdmin 
		} else {
			msg := utils.GetReqFieldMsg(ServerInputLabel["cpanelUsersAdmin"])
			g.drawNotifyMsgOkForm(msg, cpanelFormId)
			return false
		}

		cpanelUsersRoot := form.GetFormItemByLabel(ServerInputLabel["cpanelUsersRoot"]).(*tview.InputField).GetText()
		cpanelUsersRoot = strings.TrimSpace(cpanelUsersRoot)
		if utils.ValidReqField(cpanelUsersRoot) {
			cnf.Cpanel.Users.Root = cpanelUsersRoot 
		} else {
			msg := utils.GetReqFieldMsg(ServerInputLabel["cpanelUsersRoot"])
			g.drawNotifyMsgOkForm(msg, cpanelFormId)
			return false
		}

		cpanelUsersUser := form.GetFormItemByLabel(ServerInputLabel["cpanelUsersUser"]).(*tview.InputField).GetText()
		cpanelUsersUser = strings.TrimSpace(cpanelUsersUser)
		if utils.ValidReqField(cpanelUsersUser) {
			cnf.Cpanel.Users.User = cpanelUsersUser 
		} else {
			msg := utils.GetReqFieldMsg(ServerInputLabel["cpanelUsersUser"])
			g.drawNotifyMsgOkForm(msg, cpanelFormId)
			return false
		}

		cpanelUsersGuest := form.GetFormItemByLabel(ServerInputLabel["cpanelUsersGuest"]).(*tview.InputField).GetText()
		cpanelUsersGuest = strings.TrimSpace(cpanelUsersGuest)
		if utils.ValidReqField(cpanelUsersGuest) {
			cnf.Cpanel.Users.Guest = cpanelUsersGuest 
		} else {
			msg := utils.GetReqFieldMsg(ServerInputLabel["cpanelUsersGuest"])
			g.drawNotifyMsgOkForm(msg, cpanelFormId)
			return false
		}

	}

	cnf.Cpanel.Listen = cpanelListen
	cnf.Cpanel.Module = cpanelModule
	cnf.Cpanel.Whitelist = utils.StrToList(cpanelWhitelist)

	return true
}