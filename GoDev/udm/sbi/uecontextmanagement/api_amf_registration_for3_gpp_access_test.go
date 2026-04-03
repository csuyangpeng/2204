package UEContextManagement_test

// GetAmf3gppAccess - retrieve the AMF registration for 3GPP access information
//func TestRegistrationAmf3gppAccess(t *testing.T) {
//	go func() { // udm server
//		router := gin.Default()
//		Nudm_UECM_Server.AddService(router)
//
//		udmLogPath := path_util.Gofree5gcPath("free5gc/udmsslkey.log")
//		udmPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.pem")
//		udmKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.key")
//
//		server, err := http2_util.NewServer(":29503", udmLogPath, router)
//		if err == nil && server != nil {
//			rlogger.InitLog.Infoln(server.ListenAndServeTLS(udmPemPath, udmKeyPath))
//			assert.True(t, err == nil)
//		}
//	}()
//	udm_util.testInitUdmConfig()
//	go udm_handler.Handle()
//
//	go func() { // fake udr server
//		router := gin.Default()
//
//		router.PUT("/nudr-dr/v1/subscription-data/:ueId/context-data/amf-3gpp-access", func(c *gin.Context) {
//			ueId := c.Param("ueId")
//			fmt.Println("==========AMF registration for 3GPP access==========")
//			fmt.Println("ueId: ", ueId)
//
//			var amf3GppAccessRegistration models.Amf3GppAccessRegistration
//			if err := c.ShouldBindJSON(&amf3GppAccessRegistration); err != nil {
//				fmt.Println("fake udr server error")
//				c.JSON(http.StatusInternalServerError, gin.H{})
//				return
//			}
//			fmt.Println("amf3GppAccessRegistration - ", amf3GppAccessRegistration.AmfInstanceId)
//			c.JSON(http.StatusNoContent, gin.H{})
//		})
//
//		udrLogPath := path_util.Gofree5gcPath("free5gc/udrsslkey.log")
//		udrPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.pem")
//		udrKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.key")
//
//		server, err := http2_util.NewServer(":29504", udrLogPath, router)
//		if err == nil && server != nil {
//			rlogger.InitLog.Infoln(server.ListenAndServeTLS(udrPemPath, udrKeyPath))
//			assert.True(t, err == nil)
//		}
//	}()
//
//	udm_context.Init()
//	cfg := Nudm_UECM_Client.NewConfiguration()
//	cfg.SetBasePath("https://localhost:29503")
//	clientAPI := Nudm_UECM_Client.NewAPIClient(cfg)
//
//	ueId := "UECM1234"
//	var amf3GppAccessRegistration models.Amf3GppAccessRegistration
//	amf3GppAccessRegistration.AmfInstanceId = "PUT_TEST_001"
//	_, resp, err := clientAPI.AMFRegistrationFor3GPPAccessApi.Registration(context.Background(), ueId, amf3GppAccessRegistration)
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		fmt.Println("resp: ", resp)
//	}
//}
