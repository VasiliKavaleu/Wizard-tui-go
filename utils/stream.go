package utils

type StreamCluster struct {
	Node      string
	Secondary []string
}

type Reconnect struct {
	Attemps  int
	Timeout  string
	Interval string
}

type Auth struct {
	Type     string
	User     string
	Password string
}

type Access struct {
	Type string
	Auth
	Whitelist []string
	Limit     int
}

type Dvr struct {
	Enable   bool
	Location string
	Depth    string
	Capacity string
	Chunk    string
}

type Streams struct {
	Uid       string
	Enable    bool
	Stream    string
	Tracks    []string
	Broadcast []string
	Access
	Dvr
}

type StreamConfig struct {
	Camera        string
	Name          string
	Location      string
	Device        string
	Enable        bool
	StreamCluster `yaml:"cluster"`
	Reconnect
	Streams
}
