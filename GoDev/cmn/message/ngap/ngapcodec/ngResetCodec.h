#ifndef __ngResetCodec_H__
#define __ngResetCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <ngResetPdu.h>

class NgResetCodec
{
public:
	NgResetCodec();
	~NgResetCodec();

	void setCauseType(Uint16 cType);
	Uint16 getCauseType();

    void setCauseValue(Uint16 value);
	Uint16 getCauseValue();

	void setResetType(ResetTypeChoice& resetTypeData);
	ResetTypeChoice& getResetType();
    
	std::string toString();

	MsgBuffer encode(NgapOssCtxt* ctxt);

	bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);
	
	bool reset();

private:

    NGResetPdu msgPdu_m;

	Uint16 causeType_m;
    Uint16 causeValue_m;
	
	ResetTypeChoice resetType_m;
}; 
#endif