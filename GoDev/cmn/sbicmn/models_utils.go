package sbicmn

import (
	"lite5gc/openapi/models"
	"time"
)

func NewN1N2MessageTransferReqData() *models.N1N2MessageTransferReqData {

	modelData := &models.N1N2MessageTransferReqData{}

	modelData.N1MessageContainer = &models.N1MessageContainer{}
	modelData.N1MessageContainer.N1MessageContent = &models.RefToBinaryData{}

	modelData.N2InfoContainer = &models.N2InfoContainer{}
	modelData.N2InfoContainer.SmInfo = &models.N2SmInformation{}
	modelData.N2InfoContainer.SmInfo.SNssai = &models.Snssai{}
	modelData.N2InfoContainer.SmInfo.N2InfoContent = &models.N2InfoContent{}
	modelData.N2InfoContainer.SmInfo.N2InfoContent.NgapData = &models.RefToBinaryData{}

	//modelData.N2InfoContainer.RanInfo = &models.N2RanInformation{}
	//modelData.N2InfoContainer.RanInfo.N2InfoContent = &models.N2InfoContent{}
	//modelData.N2InfoContainer.RanInfo.N2InfoContent.NgapData = &models.RefToBinaryData{}
	//
	//modelData.N2InfoContainer.NrppaInfo = &models.NrppaInformation{}
	//modelData.N2InfoContainer.NrppaInfo.NrppaPdu = &models.N2InfoContent{}
	//modelData.N2InfoContainer.NrppaInfo.NrppaPdu.NgapData = &models.RefToBinaryData{}
	//
	//modelData.N2InfoContainer.PwsInfo = &models.PwsInformation{}
	//modelData.N2InfoContainer.PwsInfo.PwsContainer = &models.N2InfoContent{}
	//modelData.N2InfoContainer.PwsInfo.PwsContainer.NgapData = &models.RefToBinaryData{}
	//
	//modelData.Arp = &models.Arp{}
	//
	//modelData.AreaOfValidity = &models.AreaOfValidity{}

	return modelData
}

func NewSmContextUpdatedData() *models.SmContextUpdatedData {
	modelData := &models.SmContextUpdatedData{}

	modelData.N1SmMsg = &models.RefToBinaryData{}
	modelData.N2SmInfo = &models.RefToBinaryData{}
	return modelData
}

func NewUpdateSmContextRequest() *models.UpdateSmContextRequest {
	mData := &models.UpdateSmContextRequest{}
	mData.JsonData = &models.SmContextUpdateData{}
	mData.JsonData.Guami = &models.Guami{}
	mData.JsonData.Guami.PlmnId = &models.PlmnId{}
	mData.JsonData.ServingNetwork = &models.PlmnId{}

	mData.JsonData.UeLocation = &models.UserLocation{}
	mData.JsonData.UeLocation.EutraLocation = &models.EutraLocation{}
	mData.JsonData.UeLocation.EutraLocation.Tai = &models.Tai{}
	mData.JsonData.UeLocation.EutraLocation.Tai.PlmnId = &models.PlmnId{}
	mData.JsonData.UeLocation.EutraLocation.Ecgi = &models.Ecgi{}
	mData.JsonData.UeLocation.EutraLocation.Ecgi.PlmnId = &models.PlmnId{}
	mData.JsonData.UeLocation.EutraLocation.UeLocationTimestamp = &time.Time{}
	mData.JsonData.UeLocation.EutraLocation.GlobalNgenbId = &models.GlobalRanNodeId{}
	mData.JsonData.UeLocation.EutraLocation.GlobalNgenbId.PlmnId = &models.PlmnId{}
	mData.JsonData.UeLocation.EutraLocation.GlobalNgenbId.GNbId = &models.GNbId{}
	mData.JsonData.UeLocation.N3gaLocation = &models.N3gaLocation{}
	mData.JsonData.UeLocation.N3gaLocation.N3gppTai = &models.Tai{}
	mData.JsonData.UeLocation.N3gaLocation.N3gppTai.PlmnId = &models.PlmnId{}
	mData.JsonData.UeLocation.NrLocation = &models.NrLocation{}
	mData.JsonData.UeLocation.NrLocation.Tai = &models.Tai{}
	mData.JsonData.UeLocation.NrLocation.Tai.PlmnId = &models.PlmnId{}
	mData.JsonData.UeLocation.NrLocation.Ncgi = &models.Ncgi{}
	mData.JsonData.UeLocation.NrLocation.Ncgi.PlmnId = &models.PlmnId{}
	mData.JsonData.UeLocation.NrLocation.UeLocationTimestamp = &time.Time{}
	mData.JsonData.UeLocation.NrLocation.GlobalGnbId = &models.GlobalRanNodeId{}
	mData.JsonData.UeLocation.NrLocation.GlobalGnbId.PlmnId = &models.PlmnId{}
	mData.JsonData.UeLocation.NrLocation.GlobalGnbId.GNbId = &models.GNbId{}

	mData.JsonData.AddUeLocation = &models.UserLocation{}
	mData.JsonData.AddUeLocation.EutraLocation = &models.EutraLocation{}
	mData.JsonData.AddUeLocation.EutraLocation.Tai = &models.Tai{}
	mData.JsonData.AddUeLocation.EutraLocation.Tai.PlmnId = &models.PlmnId{}
	mData.JsonData.AddUeLocation.EutraLocation.Ecgi = &models.Ecgi{}
	mData.JsonData.AddUeLocation.EutraLocation.Ecgi.PlmnId = &models.PlmnId{}
	mData.JsonData.AddUeLocation.EutraLocation.UeLocationTimestamp = &time.Time{}
	mData.JsonData.AddUeLocation.EutraLocation.GlobalNgenbId = &models.GlobalRanNodeId{}
	mData.JsonData.AddUeLocation.EutraLocation.GlobalNgenbId.PlmnId = &models.PlmnId{}
	mData.JsonData.AddUeLocation.EutraLocation.GlobalNgenbId.GNbId = &models.GNbId{}
	mData.JsonData.AddUeLocation.N3gaLocation = &models.N3gaLocation{}
	mData.JsonData.AddUeLocation.N3gaLocation.N3gppTai = &models.Tai{}
	mData.JsonData.AddUeLocation.N3gaLocation.N3gppTai.PlmnId = &models.PlmnId{}
	mData.JsonData.AddUeLocation.NrLocation = &models.NrLocation{}
	mData.JsonData.AddUeLocation.NrLocation.Tai = &models.Tai{}
	mData.JsonData.AddUeLocation.NrLocation.Tai.PlmnId = &models.PlmnId{}
	mData.JsonData.AddUeLocation.NrLocation.Ncgi = &models.Ncgi{}
	mData.JsonData.AddUeLocation.NrLocation.Ncgi.PlmnId = &models.PlmnId{}
	mData.JsonData.AddUeLocation.NrLocation.UeLocationTimestamp = &time.Time{}
	mData.JsonData.AddUeLocation.NrLocation.GlobalGnbId = &models.GlobalRanNodeId{}
	mData.JsonData.AddUeLocation.NrLocation.GlobalGnbId.PlmnId = &models.PlmnId{}
	mData.JsonData.AddUeLocation.NrLocation.GlobalGnbId.GNbId = &models.GNbId{}

	mData.JsonData.N1SmMsg = &models.RefToBinaryData{}
	mData.JsonData.N2SmInfo = &models.RefToBinaryData{}

	mData.JsonData.TargetId = &models.NgRanTargetId{}
	mData.JsonData.TargetId.Tai = &models.Tai{}
	mData.JsonData.TargetId.Tai.PlmnId = &models.PlmnId{}
	mData.JsonData.TargetId.RanNodeId = &models.GlobalRanNodeId{}
	mData.JsonData.TargetId.RanNodeId.PlmnId = &models.PlmnId{}
	mData.JsonData.TargetId.RanNodeId.GNbId = &models.GNbId{}

	mData.JsonData.NgApCause = &models.NgApCause{}

	mData.JsonData.SNssai = &models.Snssai{}

	mData.JsonData.TraceData = &models.TraceData{}

	return mData
}
