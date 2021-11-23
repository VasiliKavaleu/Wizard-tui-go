package gui

import (
	"github.com/rivo/tview"
	"configurator/utils"
	"strconv"
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


	exit := func() {
		clusterlUsage := parentForm.GetFormItemByLabel(ServerInputLabel["clusterUsage"]).(*tview.DropDown)
		clusterlUsage.SetCurrentOption(0)
		g.app.SetFocus(clusterlUsage)
		g.pages.RemovePage("clusterForm").ShowPage("main")
	}

	form.AddButton("Ok", exit)
	form.AddButton("Cancel", exit)
	form.SetButtonsAlign(tview.AlignRight)
	form.SetCancelFunc(exit)

	g.pages.AddAndSwitchToPage("clusterForm", g.modal(form, 80, 20), true)
}