#ifndef __ngSetupRequestPdu_h__INCLUDED__
#define __ngSetupRequestPdu_h__INCLUDED__

#include "ngapOssPdu.h"

///////////////////////////////////////////////////////////////////////////////
// @class NGSetupRequestPdu
//-- **************************************************************
//--
//-- NG SETUP REQUEST
//--
//-- **************************************************************
//
//NGSetupRequest ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {NGSetupRequestIEs} },
//	...
//}
//
// NGSetupRequestIEs NGAP-PROTOCOL-IES ::= {
// { ID id-GlobalRANNodeID			CRITICALITY reject	TYPE GlobalRANNodeID	PRESENCE mandatory	}|
// { ID id-RANNodeName				CRITICALITY ignore	TYPE RANNodeName		PRESENCE optional		}|
// { ID id-SupportedTAList			CRITICALITY reject	TYPE SupportedTAList	PRESENCE mandatory	}|
// { ID id-DefaultPagingDRX		CRITICALITY ignore	TYPE PagingDRX			PRESENCE mandatory	}|
// { ID id-UERetentionInformation	CRITICALITY ignore	TYPE UERetentionInformation 	PRESENCE optional	},
// ...
//}
///////////////////////////////////////////////////////////////////////////////

class NGSetupRequestPdu:public NgapOssPdu
{
  public:
    NGSetupRequestPdu();
    //virtual
    ~NGSetupRequestPdu();
    
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
    /// @brief get NGSetupRequest pointer
    /// returns the pointer to the
    /// NGSetupRequest in the decoded message
    /////////////////////////////////////////////////////////////
    NGSetupRequest * getNgSetupRequestDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// Global Ran Node Id
    /////////////////////////////////////////////////////////////
    GlobalRANNodeID * getGlobalRanNodeIdDecodedPtr() const;
    GlobalRANNodeID * getGlobalRanNodeIdEncodePtr();

    /////////////////////////////////////////////////////////////
    /// RanNodeName 
    /////////////////////////////////////////////////////////////
    RANNodeName * getRanNodeNameDecodedPtr() const;
    RANNodeName * getRanNodeNameEncodePtr();
    bool isRanNodeNamePresent() const;

    /////////////////////////////////////////////////////////////
    /// Supported TA List
    /////////////////////////////////////////////////////////////
    SupportedTAList * getSupportedTAListDecodedPtr() const;
    SupportedTAList * getSupportedTAListEncodePtr();

    SupportedTAItem * getSupportedTAListItemDecodedPtr(unsigned int index) const;
    SupportedTAItem * getSupportedTAListItemEncodePtr(unsigned int index);

    void setNumOfSupportedTAListItem_v(unsigned int num);
    unsigned int getNumOfSupportedTAListItem() const;


    BroadcastPLMNItem* getBPlmnsDecodedPtr(SupportedTAItem * taiPtr,  
		                                        unsigned int indBPLMNs) const;
    BroadcastPLMNItem* getBPlmnsEncodePtr(unsigned int indTACs, 
		                                       unsigned int indBPLMNs);
    void setNumOfBPlmns_v(unsigned int num, unsigned int index);
    unsigned int getNumOfBPlmns(unsigned int index) const;

    SliceSupportItem* getSliceSupportListDecodedPtr(BroadcastPLMNItem* bplmns_p,
		                                                     unsigned int indSlices) const;
    SliceSupportItem* getSliceSupportListEncodePtr(unsigned int indTACs, 
                                                  unsigned int indBPLMNs, 
                                                  unsigned int indSlices);
    void setNumOfSlices_v(unsigned int num, unsigned int indexTac, unsigned int indexPlmns);
    unsigned int getNumOfSlices(unsigned int indexTac, unsigned int indexPlmns) const;

    /////////////////////////////////////////////////////////////
    /// PagingDRX
    /////////////////////////////////////////////////////////////
    PagingDRX * getPagingDrxDecodedPtr() const;
    PagingDRX * getPagingDrxEncodePtr();

    /////////////////////////////////////////////////////////////
    /// Ue Retetion Information
    /////////////////////////////////////////////////////////////
    UERetentionInformation * getUeRetentionInfoDecodedPtr() const;
    UERetentionInformation * getUeRetentionInfoEncodePtr();
    bool isUeRetentionInfoPresent() const;

  private:
  
    typedef NGSetupRequestIe ProtocolIeType;

    //OSS structures
    NGSetupRequest *NGSetupRequest_mp; //for encode

    //Global ENB ID
    GlobalRANNodeID globalRanNodeId_m; //for encode
    GlobalRANNodeID *globalRanNodeId_mp; //for decode
	unsigned char gnbIdValue_ma[maxGnbIdValue_c];

    //ENB Name
    RANNodeName ranNodeName_m;
    char ranNodeNameValue_ma[maxRanNameValue_c];
    bool ranNodeNamePresent_m;

    RANNodeName *ranNodeName_mp; 

    //Supported TAs
    SupportedTAList supportedTAList_m;
    struct SupportedTAList_ supportedTAList_ma[maxnoofTACs_c];
	unsigned int numOfSupportedTAListItem_m;
    SupportedTAList *supportedTAList_mp; 
 
    //BroadcastPLMNList
    unsigned int numofBroadcastPLMNItem_m[maxnoofTACs_c];
    struct BroadcastPLMNList_ broadcastPLMNItem_ma[maxnoofTACs_c][maxnoofBPLMNs_c];

    //SliceSupportList
    unsigned int numOfSliceSupportItem_m[maxnoofTACs_c][maxnoofBPLMNs_c];
    struct SliceSupportList_ sliceSupportItem_ma[maxnoofTACs_c][maxnoofBPLMNs_c][maxnoofSlices_c];

    //Paging DRX
    PagingDRX pagingDrx_m;   //for encode
    PagingDRX *pagingDrx_mp; //for decode

    //UE Retention Information
    UERetentionInformation  UeRetentionInfo_m;
    UERetentionInformation* UeRetentionInfo_mp;
    bool  UeRetentionInfoPresent_m;
};

#endif // __ngSetupRequestPdu_h__INCLUDED__
