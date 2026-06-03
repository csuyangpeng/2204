package configure

type CmAmfConfig struct {
	Version  VerConfig
	Logger   LoggerConfig
	N2if     IpAddr
	Sctp     SctpParameter
	Service  CmAmfService
	Nas      CmNasConfig
	Plmnlist []string      `yaml:"plmn list"`
	TaiLists []CmnTaiList  `yaml:"tai lists"`
	Nssai    []CmSNssai
	Sbi      SBI
}

type CmAmfService struct {
	AmfName             string          `yaml:"amf name"`
	RelativeAmfCapacity int             `yaml:"relative amf capacity"`
	AmfIdentifier       CmAmfIdentifier `yaml:"amf identifier"`
}

type CmAmfIdentifier struct {
	RegionId uint8  `yaml:"region id"` //8  bit Length
	SetId    uint16 `yaml:"set id"`    //10 bit Length
	Pointer  uint8  //6  bit Length
}

type CmNasConfig struct {
	EncryptAlgorithm   string `yaml:"encrypt algorithm"`
	IntegrityAlgorithm string `yaml:"integrity algorithm"`
	T3512min           int    `yaml:"t3512 min"`
	T3513Sec           int    `yaml:"t3513 sec"`
	T3502min           int    `yaml:"t3502 min"`
	T3550sec           int    `yaml:"t3550 sec"` //Retransmission of REGISTRATION ACCEPT message
	T3560sec           int    `yaml:"t3560 sec"` //Retransmission of AUTHENTICATION REQUEST message or SECURITY MODE COMMAND message
	T3570sec           int    `yaml:"t3570 sec"` //Retransmission of IDENTITY REQUEST message
	T3522sec           int    `yaml:"t3522 sec"` //Retransmission of DEREGISTRATION REQUEST message
	T3555sec           int    `yaml:"t3555 sec"` //Retransmission of CONFIGURATION UPDATE COMMAND message
	T3565sec           int    `yaml:"t3565 sec"` //Retransmission of NOTIFICATION message
}

type CmnTaiList struct {
	Plmn string
	Tac  []uint32
}

type CmSNssai struct {
	Sst uint8
	Sd  uint32
}
