package gui

import (
	"github.com/rivo/tview"
	"configurator/utils"
	"strings"
)

func (g *Gui) drawApiConfForm(cnf *utils.ServerConfig, parentForm *tview.Form) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle(labelApiForm)

	form.AddInputField(ServerInputLabel["apiListen"], cnf.Api.Listen, inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["apiModule"], cnf.Api.Module, inputWidth, nil, nil)

	form.AddDropDown(ServerInputLabel["apiSsl"], usageDropDownOptions, utils.BoolToIndexDisableAnable(cnf.Api.Ssl), 
					func(option string, optionIndex int){
						if optionIndex != 0 {
							cnf.Api.Ssl = true
						} else {
							cnf.Api.Ssl = false
						}
					})

	form.AddInputField(ServerInputLabel["apiWhitelist"], utils.ListToStr(cnf.Api.Whitelist), inputWidth, nil, nil)
	form.AddDropDown(ServerInputLabel["apiAuth"], apiAuthDropDownOptions, utils.GetIndexFromVal(apiAuthDropDownOptions, cnf.Api.Auth), 
					func(option string, optionIndex int){
						cnf.Api.Auth = option
					})

	form.AddPasswordField(ServerInputLabel["apiUsersAdmin"], cnf.Api.Users.Admin, inputWidth, passwordMask, nil)
	form.AddPasswordField(ServerInputLabel["apiUsersRoot"], cnf.Api.Users.Root, inputWidth, passwordMask, nil)

	saveApiConf := func() {
		if g.validateSaveApiConf(form, cnf) {
			g.pages.RemovePage(apiFormId)
		}
	}

	exit := func() {
		apiUsage := parentForm.GetFormItemByLabel(ServerInputLabel["apiUsage"]).(*tview.DropDown)
		apiUsage.SetCurrentOption(0)
		g.app.SetFocus(apiUsage)
		g.pages.RemovePage(apiFormId)
	}

	form.AddButton("Ok", saveApiConf)
	form.AddButton("Cancel", exit)
	form.SetButtonsAlign(tview.AlignRight)
	form.SetCancelFunc(exit)

	g.pages.AddAndSwitchToPage(apiFormId, g.modal(form, 80, 20), true)

}

func (g *Gui) validateSaveApiConf(form *tview.Form, cnf *utils.ServerConfig) bool {
	apiListen := form.GetFormItemByLabel(ServerInputLabel["apiListen"]).(*tview.InputField).GetText()
	apiListen = strings.TrimSpace(apiListen)
	if !utils.ValidReqField(apiListen) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["apiListen"])
		g.drawNotifyMsgOkForm(msg, apiFormId)
		return false
	}

	apiModule := form.GetFormItemByLabel(ServerInputLabel["apiModule"]).(*tview.InputField).GetText()
	apiModule = strings.TrimSpace(apiModule)
	if !utils.ValidReqField(apiModule) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["apiModule"])
		g.drawNotifyMsgOkForm(msg, apiFormId)
		return false
	}

	apiWhitelist := form.GetFormItemByLabel(ServerInputLabel["apiWhitelist"]).(*tview.InputField).GetText()
	apiWhitelist = strings.TrimSpace(apiWhitelist)
	if !utils.ValidReqField(apiWhitelist) {
		msg := utils.GetReqFieldMsg(ServerInputLabel["apiWhitelist"])
		g.drawNotifyMsgOkForm(msg, apiFormId)
		return false
	}

	if cnf.Api.Auth == "" {
		msg := utils.GetReqFieldMsg(ServerInputLabel["apiAuth"])
		g.drawNotifyMsgOkForm(msg, apiFormId)
		return false
	}  else if cnf.Api.Auth != "none" {
		apiUsersAdmin := form.GetFormItemByLabel(ServerInputLabel["apiUsersAdmin"]).(*tview.InputField).GetText()
		apiUsersAdmin = strings.TrimSpace(apiUsersAdmin)
		if utils.ValidReqField(apiUsersAdmin) {
			cnf.Api.Users.Admin = apiListen 
		} else {
			msg := utils.GetReqFieldMsg(ServerInputLabel["apiUsersAdmin"])
			g.drawNotifyMsgOkForm(msg, apiFormId)
			return false
		}

		apiUsersRoot := form.GetFormItemByLabel(ServerInputLabel["apiUsersRoot"]).(*tview.InputField).GetText()
		apiUsersRoot = strings.TrimSpace(apiUsersRoot)
		if utils.ValidReqField(apiUsersAdmin) {
			cnf.Api.Users.Root = apiUsersRoot 
		} else {
			msg := utils.GetReqFieldMsg(ServerInputLabel["apiUsersRoot"])
			g.drawNotifyMsgOkForm(msg, apiFormId)
			return false
		}
	}

	cnf.Api.Listen = apiListen
	cnf.Api.Module = apiModule
	cnf.Api.Whitelist = utils.StrToList(apiWhitelist) 

	return true
}
