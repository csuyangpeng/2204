#include <iostream>
#include <sstream>
#include <string>
#include "ngResetCodec.h"

using namespace std;

NgResetCodec::NgResetCodec()
{
    reset();
}

NgResetCodec::~NgResetCodec()
{
}

void 
NgResetCodec::setCauseType(Uint16 cType)
{
	causeType_m = cType;
}

Uint16 
NgResetCodec::getCauseType()
{
	return causeType_m;
}

void 
NgResetCodec::setCauseValue(Uint16 value)
{
	causeValue_m = value;
}

Uint16 
NgResetCodec::getCauseValue()
{
	return causeValue_m;
}

void NgResetCodec::setResetType(ResetTypeChoice& resetTypeData)
{
    resetType_m = resetTypeData;
}

ResetTypeChoice& NgResetCodec::getResetType()
{
    return resetType_m;
}

bool NgResetCodec::reset()
{
    msgPdu_m.reset_v();
    
	causeType_m = 0;
    causeValue_m = 0;

    resetType_m.ngInterface = 0;
    resetType_m.ngInterfacePrst = false;
    resetType_m.ueAssLogicalNgConnList.clear();
    resetType_m.ueAssLogicalNgConnListPrst = false;

    return true;
}

MsgBuffer NgResetCodec::encode(NgapOssCtxt* ctxt)
{
    std::cout << "NgResetCodec::encode()" << endl;
    
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

   //encode Cause Type and Cause Value
   Cause* causePtr = msgPdu_m.getCauseEncodePtr();
	switch (causeType_m)
		{
		case radioNetwork_chosen:
			causePtr->choice = causeType_m;
			causePtr->u.radioNetwork = (CauseRadioNetwork)causeValue_m;
			break;
		case transport_chosen:
			causePtr->choice = causeType_m;
            causePtr->u.transport = (CauseTransport)causeValue_m;
            break;
		case nas_chosen:
			causePtr->choice = causeType_m;
			causePtr->u.nas = (CauseNas)causeValue_m;
			break;
		case protocol_chosen:
			causePtr->choice = causeType_m;
            causePtr->u.protocol = (CauseProtocol)causeValue_m;
            break;
		case misc_chosen:
			causePtr->choice = causeType_m;
			causePtr->u.misc = (CauseMisc)causeValue_m;
			break;

		default:
		break;
		}

    //encode Reset Type
    ResetType* resetTypePtr = msgPdu_m.getResetTypeEncodePtr();
    if(resetType_m.ngInterfacePrst == true) {
        resetTypePtr->choice = nG_Interface_chosen;
        resetTypePtr->u.nG_Interface = ResetAll(resetType_m.ngInterface);
    } 

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool NgResetCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
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

	//decode Cause Type and Cause Value
    causeType_m=(msgPdu_m.getCauseDecodedPtr())->choice;
	switch((msgPdu_m.getCauseDecodedPtr())->choice){
		case radioNetwork_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).radioNetwork;
		break;
		case transport_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).transport;
		break;
		case nas_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).nas;
		break;
		case protocol_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).protocol;
		break;
		case misc_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).misc;
		break;
	}

    // decode Reset Type
    ResetType* resetTypePtr = msgPdu_m.getResetTypeDecodedPtr();
    if (resetTypePtr->choice == nG_Interface_chosen) {
        resetType_m.ngInterface = resetTypePtr->u.nG_Interface;
        resetType_m.ngInterfacePrst = true;
    }


    return true;
}

string NgResetCodec::toString()
{
    ostringstream oss;
    oss <<"Dump NgResetCodec:" << endl
		 <<" CauseType: "<<causeType_m
		 <<" CauseValue: "<<causeValue_m <<endl;
    oss <<"Reset Type: "<< endl
         <<"reset type prst:" << resetType_m.ngInterfacePrst 
         <<" ,value:" <<int(resetType_m.ngInterface) <<endl;
    return oss.str();
}

