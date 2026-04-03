package configure

type CmSysConfig struct {
	RedisSerAddr IpAddr `yaml:"redis server"`
	BDConfig     DBConf `yaml:"db server"`
}

type DBConf struct {
	UserName   string
	Pwd        string
	Ip         string
	Port       int
	DriverName string
	DbName     string
	Debug      bool
}

