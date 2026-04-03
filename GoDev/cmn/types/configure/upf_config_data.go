package configure

//UPF configuraiton
type UpfConfig struct {
	Version        VerConfig
	Logger         LoggerConfig
	N3             IpUpConfig
	N4             N4Config
	N6             IpUpConfig
	Nff            NffConfig
	DnnInfo        []CmDNNInformation `yaml:"dnn info"`
	Pm             Pm
	DnnNameGwIpMap map[string]string // key:dnn name; value:gw ip
}

type N4Config struct {
	Local IpConfig
	Smf   IpConfig
}

type IpConfig struct {
	Ipv4 string
	Port int
}

type IpUpConfig struct {
	PortId int `yaml:"port id"`
	Ipv4   string
	Mask   string
}

type NffConfig struct {
	DpdkArgs           string `yaml:"dpdk args"`
	CpuList            string `yaml:"cpu list"`
	StatsServerNostats bool   `yaml:"stats server nostats"`
	StatsServerAddress string `yaml:"stats server address"`
	StatsServerPort    int    `yaml:"stats server port"`
}
