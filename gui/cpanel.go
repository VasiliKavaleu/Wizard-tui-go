package gui

import (
	"configurator/utils"
	"github.com/rivo/tview"
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
		func(option string, optionIndex int) {
			if optionIndex != 0 {
				cnf.Cpanel.Ssl = true
			} else {
				cnf.Cpanel.Ssl = false
			}
		})

	form.AddInputField(ServerInputLabel["cpanelWhitelist"], utils.ListToStr(cnf.Cpanel.Whitelist), inputWidth, nil, nil)
	form.AddDropDown(ServerInputLabel["cpanelAuth"], cpanelAuthDropDownOptions, utils.GetIndexFromVal(cpanelAuthDropDownOptions, cnf.Cpanel.Auth),
		func(option string, optionIndex int) {
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
	cpanelModule := form.GetFormItemByLabel(ServerInputLabel["cpanelModule"]).(*tview.InputField).GetText()
	cpanelWhitelist := form.GetFormItemByLabel(ServerInputLabel["cpanelWhitelist"]).(*tview.InputField).GetText()

	cpanelListen = strings.TrimSpace(cpanelListen)
	cpanelModule = strings.TrimSpace(cpanelModule)
	cpanelWhitelist = strings.TrimSpace(cpanelWhitelist)

	if g.checkAndNotifyRequiredField(cpanelListen, ServerInputLabel["cpanelListen"], cpanelFormId) &&
		g.checkAndNotifyRequiredField(cpanelModule, ServerInputLabel["cpanelModule"], cpanelFormId) &&
		g.checkAndNotifyRequiredField(cpanelWhitelist, ServerInputLabel["cpanelWhitelist"], cpanelFormId) {
		cnf.Cpanel.Listen = cpanelListen
		cnf.Cpanel.Module = cpanelModule
		cnf.Cpanel.Whitelist = utils.StrToList(cpanelWhitelist)
	} else {
		return false
	}

	if cnf.Cpanel.Auth != "none" {
		cpanelUsersAdmin := form.GetFormItemByLabel(ServerInputLabel["cpanelUsersAdmin"]).(*tview.InputField).GetText()
		cpanelUsersRoot := form.GetFormItemByLabel(ServerInputLabel["cpanelUsersRoot"]).(*tview.InputField).GetText()
		cpanelUsersUser := form.GetFormItemByLabel(ServerInputLabel["cpanelUsersUser"]).(*tview.InputField).GetText()
		cpanelUsersGuest := form.GetFormItemByLabel(ServerInputLabel["cpanelUsersGuest"]).(*tview.InputField).GetText()

		cpanelUsersAdmin = strings.TrimSpace(cpanelUsersAdmin)
		cpanelUsersRoot = strings.TrimSpace(cpanelUsersRoot)
		cpanelUsersUser = strings.TrimSpace(cpanelUsersUser)
		cpanelUsersGuest = strings.TrimSpace(cpanelUsersGuest)

		if g.checkAndNotifyRequiredField(cpanelUsersAdmin, ServerInputLabel["cpanelUsersAdmin"], cpanelFormId) &&
			g.checkAndNotifyRequiredField(cpanelUsersRoot, ServerInputLabel["cpanelUsersRoot"], cpanelFormId) &&
			g.checkAndNotifyRequiredField(cpanelUsersUser, ServerInputLabel["cpanelUsersUser"], cpanelFormId) &&
			g.checkAndNotifyRequiredField(cpanelUsersGuest, ServerInputLabel["cpanelUsersGuest"], cpanelFormId) {
			cnf.Cpanel.Users.Admin = cpanelUsersAdmin
			cnf.Cpanel.Users.Root = cpanelUsersRoot
			cnf.Cpanel.Users.User = cpanelUsersUser
			cnf.Cpanel.Users.Guest = cpanelUsersGuest
		} else {
			return false
		}
	}
	return true
}
