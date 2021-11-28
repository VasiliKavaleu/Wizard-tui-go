package gui

import (
	"configurator/utils"
	"github.com/rivo/tview"
	"github.com/satori/go.uuid"
	// "gopkg.in/yaml.v3"
	// "io/ioutil"
	"strconv"
	"strings"
)

func (g *Gui) drawStreamConfForm() {
	cnf := &utils.StreamConfig{}

	form := tview.NewForm()

	var uid uuid.UUID

	uid = uuid.NewV4()
	cnf.Camera = uid.String()

	uid = uuid.NewV4()
	cnf.Streams.Uid = uid.String()

	form.AddInputField(StreamInputLabel["cameraName"], "", inputWidth, nil, func(value string) {
		cnf.Name = strings.TrimSpace(value)
	})
	form.AddInputField(StreamInputLabel["cameraLocation"], "", inputWidth, nil, func(value string) {
		cnf.Location = strings.TrimSpace(value)
	})
	form.AddInputField(StreamInputLabel["device"], "", inputWidth, nil, func(value string) {
		cnf.Device = strings.TrimSpace(value)
	})

	form.AddDropDown(StreamInputLabel["cameraEnable"], usageDropDownOptions, 1,
	func(option string, optionIndex int) {
		if optionIndex != 0 {
			cnf.Enable = true
		} else {
			cnf.Enable = false
		}
	})

	form.AddInputField(StreamInputLabel["clusterNode"], "", inputWidth, nil, nil)

	form.AddInputField(StreamInputLabel["clusterSecondary"], "", inputWidth, nil, func(value string) {		
		value = strings.TrimSpace(value)
		cnf.Secondary = utils.StrToList(value)
	})

	form.AddInputField(StreamInputLabel["reconnectAttempts"], "", inputWidth, utils.ValidInt, nil)
	form.AddInputField(StreamInputLabel["reconnectTimeout"], "", inputWidth, nil, nil)
	form.AddInputField(StreamInputLabel["reconnectInterval"], "", inputWidth, nil, nil)


	form.AddDropDown(StreamInputLabel["streamsEnable"], usageDropDownOptions, 1,
	func(option string, optionIndex int) {
		if optionIndex != 0 {
			cnf.Streams.Enable = true
		} else {
			cnf.Streams.Enable = false
		}
	})

	form.AddDropDown(StreamInputLabel["streamsStream"], streamsStreamOptions, 0, 
	func(option string, optionIndex int) {
		cnf.Streams.Stream = option
	})

	form.AddInputField(StreamInputLabel["streamsTracks"], "", inputWidth, nil, func(text string) {
		g.drawStreamsTracksForm(form, cnf)
	})

	form.AddInputField(StreamInputLabel["streamsBroadcast"], "", inputWidth, nil, func(text string) {
		g.drawStreamsBroadcastForm(form, cnf)
	})

	form.AddDropDown(StreamInputLabel["accessType"], accessTypeOptions, 0, 
	func(option string, optionIndex int) {
		cnf.Streams.Type = option
	})

	form.AddInputField(StreamInputLabel["accessAuthType"], "", inputWidth, nil, nil)
	form.AddInputField(StreamInputLabel["accessAuthUser"], "", inputWidth, nil, nil)
	form.AddInputField(StreamInputLabel["accessAuthPassword"], "", inputWidth, nil, nil)

	form.AddInputField(StreamInputLabel["accessWhitelist"], "", inputWidth, nil, func(value string) {
		value = strings.TrimSpace(value)
		cnf.Access.Whitelist = utils.StrToList(value)
	})
	form.AddInputField(StreamInputLabel["accessLimit"], "", inputWidth, utils.ValidInt, func(value string) {
		value = strings.TrimSpace(value)
		cnf.Access.Limit, _ = strconv.Atoi(value)
	})

	form.AddDropDown(StreamInputLabel["dvrEnable"], usageDropDownOptions, 1,
	func(option string, optionIndex int) {
		if optionIndex != 0 {
			cnf.Streams.Enable = true
		} else {
			cnf.Streams.Enable = false
		}
	})

	form.AddInputField(StreamInputLabel["dvrLocation"], "", inputWidth, nil, nil)
	form.AddInputField(StreamInputLabel["dvrDepth"], "", inputWidth, nil, nil)
	form.AddInputField(StreamInputLabel["dvrCapacity"], "", inputWidth, nil, nil)
	form.AddInputField(StreamInputLabel["dvrChunk"], "", inputWidth, nil, nil)

	exit := func() {
		g.drawOkCancelNotifyForm(notSaveMsg, "Ok", "Cancel", serverFormId, "main")
	}

	saveStreamConf := func() {
		if g.validateSaveStreamConf(form, cnf) {
			g.drawSaveStreamNotifyForm(cnf)
		}
	}

	form.AddButton("Save", saveStreamConf)
	form.AddButton("Cancel", exit)
	form.SetCancelFunc(exit)
	form.SetButtonsAlign(tview.AlignRight)
	form.SetBorder(true).SetTitle(labelServerForm)

	createStreamConfigPage := createFormLayout(form)
	g.pages.AddAndSwitchToPage(streamFormId, createStreamConfigPage, true)
}


func (g *Gui) validateSaveStreamConf(form *tview.Form, cnf *utils.StreamConfig) bool {
	clusterNode := form.GetFormItemByLabel(StreamInputLabel["clusterNode"]).(*tview.InputField).GetText()
	reconnectAttempts := form.GetFormItemByLabel(StreamInputLabel["reconnectAttempts"]).(*tview.InputField).GetText()
	reconnectTimeout := form.GetFormItemByLabel(StreamInputLabel["reconnectTimeout"]).(*tview.InputField).GetText()
	reconnectInterval := form.GetFormItemByLabel(StreamInputLabel["reconnectInterval"]).(*tview.InputField).GetText()
	accessAuthType := form.GetFormItemByLabel(StreamInputLabel["accessAuthType"]).(*tview.InputField).GetText()
	accessAuthUser := form.GetFormItemByLabel(StreamInputLabel["accessAuthUser"]).(*tview.InputField).GetText()
	accessAuthPassword := form.GetFormItemByLabel(StreamInputLabel["accessAuthPassword"]).(*tview.InputField).GetText()

	dvrLocation := form.GetFormItemByLabel(StreamInputLabel["dvrLocation"]).(*tview.InputField).GetText()
	dvrDepth := form.GetFormItemByLabel(StreamInputLabel["dvrDepth"]).(*tview.InputField).GetText()
	dvrCapacity := form.GetFormItemByLabel(StreamInputLabel["dvrCapacity"]).(*tview.InputField).GetText()
	dvrChunk := form.GetFormItemByLabel(StreamInputLabel["dvrChunk"]).(*tview.InputField).GetText()

	clusterNode = strings.TrimSpace(clusterNode)
	reconnectAttempts = strings.TrimSpace(reconnectAttempts)
	reconnectTimeout = strings.TrimSpace(reconnectTimeout)
	reconnectInterval = strings.TrimSpace(reconnectInterval)
	accessAuthType = strings.TrimSpace(accessAuthType)
	accessAuthUser = strings.TrimSpace(accessAuthUser)
	accessAuthPassword = strings.TrimSpace(accessAuthPassword)
	dvrLocation = strings.TrimSpace(dvrLocation)
	dvrDepth = strings.TrimSpace(dvrDepth)
	dvrCapacity = strings.TrimSpace(dvrCapacity)
	dvrChunk = strings.TrimSpace(dvrChunk)

	if g.checkAndNotifyRequiredField(clusterNode, StreamInputLabel["clusterNode"], streamFormId) &&

	g.checkAndNotifyRequiredField(reconnectAttempts, StreamInputLabel["reconnectAttempts"], streamFormId) &&
	g.checkAndNotifyRequiredField(reconnectTimeout, StreamInputLabel["reconnectTimeout"], streamFormId) &&
	g.checkAndNotifyRequiredField(reconnectInterval, StreamInputLabel["reconnectInterval"], streamFormId) &&

	g.checkAndNotifyRequiredField(accessAuthType, StreamInputLabel["accessAuthType"], streamFormId) &&
	g.checkAndNotifyRequiredField(accessAuthUser, StreamInputLabel["accessAuthUser"], streamFormId) &&
	g.checkAndNotifyRequiredField(accessAuthPassword, StreamInputLabel["accessAuthPassword"], streamFormId) {

		cnf.Node = clusterNode
		cnf.Reconnect.Attemps, _ = strconv.Atoi(reconnectAttempts)
		cnf.Reconnect.Timeout = reconnectTimeout
		cnf.Reconnect.Interval = reconnectInterval
		cnf.Access.Auth.Type = accessAuthType
		cnf.Access.Auth.User = accessAuthUser
		cnf.Access.Auth.Password = accessAuthPassword

	} else {
		return false
	}

	if cnf.Streams.Enable {
		if 	g.checkAndNotifyRequiredField(dvrLocation, StreamInputLabel["dvrLocation"], streamFormId) &&
		g.checkAndNotifyRequiredField(dvrDepth, StreamInputLabel["dvrDepth"], streamFormId) &&
		g.checkAndNotifyRequiredField(dvrCapacity, StreamInputLabel["dvrCapacity"], streamFormId) &&
		g.checkAndNotifyRequiredField(dvrChunk, StreamInputLabel["dvrChunk"], streamFormId) {
			cnf.Streams.Location = dvrLocation
			cnf.Streams.Depth = dvrDepth
			cnf.Streams.Capacity = dvrCapacity
			cnf.Streams.Chunk = dvrChunk
		} else {
			return false
		}
	}

	if cnf.Streams.Stream == "" {
		msg := utils.GetReqFieldMsg(ServerInputLabel["streamsStream"])
		g.drawNotifyMsgOkForm(msg, streamFormId)
		return false
	}
	return true

}

func (g *Gui) drawSaveStreamNotifyForm(cnf *utils.StreamConfig) {
	modal := tview.NewModal().
	SetText(saveMsg).
	AddButtons([]string{"Save", "Cancel"}).
	SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Cancel" {
			g.pages.RemovePage("saveStreamNotify")
		} else {
			g.pages.RemovePage("saveStreamNotify")
			utils.SaveConfigToFile(cnf, utils.GetFileNameByUuid(cnf.Streams.Uid))
			g.drawNotifyMsgOkForm(saveSuccessConfMsg, "main")
		}
	})
g.pages.AddAndSwitchToPage("saveStreamNotify", modal, true).ShowPage(streamFormId)
}

// form reflecting checkboxes
func (g *Gui) drawStreamsTracksForm(parentForm *tview.Form, cnf *utils.StreamConfig) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle(" Choose RTSP channels ")
	events := map[string]bool{}

	form.AddCheckbox(streamsTracksOptions[0], false, func(checked bool) {
		if checked {
			events[streamsTracksOptions[0]] = checked
		} else {
			delete(events, streamsTracksOptions[0])
		}
	})
	form.AddCheckbox(streamsTracksOptions[1], false, func(checked bool) {
		if checked {
			events[streamsTracksOptions[1]] = checked
		} else {
			delete(events, streamsTracksOptions[1])
		}
	})

	exit := func() {
		eventsList := []string{}
		for key := range events {
			eventsList = append(eventsList, key)
		}
		cnf.Tracks = eventsList

		streamsTracks := parentForm.GetFormItemByLabel(StreamInputLabel["streamsTracks"]).(*tview.InputField)
		streamsTracks.SetText(utils.ListToStr(eventsList))
		g.app.SetFocus(streamsTracks)

		g.pages.RemovePage("streamsTracks")
	}

	form.AddButton("Ok", exit)
	form.SetButtonsAlign(tview.AlignRight)
	form.SetCancelFunc(exit)

	g.pages.AddAndSwitchToPage("streamsTracks", g.modal(form, 40, 10), true).ShowPage(streamFormId)
}

// form reflecting checkboxes
func (g *Gui) drawStreamsBroadcastForm(parentForm *tview.Form, cnf *utils.StreamConfig) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle(" Choose restreaming protocols ")
	events := map[string]bool{}

	form.AddCheckbox(streamsBroadcastOptions[0], false, func(checked bool) {
		if checked {
			events[streamsBroadcastOptions[0]] = checked
		} else {
			delete(events, streamsBroadcastOptions[0])
		}
	})
	form.AddCheckbox(streamsBroadcastOptions[1], false, func(checked bool) {
		if checked {
			events[streamsBroadcastOptions[1]] = checked
		} else {
			delete(events, streamsBroadcastOptions[1])
		}
	})
	form.AddCheckbox(streamsBroadcastOptions[2], false, func(checked bool) {
		if checked {
			events[streamsBroadcastOptions[2]] = checked
		} else {
			delete(events, streamsBroadcastOptions[2])
		}
	})
	form.AddCheckbox(streamsBroadcastOptions[3], false, func(checked bool) {
		if checked {
			events[streamsBroadcastOptions[3]] = checked
		} else {
			delete(events, streamsBroadcastOptions[3])
		}
	})

	exit := func() {
		eventsList := []string{}
		for key := range events {
			eventsList = append(eventsList, key)
		}
		cnf.Broadcast = eventsList

		streamsTracks := parentForm.GetFormItemByLabel(StreamInputLabel["streamsBroadcast"]).(*tview.InputField)
		streamsTracks.SetText(utils.ListToStr(eventsList))
		g.app.SetFocus(streamsTracks)

		g.pages.RemovePage("streamsBroadcast")
	}

	form.AddButton("Ok", exit)
	form.SetButtonsAlign(tview.AlignRight)
	form.SetCancelFunc(exit)

	g.pages.AddAndSwitchToPage("streamsBroadcast", g.modal(form, 40, 13), true).ShowPage(streamFormId)
}