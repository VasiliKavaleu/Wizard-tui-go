package utils


type ServerConfig struct {
	Server Server `yaml:"server"`
	Controller Controller
	Api Api
	Cpanel Cpanel
	Cluster Cluster
}

type Media struct {
	Storages  []string
	Threads int
	Streams []string
}

type Token struct {
	Secret  string
	Ttl int
}

type Web struct {
	Listen  string
	Threads int
}

type Publish struct {
	Listen  string
	Threads int
}

type Rtsp struct {
	Listen  string
	Threads int
}

type Broadcast struct {
	Ssl  bool
	Whitelist []string
	Web Web
	Publish Publish
	Rtsp Rtsp
}

type ApiUsers struct {
	Admin string
	Root string
}

type CpanelUsers struct {
	Admin string
	Root string
	User string
	Guest string
}

type Api struct {
	Enable bool
	Listen string
	Module string
	Ssl bool
	Whitelist []string
	Auth  string
	Users ApiUsers
}

type Cpanel struct {
	Enable bool
	Listen string
	Module string
	Ssl bool
	Whitelist []string
	Auth  string
	Users CpanelUsers
}

type Healthcheck struct {
	WarmingUp string `yaml:"warming-up"`
	Retries int `yaml:"retries"`
	Interval string `yaml:"interval"`
}

type Cluster struct {
	Enable bool
	Id int
	Node int
	Pool []map[int]string
	Healthcheck  `yaml:"healthcheck"`
}

type Controller struct {
	Dsn string
	Events []string
}

type Server struct {
	Media Media `yaml:"media"`
	Token Token `yaml:"token"`
	Broadcast Broadcast
}
