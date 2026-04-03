#ifndef __pagingCodec_H__
#define __pagingCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pagingPdu.h>

class PagingCodec
{
  public:

    PagingCodec();
    ~PagingCodec();

    void setUePagingIdentity(UePagingIdentity& ident);
    UePagingIdentity& getUePagingIdentity();

	void setPagingDrx(Uint8 pdrx);
	Uint8 getPagingDrx();
    bool  isPagingDrxPrst();

	void addTaiPagingList(TaiType& pTai);
    TaiPagingList& getTaiPagingList();

	void setPagingPriority(Uint8 pPriority);
	Uint8 getPagingPriority();
    bool  isPagingPriorityPrst();
  
    void setUeRadioCapForPaging(UeRadioCapForPaging& val);
    UeRadioCapForPaging& getUeRadioCapForPaging();
    bool isUeRadioCapForPagingPrst();

	void setPagingOrigin(Uint8 pOrigin);
	Uint8 getPagingOrigin();
    bool  isPagingOriginPrst();

    void setAssDateForPaging(AssDateForPaging& pDate);
    AssDateForPaging& getAssDateForPaging();
    bool  isAssDataForPagingPrst();

    MsgBuffer encode(NgapOssCtxt* ctxt);
    //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();

  private:

    PagingPdu msgPdu_m;

    UePagingIdentity     uePagingIdentity_m;
	Uint8                pagingDrx_m;
    bool                 pagingDrxPrst_m;
    TaiPagingList        taiPagingList_m;
    Uint8                pagingPriority_m;
    bool                 pagingPriorityPrst_m;
    UeRadioCapForPaging  ueRadioCapForPaging_m;
    bool                 ueRadioCapForPagingPrst_m;
    Uint8                pagingOrigin_m;
    bool                 pagingOriginPrst_m;
    AssDateForPaging     assDataForPaging_m;
    bool                 assDataForPagingPrst_m;
};

#endif