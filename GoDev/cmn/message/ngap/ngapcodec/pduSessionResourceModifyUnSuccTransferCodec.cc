#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceModifyUnSuccTransferCodec.h"

using namespace std;

PduSessResMdfyUnSuccTransferCodec::PduSessResMdfyUnSuccTransferCodec() {
  reset();
}

PduSessResMdfyUnSuccTransferCodec::~PduSessResMdfyUnSuccTransferCodec() {}

void PduSessResMdfyUnSuccTransferCodec::setCause(Uint8 cType, Uint8 cValue)
{
	causeType_m = cType;
	causeValue_m = cValue;
}

Uint8 PduSessResMdfyUnSuccTransferCodec::getCauseType()
{
    return causeType_m;
}

Uint8 PduSessResMdfyUnSuccTransferCodec::getCauseValue()
{
    return causeValue_m;
}

//bool PduSessResMdfyUnSuccTransferCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSessResMdfyUnSuccTransferCodec::encode(NgapOssCtxt *ctxt)
{
	MsgBuffer msgBuffer;
	msgBuffer.length = 0;
	msgBuffer.value = NULL;

    // qos flow per tul information
    Cause* causePtr = msgPdu_m.getCauseEncodePtr();

	switch(causeType_m){
	case radioNetwork_chosen:
		causePtr->choice = radioNetwork_chosen;
		causePtr->u.radioNetwork = CauseRadioNetwork(causeValue_m);
	    break;
	case transport_chosen:
		causePtr->choice = transport_chosen;
		causePtr->u.transport = CauseTransport(causeValue_m);
	    break;
	case nas_chosen:
		causePtr->choice = nas_chosen;
		causePtr->u.nas = CauseNas(causeValue_m);
	    break;
	case protocol_chosen:
		causePtr->choice = protocol_chosen;
		causePtr->u.protocol = CauseProtocol(causeValue_m);
	    break;
	case misc_chosen:
		causePtr->choice = misc_chosen;
		causePtr->u.misc = CauseMisc(causeValue_m);
	    break;
	}

	msgPdu_m.encodeOssToPer(ctxt);

    OssBuf *perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;

}
bool PduSessResMdfyUnSuccTransferCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf)
{
	reset();

	OssBuf ossbuf;
	ossbuf.length = msgBuf.length;
	ossbuf.value = msgBuf.value;

	msgPdu_m.setPerBufRef_v(ossbuf);

	if (0 != msgPdu_m.decodePerToOss(ctxt)) {
	  return false;
	}


    Cause* causePtr = msgPdu_m.getCauseDecodePtr();

	switch (causePtr->choice){
		case radioNetwork_chosen:
			causeType_m = radioNetwork_chosen;
			causeValue_m = causePtr->u.radioNetwork;
			break;
		case transport_chosen:
			causeType_m = transport_chosen;
			causeValue_m = causePtr->u.transport;
			break;
		case nas_chosen:
			causeType_m = nas_chosen;
			causeValue_m = causePtr->u.nas;
			break;
		case protocol_chosen:
			causeType_m = protocol_chosen;
			causeValue_m = causePtr->u.protocol;
			break;
		case misc_chosen:
			causeType_m = misc_chosen;
			causeValue_m = causePtr->u.misc;
			break;
	}

}

bool PduSessResMdfyUnSuccTransferCodec::reset()
{
 	msgPdu_m.reset_v();
	causeType_m = 0;
	causeValue_m = 0;
}

std::string PduSessResMdfyUnSuccTransferCodec::toString()
{

  ostringstream oss;

  oss << "PduSessResMdfyUnSuccTransferCodec: type("<<(int)causeType_m<<"," <<(int)causeValue_m<<")";

  return oss.str();
}



