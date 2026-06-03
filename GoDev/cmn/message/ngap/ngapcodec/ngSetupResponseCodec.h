#ifndef __ngSetupResponseCodec_H__
#define __ngSetupResponseCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <ngSetupResponsePdu.h>

class NgSetupResponseCodec
{
public:
	NgSetupResponseCodec();
	~NgSetupResponseCodec();

	void setAmfName(std::string name);
	std::string getAmfName();

	void setRelativeAmfCapacity(Uint16 cap);
	Uint16 getRelativeAmfCapacity();

	void addServedGuamiList(ServedGuamiItem& sguami);
    ServedGuamiList& getServedGuamiList();

	void addPlmnList(BPlmnItem& plmnItem);
    BPlmnList& getPlmnList();

	void setUeRetentionInfo(Uint8 ueRetentionInfo);
	bool getUeRetentionInfoPrst();
	Uint8 getUeRetentionInfo();
	
	MsgBuffer encode(NgapOssCtxt* ctxt);
	//bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
	bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

	bool reset();
	std::string toString();
	
private:

    NGSetupResponsePdu msgPdu_m;

	std::string amfName_m;
	
	Uint16 relativeAmfCapacity_m;
	
	ServedGuamiList serGuamiList_m;
	
	BPlmnList plmnSupportist_m;

	bool  UeRetentionInfoPrst_m;
	Uint8 UeRetentionInfo_m;
}; 
#endif