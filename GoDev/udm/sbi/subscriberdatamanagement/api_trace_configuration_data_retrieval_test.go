/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 11/23/20 9:36 PM
 * Description:
 */

package subscriberdatamanagement_test

import (
	"testing"
)

// GetTraceData - retrieve a UE's Trace Configuration Data
func TestGetTraceData(t *testing.T) {

	//go func() { // udm server
	//	router := gin.Default()
	//	Nudm_SDM_Server.AddService(router)
	//
	//	udmLogPath := path_util.Gofree5gcPath("free5gc/udmsslkey.log")
	//	udmPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.pem")
	//	udmKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.key")
	//
	//	server, err := http2_util.NewServer(":29503", udmLogPath, router)
	//	if err == nil && server != nil {
	//		logger.InitLog.Infoln(server.ListenAndServeTLS(udmPemPath, udmKeyPath))
	//		assert.True(t, err == nil)
	//	}
	//}()
	//
	//udm_util.testInitUdmConfig()
	//go udm_handler.Handle()
	//
	//go func() { // fake udr server
	//	router := gin.Default()
	//
	//	router.GET("/nudr-dr/v1/subscription-data/:ueId/:servingPlmnId/provisioned-data/trace-data", func(c *gin.Context) {
	//		supi := c.Param("supi")
	//		fmt.Println("==========GetTraceData - retrieve a UE's Trace Configuration Data==========")
	//		fmt.Println("supi: ", supi)
	//
	//		var traceData models.TraceData
	//		traceData.TraceRef = "Test_00"
	//		fmt.Println("traceData - ", traceData.TraceRef)
	//		c.JSON(http.StatusNoContent, gin.H{})
	//	})
	//
	//	udrLogPath := path_util.Gofree5gcPath("free5gc/udrsslkey.log")
	//	udrPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.pem")
	//	udrKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.key")
	//
	//	server, err := http2_util.NewServer(":29504", udrLogPath, router)
	//	if err == nil && server != nil {
	//		logger.InitLog.Infoln(server.ListenAndServeTLS(udrPemPath, udrKeyPath))
	//		assert.True(t, err == nil)
	//	}
	//}()
	//
	//udm_context.Init()
	//cfg := Nudm_SDM_Client.NewConfiguration()
	//cfg.SetBasePath("https://localhost:29503")
	//clientAPI := Nudm_SDM_Client.NewAPIClient(cfg)
	//
	//supi := "SDM1234"
	//_, resp, err := clientAPI.TraceConfigurationDataRetrievalApi.GetTraceData(context.TODO(), supi, nil)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println("resp: ", resp)
	//}
}
