#ifndef __pagingPdu_h__INCLUDED__
#define __pagingPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PagingPdu
// -- **************************************************************
// --
// -- PAGING ELEMENTARY PROCEDURE
// --
// -- **************************************************************

// -- **************************************************************
// --
// -- PAGING
// --
// -- **************************************************************

// Paging ::= SEQUENCE {
// 	protocolIEs		ProtocolIE-Container		{ {PagingIEs} },
// 	...
// }

// PagingIEs NGAP-PROTOCOL-IES ::= {
// 	{ ID id-UEPagingIdentity				CRITICALITY ignore	TYPE UEPagingIdentity				PRESENCE mandatory	}|
// 	{ ID id-PagingDRX						CRITICALITY ignore	TYPE PagingDRX						PRESENCE optional		}|
// 	{ ID id-TAIListForPaging				CRITICALITY ignore	TYPE TAIListForPaging				PRESENCE mandatory	}|
// 	{ ID id-PagingPriority					CRITICALITY ignore	TYPE PagingPriority					PRESENCE optional		}|
// 	{ ID id-UERadioCapabilityForPaging		CRITICALITY ignore	TYPE UERadioCapabilityForPaging		PRESENCE optional		}|
// 	{ ID id-PagingOrigin					CRITICALITY ignore	TYPE PagingOrigin					PRESENCE optional		}|
// 	{ ID id-AssistanceDataForPaging			CRITICALITY ignore	TYPE AssistanceDataForPaging		PRESENCE optional		},
// 	...
// }
///////////////////////////////////////////////////////////////////////////////

class PagingPdu:public NgapOssPdu
{
  public:
    PagingPdu();
    //virtual
    ~PagingPdu();
    
    /////////////////////////////////////////////////////////////
    /// @brief reset_v
    ///  resets all data members to default value
    /////////////////////////////////////////////////////////////
    //virtual
    void reset_v();

    /////////////////////////////////////////////////////////////
    /// @brief encode Oss to PER using ASN.1
    ///  builds and processes the OSS PDU to encode
    ///  it to PER using ASN.1
    /////////////////////////////////////////////////////////////
    //virtual
    unsigned int encodeOssToPer(NgapOssCtxt* ctxt);

    /////////////////////////////////////////////////////////////
    /// @brief get Decoded Oss Data Pointers
    ///  processes the decoded OSS buffer and populates
    ///  the data access pointers in the class for the IEs present in the
    ///  decoded OSS buffer
    /////////////////////////////////////////////////////////////
    //virtual
    void getDecodedOssDataPtrs_v();

    /////////////////////////////////////////////////////////////
    /// @brief get Paging pointer
    /// returns the pointer to the
    /// Paging in the decoded message
    /////////////////////////////////////////////////////////////
    Paging * getPagingDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// UEPagingIdentity
    /////////////////////////////////////////////////////////////
    UEPagingIdentity* getUePagingIdentityDecodedPtr() const;
    UEPagingIdentity* getUePagingIdentityEncodePtr();

    /////////////////////////////////////////////////////////////
    /// PagingDRX
    /////////////////////////////////////////////////////////////
    PagingDRX * getPagingDrxDecodedPtr() const;
    PagingDRX * getPagingDrxEncodePtr();
    bool isPagingDrxPresent();

    /////////////////////////////////////////////////////////////
    /// TAIListForPaging
    /////////////////////////////////////////////////////////////
    TAIListForPaging* getTAIListForPagingDecodedPtr() const;
    TAIListForPaging* getTAIListForPagingEncodePtr();
    
    TAIListForPagingItem* getTAIListForPagingItemDecodedPtr(unsigned int index) const;
    TAIListForPagingItem* getTAIListForPagingItemEncodePtr(unsigned int index);

    void setNumofTAIListForPagingItem_v(unsigned int num);
    unsigned int getNumofTAIListForPagingItem() const;

    /////////////////////////////////////////////////////////////
    /// PagingPriority
    /////////////////////////////////////////////////////////////
    PagingPriority *getPagePriorityDecodedPtr() const;
    PagingPriority *getPagePriorityEncodePtr();
    bool isPagePriorityPresent();

    /////////////////////////////////////////////////////////////
    /// UERadioCapabilityForPaging
    /////////////////////////////////////////////////////////////
    UERadioCapabilityForPaging* getUeRadioCapForPagingDecodedPtr() const;
    UERadioCapabilityForPaging* getUeRadioCapForPagingEncodePtr();
    bool isUeRadioCapabilityForPagingPresent();

    /////////////////////////////////////////////////////////////
    /// PagingOrigin
    /////////////////////////////////////////////////////////////
    PagingOrigin* getPagingOriginDecodedPtr() const;
    PagingOrigin* getPagingOriginEncodePtr();
    bool isPagingOriginPresent();

    /////////////////////////////////////////////////////////////
    /// AssistanceDataForPaging
    /////////////////////////////////////////////////////////////
    AssistanceDataForPaging* getAssDataForPagingDecodedPtr() const;
    AssistanceDataForPaging* getAssDataForPagingEncodePtr();
    bool isAssDataForPagingPresent();
    
    RecommendedCellItem * getRecommendedCellItemDecodedPtr(unsigned int index) const;
    RecommendedCellItem * getRecommendedCellItemEncodePtr(unsigned int index);

    void setNumOfRecommendedCellItem_v(unsigned int num);
    unsigned int getNumOfRecommendedCellItem() const;

  private:
  
    typedef PagingIe ProtocolIeType;

    //OSS structures
    Paging *paging_mp; //for encode

    //UEPagingIdentity
    UEPagingIdentity uePagingIdentity_m;
    UEPagingIdentity* uePagingIdentity_mp;
    UePagingIdentity uePagingIdent_ma;

    //Paging DRX
    PagingDRX pagingDrx_m; 
    PagingDRX *pagingDrx_mp;
    bool pagingDrxPresent_m; 

    //TAIListForPaging
    TAIListForPaging taiListForPaging_m;
    struct TAIListForPaging_ taiListForPaging_ma[maxnoofTAIforPaging];
    unsigned int numofTAIListForPagingItem_m;
    TAIListForPaging* taiListForPaging_mp;

    //PagingPriority
    PagingPriority pagingPriority_m;
    PagingPriority *pagingPriority_mp;
    bool pagingPriorityPresent_m;

    //UERadioCapabilityForPaging
    UERadioCapabilityForPaging ueRadioCapabilityForPaging_m;
    unsigned char ueRadioCapabilityForPagingNr_ma[maxUeRadioCapabilityValue_c];
	unsigned char ueRadioCapabilityForPagingEutra_ma[maxUeRadioCapabilityValue_c];
    UERadioCapabilityForPaging* ueRadioCapabilityForPaging_mp;
    bool  ueRadioCapabilityForPagingPresent_m;

    //PagingOrigin
    PagingOrigin pagingOrigin_m;
    PagingOrigin *pagingOrigin_mp;
    bool pagingOriginPresent_m;

    //AssistanceDataForPaging
    AssistanceDataForPaging assDataForPaging_m;
    AssistanceDataForPaging* assDataForPaging_mp;
    AssDateForPaging AssDateForPaging_ma;
    bool assDataForPagingPresent_m;

    struct RecommendedCellList_ recommendedCellList_ma[maxnoofRecommendedCells_c];
    unsigned char nrCellmaId[maxnoofRecommendedCells_c];
    unsigned int numOfRecommendedCellItem_m;
};

#endif