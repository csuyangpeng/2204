#include <iostream>
#include <sstream>
#include <string>

#include "ngSetupFailureCodec.h"

using namespace std;

NgSetupFailureCodec::NgSetupFailureCodec()
{
    reset();
}

NgSetupFailureCodec::~NgSetupFailureCodec()
{
}

bool
NgSetupFailureCodec::reset()
{
    msgPdu_m.reset_v();

	causeType_m = 0;

    causeValue_m = 0;

    relativeTimeToWait_m = 0;
	timeToWaitPrst_m = false;

	return true;
}

void 
NgSetupFailureCodec::setCauseType(Uint16 cType)
{
	causeType_m = cType;
}

Uint16 
NgSetupFailureCodec::getCauseType()
{
	return causeType_m;
}

void 
NgSetupFailureCodec::setCauseValue(Uint16 value)
{
	causeValue_m = value;
}

Uint16 
NgSetupFailureCodec::getCauseValue()
{
	return causeValue_m;
}

void 
NgSetupFailureCodec::setRelativeTimeToWait(Uint32 time)
{
    timeToWaitPrst_m = true;
	relativeTimeToWait_m = time;
}

Uint32 
NgSetupFailureCodec::getRelativeTimeToWait()
{
	return relativeTimeToWait_m;
}

bool 
NgSetupFailureCodec::isTimeToWaitPrst()
{
	return timeToWaitPrst_m;
}


//bool NgSetupFailureCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer NgSetupFailureCodec::encode(NgapOssCtxt* ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

   //encode Time To Wait
   if(timeToWaitPrst_m)
   {
       TimeToWait* timeToWaitPtr = msgPdu_m.getTimeToWaitEncodePtr();
       *timeToWaitPtr=(TimeToWait)relativeTimeToWait_m;
   }

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

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;

}

bool 
NgSetupFailureCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
{
	reset();
	
    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;
    
    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)){
		cout << "Failed to decode Per to OSS" << endl;
		return false;
	}
    
	//decode Time To Wait
	if(msgPdu_m.isTimeToWaitPresent())
	{
	    relativeTimeToWait_m=*(msgPdu_m.getTimeToWaitDecodedPtr());
		timeToWaitPrst_m = true;
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
	return true;
}

string NgSetupFailureCodec::toString()
{
	ostringstream oss;

    oss <<"NgSetupFailureCodec:" << endl
        <<"RelativeTimeToWait:prst "<<timeToWaitPrst_m<<" value: "<<relativeTimeToWait_m<<endl
		<<" CauseType: "<<causeType_m
		<<" CauseValue: "<<causeValue_m;

	return oss.str();
}
