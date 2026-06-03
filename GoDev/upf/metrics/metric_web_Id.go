package metrics

/*
import (
	"lite5gc/oam/agent/webTypes"
	"lite5gc/cmn/types/configure"
)

const (
	numTotalPackets                          int = 4040010
	numTotalBits                             int = 4040020
	numTotalPacketsSec                       int = 4040030
	numTotalBitsSec                          int = 4040040
	numTotalDiscardedPackets                 int = 4040050
	numUpLinkTotalPacketsSent                int = 4041010
	numUpLinkPacketsSentSec                  int = 4041020
	numUpLinkTotalPacketsReceived            int = 4041030
	numUpLinkPacketsReceivedSec              int = 4041040
	numUpLinkTotalBitsSent                   int = 4041050
	numUpLinkBitsSentSec                     int = 4041060
	numUpLinkTotalBitsReceived               int = 4041070
	numUpLinkBitsReceivedSec                 int = 4041080
	numUpLinkTotalReceivedPacketsDiscarded   int = 4041090
	numUpLinkTotalSentPacketsDiscarded       int = 4041100
	numUpLinkTotalInvalidPackets             int = 4041110
	numUpLinkTotalDiscardedPackets           int = 4041120
	numDownLinkTotalPacketsSent              int = 4042010
	numDownLinkPacketsSentSec                int = 4042020
	numDownLinkTotalPacketsReceived          int = 4042030
	numDownLinkPacketsReceivedSec            int = 4042040
	numDownLinkTotalBitsSent                 int = 4042050
	numDownLinkBitsSentSec                   int = 4042060
	numDownLinkTotalBitsReceived             int = 4042070
	numDownLinkBitsReceivedSec               int = 4042080
	numDownLinkTotalReceivedPacketsDiscarded int = 4042090
	numDownLinkTotalSentPacketsDiscarded     int = 4042100
	numDownLinkTotalInvalidPackets           int = 4042110
	numDownLinkTotalDiscardedPackets         int = 4042120
)

func ReportToLems() *webTypes.NfPerformanceData {
	pms := webTypes.NewNfPerformanceData()
	pms.NfNo = configure.LemsUpfAgentConf.NfNo.UpfNfNo
	//pms.Params = append(pms.Params, webTypes.Param{111, 100})
	var params = make([]webTypes.Param, 0, 30)

	// upf module counter
	params = append(params, webTypes.Param{numTotalPackets, UpfmoduleSet.TotalPackets.Count()})
	params = append(params, webTypes.Param{numTotalBits, UpfmoduleSet.TotalBits.Count()})
	params = append(params, webTypes.Param{numTotalDiscardedPackets, UpfmoduleSet.TotalDiscardedPackets.Count()})
	params = append(params, webTypes.Param{numUpLinkTotalPacketsSent, UpfmoduleSet.UpLinkTotalPacketsSent.Count()})
	params = append(params, webTypes.Param{numUpLinkTotalPacketsReceived, UpfmoduleSet.UpLinkTotalPacketsReceived.Count()})
	params = append(params, webTypes.Param{numUpLinkTotalBitsSent, UpfmoduleSet.UpLinkTotalBitsSent.Count() * 8})
	params = append(params, webTypes.Param{numUpLinkTotalBitsReceived, UpfmoduleSet.UpLinkTotalBitsReceived.Count() * 8})
	params = append(params, webTypes.Param{numUpLinkTotalReceivedPacketsDiscarded, UpfmoduleSet.UpLinkTotalReceivedPacketsDiscarded.Count()})
	params = append(params, webTypes.Param{numUpLinkTotalSentPacketsDiscarded, UpfmoduleSet.UpLinkTotalSentPacketsDiscarded.Count()})
	params = append(params, webTypes.Param{numUpLinkTotalInvalidPackets, UpfmoduleSet.UpLinkTotalInvalidPackets.Count()})
	params = append(params, webTypes.Param{numUpLinkTotalDiscardedPackets, UpfmoduleSet.UpLinkTotalDiscardedPackets.Count()})
	params = append(params, webTypes.Param{numDownLinkTotalPacketsSent, UpfmoduleSet.DownLinkTotalPacketsSent.Count()})
	params = append(params, webTypes.Param{numDownLinkTotalPacketsReceived, UpfmoduleSet.DownLinkTotalPacketsReceived.Count()})
	params = append(params, webTypes.Param{numDownLinkTotalBitsSent, UpfmoduleSet.DownLinkTotalBitsSent.Count() * 8})
	params = append(params, webTypes.Param{numDownLinkTotalBitsReceived, UpfmoduleSet.DownLinkTotalBitsReceived.Count() * 8})
	params = append(params, webTypes.Param{numDownLinkTotalReceivedPacketsDiscarded, UpfmoduleSet.DownLinkTotalReceivedPacketsDiscarded.Count()})
	params = append(params, webTypes.Param{numDownLinkTotalSentPacketsDiscarded, UpfmoduleSet.DownLinkTotalSentPacketsDiscarded.Count()})
	params = append(params, webTypes.Param{numDownLinkTotalInvalidPackets, UpfmoduleSet.DownLinkTotalInvalidPackets.Count()})
	params = append(params, webTypes.Param{numDownLinkTotalDiscardedPackets, UpfmoduleSet.DownLinkTotalDiscardedPackets.Count()})

	// upf module meter P-I
	params = append(params, webTypes.Param{numTotalPacketsSec, UpfmodulePISet.TotalPacketsPerSec.Rate1()})
	params = append(params, webTypes.Param{numTotalBitsSec, UpfmodulePISet.TotalBitsPerSec.Rate1()})
	params = append(params, webTypes.Param{numUpLinkPacketsSentSec, UpfmodulePISet.UpLinkPacketsSentPerSec.Rate1()})
	params = append(params, webTypes.Param{numUpLinkPacketsReceivedSec, UpfmodulePISet.UpLinkPacketsReceivedPerSec.Rate1()})
	params = append(params, webTypes.Param{numUpLinkBitsSentSec, UpfmodulePISet.UpLinkBitsSentPerSec.Rate1()})
	params = append(params, webTypes.Param{numUpLinkBitsReceivedSec, UpfmodulePISet.UpLinkBitsReceivedPerSec.Rate1()})
	params = append(params, webTypes.Param{numDownLinkPacketsSentSec, UpfmodulePISet.DownLinkPacketsSentPerSec.Rate1()})
	params = append(params, webTypes.Param{numDownLinkPacketsReceivedSec, UpfmodulePISet.DownLinkPacketsReceivedPerSec.Rate1()})
	params = append(params, webTypes.Param{numDownLinkBitsSentSec, UpfmodulePISet.DownLinkBitsSentPerSec.Rate1()})
	params = append(params, webTypes.Param{numDownLinkBitsReceivedSec, UpfmodulePISet.DownLinkBitsReceivedPerSec.Rate1()})

	pms.Params = params

	return pms
}

// report 5 second
// func StartPMReport() {
// 	agent.StartPerformanceReport(5e9, ReportToLems)
// }
*/
