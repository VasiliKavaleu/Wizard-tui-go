package gui

var (
	// messages
	mainMenuTitle      = " Select an option "
	ValidationTitle    = "Check input value"
	quitMsg            = "Exit? Server configuration will be completed."
	saveMsg            = "Save configuration?"
	saveSuccessConfMsg = "Configuration file saved successfully!"
	notSaveMsg         = "Exit? Changes will not be saved!"
	examplePoolFormat  = "1001 : node_host_name-1, 1002: node_host_name-2"

	// form labes
	labelServerForm = " Fill out server configuration form "
	labelApiForm    = " API configuration "
	labelCpanelForm = " Cpanel configuration "
	navigate        = "<Tab>: Down, <Shift+Tab>: Up, <Enter>: Select an option, <Ctrl+S>: Save, <Esc>: Exit"

	menuLabes = map[string]string{
		"createServer": "Create server configuration",
		"createStream": "Create stream configuration",
		"changeServer": "Change server configuration",
		"exit":         "Quit",
	}

	// form id
	serverFormId  = "serverForm"
	apiFormId     = "apiForm"
	cpanelFormId  = "cpanelForm"
	clusterFormId = "clusterForm"
	streamFormId  = "streamForm"

	// initial values
	usageDropDownOptions      = []string{"Disable", "Enable"}
	controllerDsnOptions      = []string{1: "http://", 2: "ws://", 3: "mysql://", 4: "rabbitmq://", 5: "tarantool://"}
	apiAuthDropDownOptions    = []string{"none", "basic", "digest", "token"}
	cpanelAuthDropDownOptions = []string{"none", "basic", "digest"}
	controllerEventsOptions   = []string{"up", "state", "stream", "cluster"}
	accessTypeOptions         = []string{"public", "private", "protected"}
	streamsStreamOptions      = []string{1: "rtsp://", 2: "publish://"}
	streamsTracksOptions      = []string{"video", "audio"}
	streamsBroadcastOptions   = []string{"rtsp", "hls", "mse", "webrtc"}

	inputWidth     = 50
	serverFilePath = "server.yaml"
	passwordMask   = '*'

	// input info
	ServerInputLabel = map[string]string{
		"mediaThreads":  "Number of mediaserver threads",
		"mediaStreams":  "List paths or file streams configuration",
		"mediaStorages": "List mediaserver storages",

		"rtspThreads":    "Number of RTSP threads",
		"rtspListen":     "RTSP socket",
		"publishThreads": "Number of publish threads",
		"publishListen":  "Publish socket",
		"webThreads":     "Number of web threads",
		"webListen":      "Web socket",

		"broadcastSsl":       "Broadcast SSL",
		"broadcastWhitelist": "Broadcast whitelist",

		"tokenSecret":      "Token secret",
		"tokenTtl":         "Token TTL",
		"controllerDsn":    "Сontroller DSN",
		"controllerEvents": "Сontroller events",

		"apiUsage":      "API usage",
		"apiListen":     "Socket",
		"apiModule":     "Module",
		"apiSsl":        "SSL",
		"apiWhitelist":  "Whitelist",
		"apiAuth":       "Authorization",
		"apiUsersAdmin": "Password for admin",
		"apiUsersRoot":  "Password for root",

		"cpanelUsage":      "Cpanel usage",
		"cpanelListen":     "Socket",
		"cpanelModule":     "Module",
		"cpanelSsl":        "SSL",
		"cpanelWhitelist":  "Whitelist",
		"cpanelAuth":       "Authorization",
		"cpanelUsersAdmin": "Password for admin",
		"cpanelUsersRoot":  "Password for root",
		"cpanelUsersUser":  "Password for user",
		"cpanelUsersGuest": "Password for guest",

		"clusterUsage":     "Cluster usage",
		"clusterId":        "Cluster ID",
		"clusterNode":      "Cluster node",
		"clusterPool":      "Cluster pool",
		"clusterWarmingUp": "Warming up time",
		"clusterRetries":   "Number of retries",
		"clusterInterval":  "Interval for reconnect",
	}

	StreamInputLabel = map[string]string{
		"cameraName":     "Camera name",
		"cameraLocation": "Camera location",
		"device":         "Device ID",
		"cameraEnable":   "Camera usage",

		"clusterNode":      "Cluster node",
		"clusterSecondary": "Cluster node secondary",

		"reconnectAttempts": "Number of reconnects",
		"reconnectTimeout":  "Timeout before reconnect attempt",
		"reconnectInterval": "Timeout interval for reconnect",

		"streamsEnable":    "Stream usage",
		"streamsStream":    "RTSP source URL",
		"streamsTracks":    "Enabled rtsp channels",
		"streamsBroadcast": "Enabled restreaming protocols",

		"accessType":         "Access type",
		"accessAuthType":     "Authorization type",
		"accessAuthUser":     "User",
		"accessAuthPassword": "Password",
		"accessWhitelist":    "Access whitelist",
		"accessLimit":        "Access limit of clients",

		"dvrEnable":   "DVR usage",
		"dvrLocation": "Root DVR directory",
		"dvrDepth":    "Duration of DVR keeping",
		"dvrCapacity": "Limit of DVR size",
		"dvrChunk":    "Chunk duration",
	}
)
