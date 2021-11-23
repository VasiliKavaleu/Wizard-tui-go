package main

import (
	"configurator/gui"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	// "io/ioutil"
    "log"
	"os"
	// "strconv"
	// "strings"

    //  "gopkg.in/yaml.v3"
)

// type serverConfig struct {
// 	Server Server `yaml:"server"`
// }

// type Media struct {
// 	Storages  []string
// 	Threads int
// 	Streams string
// }

// type Token struct {
// 	Secret  string
// 	Ttl string
// }

// type Broadcast struct {
// 	Ssl  bool
// 	Whitelist []string
// }

// type Server struct {
// 	Media Media `yaml:"media"`
// 	Token Token `yaml:"token"`
// 	Broadcast Broadcast
// }



// func createCommandList() (commandList *tview.List) {
// 	///// Commands /////
// 	commandList = tview.NewList()                         // создает обертку для команд
// 	commandList.SetBorder(true).SetTitle("Server Master") // с границей и надписью сверху
// 	commandList.ShowSecondaryText(false)                  // показывать ли вторичный текст подсказки
// 	return commandList
// }

// func shooseConf() func() {
// 	return func() {}
// }

// func changeConf() func() {
// 	return func() {}
// }

// func saveConfig(config_data interface{}, config_path string) {
// 	data, err := yaml.Marshal(&config_data)
// 	if err != nil {
// 		log.Fatal(err)
//    }
//    err2 := ioutil.WriteFile(config_path, data, 0644)
//    if err2 != nil {
// 		log.Fatal(err2)
//    }
// }

// func strToList(value string) (values []string) {
// 	values = strings.Split(value, ",")
// 	return
// }

// func listToStr(values []string) (value string) {
// 	value = strings.Join(values, ",")
// 	return
// }

// func validInt(value string, lastChar rune) bool {
// 	_, err := strconv.Atoi(value)
// 	return err == nil
// }

// func apiForm() {
// 	form := tview.NewForm()
// 	form.SetBorder(true)
// 	form.SetTitleAlign(tview.AlignLeft)
// 	form.SetTitle("API config")
// 	form.AddInputField("Repository", "", 50, nil, nil)
// }

// func enableApi(checked bool) {
// 	if checked {
// 		apiForm()
// 	}
// }


// func createConf(pages *tview.Pages, conf *serverConfig) func() {  // функция для вывода страницы создания конфигурации файла
// 	return func() {
// 		// serverTitle := tview.NewTextView().SetText("Server")  // TextView is a box which displays text.
// 		form := tview.NewForm()
// 		form.AddInputField("Media storages:", listToStr(conf.Server.Media.Storages), 50, nil, func(value string) { conf.Server.Media.Storages = strToList(value) })
// 		form.AddInputField("Media threads:", strconv.Itoa(conf.Server.Media.Threads), 50, validInt, func(value string) { conf.Server.Media.Threads, _ = strconv.Atoi(value) })
// 		form.AddInputField("Media streams:", conf.Server.Media.Streams, 50, nil, func(value string) { conf.Server.Media.Streams = value })

// 		form.AddCheckbox("Enable/disable broadcast SSL:", conf.Server.Broadcast.Ssl, func(checked bool) { conf.Server.Broadcast.Ssl = checked })
// 		form.AddInputField("Broadcast whitelist:", listToStr(conf.Server.Broadcast.Whitelist), 50, nil, func(value string) { conf.Server.Broadcast.Whitelist = strToList(value) })

// 		form.AddInputField("Token secret:", conf.Server.Token.Secret, 50, nil, func(value string) { conf.Server.Token.Secret = value })
// 		form.AddInputField("Token ttl:", conf.Server.Token.Ttl, 50, nil, func(value string) { conf.Server.Token.Ttl = value })

// 		form.AddCheckbox("API usage enable/disable:", false, enableApi)

// 		cancelFunc := func() {                                              // handler func for 'cancel' button 
// 			pages.SwitchToPage("main")									    // sets a page's visibility to "true" and all other pages' visibility to "false". 
// 			pages.RemovePage("modal")                                       // removes the page with the given name. If that page was the only visible page, visibility is assigned to the last page
// 		}

// 		saveServerConf := func() {
// 			saveConfig(conf, "server.yaml")
// 			cancelFunc()
// 		}

// 		form.AddButton("Save", saveServerConf)                                         // add button
// 		form.AddButton("Cancel", cancelFunc)                                // add button
// 		form.SetCancelFunc(cancelFunc)                                      // sets a handler which is called when the user hits the Escape key
// 		form.SetButtonsAlign(tview.AlignCenter) 							// aligning form button
// 		form.SetBorder(true).SetTitle("Fill out the server config form")    // set border and title



// 		// modal := createModalForm(pages, form, 13, 55)
// 		createSereverConfigPage := func(form tview.Primitive) tview.Primitive {
// 			createPage := tview.NewFlex().SetDirection(tview.FlexRow)
// 			// createPage.AddItem(serverTitle, 0, 1, false)
// 			createPage.AddItem(form, 0, 2, true)
// 			return createPage
// 		}(form)

// 		pages.AddPage("modal", createSereverConfigPage, true, true)
// 	}
// }

// func createMainLayout(commandList tview.Primitive) (layout *tview.Frame) {
// 	///// Main Layout /////
// 	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow). // NewFlex returns a new flexbox layout container with no primitives
// 									AddItem(commandList, 10, 1, true) // adds a new item to the container.
// 		// .AddItem(outputPanel, 0, 4, false)

// 	info := tview.NewTextView() // returns a new text view.
// 	info.SetBorder(true)
// 	info.SetText("UDP Packet Tester v1.0 - Copyright 2019 Andrew Young <andrew@vaelen.org>") // устанавливает текст
// 	info.SetTextAlign(tview.AlignCenter)                                                     // устанавливает расположение текста

	
// 	inner_layout := tview.NewFlex().SetDirection(tview.FlexRow).
// 		AddItem(mainLayout, 0, 1, true). // 3е значение proportion - пропорция между AddItem в рамках текущего layout, только при fixedSize
// 		AddItem(info, 3, 1, false)

// 	clayout := tview.NewFrame(inner_layout).SetBorders(2, 2, 2, 2, 4, 4)

// 	return clayout
// }

// func createApplication() (app *tview.Application) {

// 	app = tview.NewApplication() // создает новое приложение
// 	pages := tview.NewPages()    // создает pop для переключения

// 	// infoUI := createInfoPanel(app)
// 	// logPanel := createTextViewPanel(app, "Log")
// 	// outputPanel := createOutputPanel(app, infoUI.panel, logPanel)

// 	// log.SetOutput(logPanel)

// 	conf := &serverConfig{}

// 	commandList := createCommandList()
// 	commandList.AddItem("Choose existing configuration", "", '1', shooseConf()) // listenCommand(pages, infoUI, conf)
// 	commandList.AddItem("Change server configuration", "", '2', changeConf())
// 	commandList.AddItem("Create server configuration", "", '3', createConf(pages, conf))
// 	commandList.AddItem("Create strea configuration", "", '4', changeConf())
// 	commandList.AddItem("Quit", "", 'q', func() {
// 		app.Stop()
// 	})

// 	layout := createMainLayout(commandList)   // главный шаблон страницы
// 	pages.AddPage("main", layout, true, true) // добавляет шаблон для прорисовки страницы с именем main, layout - сам шаблон, with resize, visible bool

// 	app.SetRoot(pages, true) // вызывается один раз или никогда, trure - заполняет экран

// 	return app
// }

// func main() {

// 	app := createApplication()

// 	if err := app.Run(); err != nil {
// 		panic(err)
// 	}

// }

func run() int {
	tview.Styles.ContrastBackgroundColor = tcell.ColorSlateGrey
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorTan
	tview.Styles.BorderColor = tcell.ColorTan
	tview.Styles.TitleColor = tcell.ColorLightGoldenrodYellow
	gui := gui.New()

	if err := gui.Start(); err != nil {
		log.Fatal("Cannot start Wizard: %s", err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run())
}
