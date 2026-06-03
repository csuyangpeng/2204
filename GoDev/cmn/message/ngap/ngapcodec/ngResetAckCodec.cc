#include <iostream>
#include <sstream>
#include <string>
#include "ngResetAckCodec.h"

using namespace std;

NgResetAckCodec::NgResetAckCodec()
{
    reset();
}

NgResetAckCodec::~NgResetAckCodec()
{
}

bool NgResetAckCodec::isUeAssLogicalNgConnPrst()
{
    return ueAssLogicalNgConnPrst_m;
}

void NgResetAckCodec::addUeAssLogicalNgConnList(UeAssLogicalNgConn& val) 
{
    ueAssLogicalNgConnPrst_m = true;
    ueAssLogicalNgConnList_m.push_back(val);
}

UeAssLogicalNgConnList& NgResetAckCodec::getUeAssLogicalNgConnList() 
{
    return ueAssLogicalNgConnList_m;
}


bool NgResetAckCodec::reset()
{
    msgPdu_m.reset_v();
    
    ueAssLogicalNgConnPrst_m = false;
    ueAssLogicalNgConnList_m.clear();

    return true;
}

MsgBuffer NgResetAckCodec::encode(NgapOssCtxt* ctxt)
{
    std::cout << "NgResetCodec::encode()" << endl;
    
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    // encode Ue Associate Logical NG-connection
    if(ueAssLogicalNgConnPrst_m)
    {
        unsigned int numofUeAssLogicalNgConn = ueAssLogicalNgConnList_m.size();
        msgPdu_m.setNumOfUEAssNGConnListItem_v(numofUeAssLogicalNgConn);
        for (int i = 0; i < numofUeAssLogicalNgConn; i++) 
        {
            UE_associatedLogicalNG_connectionItem* itemPtr = msgPdu_m.getUEAssNGConnListItemEncodePtr(i);
            if (itemPtr != NULL) 
            {
                if (ueAssLogicalNgConnList_m[i].amfUeNgapIdPrst) 
                {
                    itemPtr->bit_mask |= aMF_UE_NGAP_ID_present;
                    itemPtr->aMF_UE_NGAP_ID = ueAssLogicalNgConnList_m[i].amfUeNgapId;
                }
                if (ueAssLogicalNgConnList_m[i].ranUeNgapIdPrst) 
                {
                    itemPtr->bit_mask |= rAN_UE_NGAP_ID_present;
                    itemPtr->rAN_UE_NGAP_ID = ueAssLogicalNgConnList_m[i].ranUeNgapId;
                }
            }
	    }
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool NgResetAckCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
{
	
    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;
    
    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)){
		std::cout << "Failed to decode Per to OSS" << endl;
		return false;
	}

	//decode Ue Associate Logical NG-connection
	if (msgPdu_m.isUEAssNGConnPresent())
	{
        for (int i = 0; i < msgPdu_m.getNumOfUEAssNGConnListItem(); i++)
        {
		    UE_associatedLogicalNG_connectionItem* ptr = msgPdu_m.getUEAssNGConnListItemDecodedPtr(i);

            UeAssLogicalNgConn ueAssLogicalNgConn;

            if (ptr->bit_mask & aMF_UE_NGAP_ID_present) {
                ueAssLogicalNgConn.amfUeNgapId = ptr->aMF_UE_NGAP_ID;
                ueAssLogicalNgConn.amfUeNgapIdPrst = true;
            }
            if (!(ptr->bit_mask & aMF_UE_NGAP_ID_present)) {
                ueAssLogicalNgConn.amfUeNgapIdPrst = false;
            }
            if (ptr->bit_mask & rAN_UE_NGAP_ID_present) {
                ueAssLogicalNgConn.ranUeNgapId = ptr->rAN_UE_NGAP_ID;
                ueAssLogicalNgConn.ranUeNgapIdPrst = true;
            }
            if (!(ptr->bit_mask & rAN_UE_NGAP_ID_present)) {
                ueAssLogicalNgConn.ranUeNgapIdPrst = false;
            }

            ueAssLogicalNgConnList_m.push_back(ueAssLogicalNgConn);
        }
        ueAssLogicalNgConnPrst_m = true;
    }


    return true;
}

string NgResetAckCodec::toString()
{
    ostringstream oss;
    oss <<"NgResetAckCodec:" << endl
        <<"Ue Associated Logical Ng Connection" << endl;
    vector<UeAssLogicalNgConn>::iterator it;
    for (it = ueAssLogicalNgConnList_m.begin();it < ueAssLogicalNgConnList_m.end(); it++) {
        oss <<"amf Ue Ngap Id Prst(" << it->amfUeNgapIdPrst <<"), value: " << it->amfUeNgapId <<endl
            <<"ran Ue Ngap Id Prst(" << it->ranUeNgapIdPrst <<"), value: " << it->ranUeNgapId <<endl;
    }
    return oss.str();
}

