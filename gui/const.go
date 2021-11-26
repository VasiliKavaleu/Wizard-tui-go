package gui

var (
	// messages
	mainMenuTitle = " Select an option "
	ValidationTitle = "Check input value"
	quitMsg = "Exit? Server configuration will be completed."
	saveMsg = "Save configuration?"
	saveSuccessConfMsg = "Configuration file saved successfully!"
	notSaveMsg = "Exit? Changes will not be saved!"
	examplePoolFormat = "1001 : node_host_name-1, 1002: node_host_name-2"

	// form labes
	labelServerForm = " Fill out server configuration form "
	labelApiForm = " API configuration "
	labelCpanelForm = " Cpanel configuration "

	menuLabes = map[string]string{
		"createServer": "Create server configuration",
		"changeServer": "Change server configuration",
		"exit": "Quit",
	}

	// form id
	serverFormId = "serverForm"
	apiFormId = "apiForm"
	cpanelFormId = "cpanelForm"
	clusterFormId = "clusterForm"

	// initial values
	usageDropDownOptions = []string{"Disable", "Enable"}
	controllerDsnOptions = []string{1: "http://", 2: "ws://", 3: "mysql://", 4: "rabbitmq://", 5: "tarantool://"}
	apiAuthDropDownOptions = []string{"none", "basic", "digest", "token"}
	cpanelAuthDropDownOptions = []string{"none", "basic", "digest"}
	controllerEventsOptions = []string{"up", "state", "stream", "cluster"}

	inputWidth = 50
	serverFilePath = "server.yaml"
	passwordMask = '*'

	// input info
	ServerInputLabel = map[string]string{
		"mediaThreads": "Number of mediaserver threads",
		"mediaStreams": "List paths or file streams configuration",
		"mediaStorages": "List mediaserver storages",
	
		"rtspThreads": "Number of RTSP threads",
		"rtspListen": "RTSP socket",
		"publishThreads": "Number of publish threads",
		"publishListen": "Publish socket",
		"webThreads": "Number of web threads",
		"webListen": "Web socket",
	
		"broadcastSsl": "Broadcast SSL",
		"broadcastWhitelist": "Broadcast whitelist",
	
		"tokenSecret": "Token secret",
		"tokenTtl": "Token TTL",
		"controllerDsn": "Сontroller DSN",
		"controllerEvents": "Сontroller events",
		
		"apiUsage": "API usage",
		"apiListen": "Socket",
		"apiModule": "Module",
		"apiSsl": "SSL",
		"apiWhitelist": "Whitelist",
		"apiAuth": "Authorization",
		"apiUsersAdmin": "Password for admin",
		"apiUsersRoot": "Password for root",
	
		"cpanelUsage": "Cpanel usage",
		"cpanelListen": "Socket",
		"cpanelModule": "Module",
		"cpanelSsl": "SSL",
		"cpanelWhitelist": "Whitelist",
		"cpanelAuth": "Authorization",
		"cpanelUsersAdmin": "Password for admin",
		"cpanelUsersRoot": "Password for root",
		"cpanelUsersUser": "Password for user",
		"cpanelUsersGuest": "Password for guest",
	
		"clusterUsage": "Cluster usage",
		"clusterId": "Cluster ID",
		"clusterNode": "Cluster node",
		"clusterPool": "Cluster pool",
		"clusterWarmingUp": "Warming up time",
		"clusterRetries": "Number of retries",
		"clusterInterval": "Interval for reconnect",
	}

)
