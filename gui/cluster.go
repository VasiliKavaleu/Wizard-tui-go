package gui

import (
	"github.com/rivo/tview"
	"configurator/utils"
	"strconv"
	"strings"
)

func (g *Gui) drawClusterConfForm(cnf *utils.ServerConfig, parentForm *tview.Form) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle(" Cluster configuration ")

	form.AddInputField(ServerInputLabel["clusterId"], strconv.Itoa(cnf.Cluster.Id), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["clusterNode"], strconv.Itoa(cnf.Cluster.Node), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["clusterPool"], cnf.Cluster.Pool, inputWidth, nil, nil)

	form.AddInputField(ServerInputLabel["clusterWarmingUp"], cnf.Cluster.Healthcheck.WarmingUp, inputWidth, nil, nil)
	form.AddInputField(ServerInputLabel["clusterRetries"], strconv.Itoa(cnf.Cluster.Healthcheck.Retries), inputWidth, utils.ValidInt, nil)
	form.AddInputField(ServerInputLabel["clusterInterval"], cnf.Cluster.Healthcheck.Interval, inputWidth, nil, nil)

	saveClusterConf := func() {
		if g.validateSaveClusterConf(form, cnf) {
			g.pages.RemovePage(clusterFormId)
		}
	}

	exit := func() {
		clusterlUsage := parentForm.GetFormItemByLabel(ServerInputLabel["clusterUsage"]).(*tview.DropDown)
		clusterlUsage.SetCurrentOption(0)
		g.app.SetFocus(clusterlUsage)
		g.pages.RemovePage(clusterFormId)
	}

	form.AddButton("Ok", saveClusterConf)
	form.AddButton("Cancel", exit)
	form.SetButtonsAlign(tview.AlignRight)
	form.SetCancelFunc(exit)

	g.pages.AddAndSwitchToPage(clusterFormId, g.modal(form, 80, 20), true)
}


func (g *Gui) validateSaveClusterConf(form *tview.Form, cnf *utils.ServerConfig) bool {
	clusterId := form.GetFormItemByLabel(ServerInputLabel["clusterId"]).(*tview.InputField).GetText()
	clusterNode := form.GetFormItemByLabel(ServerInputLabel["clusterNode"]).(*tview.InputField).GetText()
	clusterPool := form.GetFormItemByLabel(ServerInputLabel["clusterPool"]).(*tview.InputField).GetText()
	clusterWarmingUp := form.GetFormItemByLabel(ServerInputLabel["clusterWarmingUp"]).(*tview.InputField).GetText()
	clusterRetries := form.GetFormItemByLabel(ServerInputLabel["clusterRetries"]).(*tview.InputField).GetText()
	clusterInterval := form.GetFormItemByLabel(ServerInputLabel["clusterInterval"]).(*tview.InputField).GetText()

	clusterId = strings.TrimSpace(clusterId)
	clusterNode = strings.TrimSpace(clusterNode)
	clusterPool = strings.TrimSpace(clusterPool)
	clusterWarmingUp = strings.TrimSpace(clusterWarmingUp)
	clusterRetries = strings.TrimSpace(clusterRetries)
	clusterInterval = strings.TrimSpace(clusterInterval)

	if g.checkAndNotifyRequiredField(clusterId, ServerInputLabel["clusterId"], clusterFormId) &&
	g.checkAndNotifyRequiredField(clusterNode, ServerInputLabel["clusterNode"], clusterFormId) &&
	g.checkAndNotifyRequiredField(clusterPool, ServerInputLabel["clusterPool"], clusterFormId) &&
	g.checkAndNotifyRequiredField(clusterWarmingUp, ServerInputLabel["clusterWarmingUp"], clusterFormId) &&
	g.checkAndNotifyRequiredField(clusterRetries, ServerInputLabel["clusterRetries"], clusterFormId) &&
	g.checkAndNotifyRequiredField(clusterInterval, ServerInputLabel["clusterInterval"], clusterFormId){
		cnf.Cluster.Id, _ = strconv.Atoi(clusterId)
		cnf.Cluster.Node, _ = strconv.Atoi(clusterNode)
		cnf.Cluster.Pool = clusterPool
		cnf.Cluster.WarmingUp = clusterWarmingUp
		cnf.Cluster.Retries, _ = strconv.Atoi(clusterRetries)
		cnf.Cluster.Interval = clusterInterval
		return true
	}

	return false
}


