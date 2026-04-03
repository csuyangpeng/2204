package configure

type SBI struct {
	Amf SBIConfig
	Udm SBIConfig
	Smf SBIConfig
}

type SBIConfig struct {
	Addr   IpAddr
	Scheme string
}

type IpAddr struct {
	Ip string
	Port int
}

type LoggerConfig struct {
	Level   string
	Path    string
	Control int8
}

type VerConfig struct {
	Main  string
	Patch string
}
