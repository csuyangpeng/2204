package statistics

import (
	"github.com/rcrowley/go-metrics"
	"lite5gc/oam/pm"
)

var NgapRegistry metrics.Registry
var MmRegistry metrics.Registry
var SmRegistry metrics.Registry
var AmfStatisticRegistry metrics.Registry

var NGSetupRequestCounter metrics.Counter
var NGSetupResponseCounter metrics.Counter
var NGReSetCounter metrics.Counter
var InitialUEMessageCounter metrics.Counter
var InitialContextSetupRequestCounter metrics.Counter
var InitialContextSetupResponseCounter metrics.Counter
var UpLinkNASTransportCounter metrics.Counter
var DownLinkNASTransportCounter metrics.Counter
var PDUSessionResourceSetupRequestCounter metrics.Counter
var PDUSessionResourceSetupResponseCounter metrics.Counter
var UEContextReleaseRequestCounter metrics.Counter
var UEContextReleaseCommandCounter metrics.Counter
var UEContextReleaseCompleteCounter metrics.Counter
var UERadioCapabilityInfoIndicationCounter metrics.Counter
var PagingCounter metrics.Counter

var RegistrationRequestCounter metrics.Counter
var RegistrationAcceptCounter metrics.Counter
var RegistrationCompleteCounter metrics.Counter
var RegistrationRejectCounter metrics.Counter
var DeregistrationRequestUeCounter metrics.Counter
var DeregistrationAcceptUeCounter metrics.Counter

var ServiceRequestCounter metrics.Counter
var ServiceRejectCounter metrics.Counter
var ServiceAcceptCounter metrics.Counter
var ULNasTransportCounter metrics.Counter
var DLNasTransportCounter metrics.Counter
var AuthenticationRequestCounter metrics.Counter
var AuthenticationResponseCounter metrics.Counter
var AuthenticationFailureCounter metrics.Counter
var AuthenticationRejectCounter metrics.Counter
var SecurityModeCommandCounter metrics.Counter
var SecurityModeCompleteCounter metrics.Counter
var SecurityModeRejectCounter metrics.Counter
var IdentifyRequestCounter metrics.Counter
var IdentifyResponseCounter metrics.Counter

var PduSessEstablishRequestCounter metrics.Counter
var PduSessEstablishAcceptCounter metrics.Counter
var PduSessEstablishRejectCounter metrics.Counter
var PduSessReleaseRequestCounter metrics.Counter
var PduSessReleaseCommandCounter metrics.Counter
var PduSessReleaseCompleteCounter metrics.Counter

var OnLineUserCounter metrics.Counter
var OffLineUserCounter metrics.Counter
var ActiveUserCounter metrics.Counter
var IdleUserCounter metrics.Counter

var offlinePres int64

func Init() {
	NgapRegistry = metrics.NewRegistry()
	MmRegistry = metrics.NewRegistry()
	SmRegistry = metrics.NewRegistry()
	AmfStatisticRegistry = metrics.NewRegistry()

	NGSetupRequestCounter = pm.CreateCounter(NgapRegistry, NGSetupRequestCounterKey)
	NGSetupResponseCounter = pm.CreateCounter(NgapRegistry, NGSetupResponseCounterKey)
	NGReSetCounter = pm.CreateCounter(NgapRegistry, NGReSetCounterKey)
	InitialUEMessageCounter = pm.CreateCounter(NgapRegistry, InitialUEMessageCounterKey)
	InitialContextSetupRequestCounter = pm.CreateCounter(NgapRegistry, InitialContextSetupRequestCounterKey)
	InitialContextSetupResponseCounter = pm.CreateCounter(NgapRegistry, InitialContextSetupResponseCounterKey)
	UpLinkNASTransportCounter = pm.CreateCounter(NgapRegistry, UpLinkNASTransportCounterKey)
	DownLinkNASTransportCounter = pm.CreateCounter(NgapRegistry, DownLinkNASTransportCounterKey)
	PDUSessionResourceSetupRequestCounter = pm.CreateCounter(NgapRegistry, PDUSessionResourceSetupRequestCounterKey)
	PDUSessionResourceSetupResponseCounter = pm.CreateCounter(NgapRegistry, PDUSessionResourceSetupResponseCounterKey)
	UEContextReleaseRequestCounter = pm.CreateCounter(NgapRegistry, UEContextReleaseRequestCounterKey)
	UEContextReleaseCommandCounter = pm.CreateCounter(NgapRegistry, UEContextReleaseCommandCounterKey)
	UEContextReleaseCompleteCounter = pm.CreateCounter(NgapRegistry, UEContextReleaseCompleteCounterKey)
	UERadioCapabilityInfoIndicationCounter = pm.CreateCounter(NgapRegistry, UERadioCapabilityInfoIndicationCounterKey)
	PagingCounter = pm.CreateCounter(NgapRegistry, PagingCounterKey)

	RegistrationRequestCounter = pm.CreateCounter(MmRegistry, RegistrationRequestCounterKey)
	RegistrationAcceptCounter = pm.CreateCounter(MmRegistry, RegistrationAcceptCounterKey)
	RegistrationCompleteCounter = pm.CreateCounter(MmRegistry, RegistrationCompleteCounterKey)
	RegistrationRejectCounter = pm.CreateCounter(MmRegistry, RegistrationRejectCounterKey)
	DeregistrationRequestUeCounter = pm.CreateCounter(MmRegistry, DeregistrationRequestUeCounterKey)
	DeregistrationAcceptUeCounter = pm.CreateCounter(MmRegistry, DeregistrationAcceptUeCounterKey)
	ServiceRequestCounter = pm.CreateCounter(MmRegistry, ServiceRequestCounterKey)
	ServiceRejectCounter = pm.CreateCounter(MmRegistry, ServiceRejectCounterKey)
	ServiceAcceptCounter = pm.CreateCounter(MmRegistry, ServiceAcceptCounterKey)
	ULNasTransportCounter = pm.CreateCounter(MmRegistry, ULNasTransportCounterKey)
	DLNasTransportCounter = pm.CreateCounter(MmRegistry, DLNasTransportCounterKey)
	AuthenticationRequestCounter = pm.CreateCounter(MmRegistry, AuthenticationRequestCounterKey)
	AuthenticationResponseCounter = pm.CreateCounter(MmRegistry, AuthenticationResponseCounterKey)
	AuthenticationFailureCounter = pm.CreateCounter(MmRegistry, AuthenticationFailureCounterKey)
	AuthenticationRejectCounter = pm.CreateCounter(MmRegistry, AuthenticationRejectCounterKey)
	SecurityModeCommandCounter = pm.CreateCounter(MmRegistry, SecurityModeCommandCounterKey)
	SecurityModeCompleteCounter = pm.CreateCounter(MmRegistry, SecurityModeCompleteCounterKey)
	SecurityModeRejectCounter = pm.CreateCounter(MmRegistry, SecurityModeRejectCounterKey)
	IdentifyRequestCounter = pm.CreateCounter(MmRegistry, IdentifyRequestCounterKey)
	IdentifyResponseCounter = pm.CreateCounter(MmRegistry, IdentifyResponseCounterKey)

	PduSessEstablishRequestCounter = pm.CreateCounter(SmRegistry, PduSessEstablishRequestCounterKey)
	PduSessEstablishAcceptCounter = pm.CreateCounter(SmRegistry, PduSessEstablishAcceptCounterKey)
	PduSessEstablishRejectCounter = pm.CreateCounter(SmRegistry, PduSessEstablishRejectCounterKey)
	PduSessReleaseRequestCounter = pm.CreateCounter(SmRegistry, PduSessReleaseRequestCounterKey)
	PduSessReleaseCommandCounter = pm.CreateCounter(SmRegistry, PduSessReleaseCommandCounterKey)
	PduSessReleaseCompleteCounter = pm.CreateCounter(SmRegistry, PduSessReleaseCompleteCounterKey)

	OnLineUserCounter = pm.CreateCounter(AmfStatisticRegistry, OnLineUserCounterKey)
	OffLineUserCounter = pm.CreateCounter(AmfStatisticRegistry, OffLineUserCounterKey)
	ActiveUserCounter = pm.CreateCounter(AmfStatisticRegistry, ActiveUserCounterKey)
	IdleUserCounter = pm.CreateCounter(AmfStatisticRegistry, IdleUserCounterKey)

	offlinePres = 0

	return
}
