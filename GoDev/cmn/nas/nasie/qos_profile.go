package nasie

import t3 "lite5gc/cmn/types3gpp"

type QosProfile struct {
	//mandatory
	//Default 5G QoS identifier
	QI5 QI5Contents

	//allocation and retention priority
	Arp t3.ARP

	// priority level
	IsPriorityLevelPrst bool
	PriorityLevel       uint8
	//optional
	//This attribute may only be used for a standardized or pre-configured 5QI.
	// When present, this attribute provides QoS characteristics that
	// override the default values for a standardized or pre-configured 5QI,
	IsNonDync5qiPrst bool
	NonDynamic5qi    t3.NonDynamic5QI

	//This attribute shall only be used for dynamically-assigned 5QIs.
	// When present, this attribute provides an explicit set of QoS characteristics.
	IsDync5qiPrst bool
	Dynamic5qi    Dynamic5Qi
}
