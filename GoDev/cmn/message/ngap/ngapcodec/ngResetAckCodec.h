#ifndef __ngResetAckCodec_H__
#define __ngResetAckCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <ngResetAckPdu.h>

class NgResetAckCodec
{
public:
	NgResetAckCodec();
	~NgResetAckCodec();

    void addUeAssLogicalNgConnList(UeAssLogicalNgConn &val);
    UeAssLogicalNgConnList& getUeAssLogicalNgConnList();
    bool isUeAssLogicalNgConnPrst();

	std::string toString();

	MsgBuffer encode(NgapOssCtxt* ctxt);

	bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);
	
	bool reset();

private:
    NGResetAckPdu msgPdu_m;

    UeAssLogicalNgConnList ueAssLogicalNgConnList_m;
    bool ueAssLogicalNgConnPrst_m;
}; 
#endif