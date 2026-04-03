#ifndef __ngSetupResponsePdu_h__INCLUDED__
#define __ngSetupResponsePdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class NGSetupResponsePdu
//-- **************************************************************
//--
//-- NG SETUP RESPONSE
//--
//-- **************************************************************
//
//NGSetupResponse ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {NGSetupResponseIEs} },
//	...
//}
//
//NGSetupResponseIEs NGAP-PROTOCOL-IES ::= {
//	{ ID id-AMFName						CRITICALITY reject	TYPE AMFName					PRESENCE mandatory	}|
//	{ ID id-ServedGUAMIList				CRITICALITY reject	TYPE ServedGUAMIList			PRESENCE mandatory	}|
//	{ ID id-RelativeAMFCapacity			CRITICALITY ignore	TYPE RelativeAMFCapacity		PRESENCE mandatory	}|
//	{ ID id-PLMNSupportList				CRITICALITY reject	TYPE PLMNSupportList			PRESENCE mandatory	}|
//	{ ID id-CriticalityDiagnostics		CRITICALITY ignore	TYPE CriticalityDiagnostics		PRESENCE optional	}|
//  { ID id-UERetentionInformation	    CRITICALITY ignore	TYPE UERetentionInformation 	PRESENCE optional	},
//	...
//}
///////////////////////////////////////////////////////////////////////////////

class NGSetupResponsePdu:public NgapOssPdu
{
  public:
    NGSetupResponsePdu();
    //virtual
    ~NGSetupResponsePdu();
    
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
    /// @brief get NGSetupResponse pointer
    /// returns the pointer to the
    /// NGSetupResponse in the decoded message
    /////////////////////////////////////////////////////////////
    NGSetupResponse * getNGSetupResponseDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// AmfName 
    /////////////////////////////////////////////////////////////
    AMFName* getAmfNameDecodedPtr() const;
    AMFName* getAmfNameEncodePtr();

    /////////////////////////////////////////////////////////////
    /// ServedGUAMIList 
    /////////////////////////////////////////////////////////////
    ServedGUAMIList* getServedGUAMIListDecodedPtr() const;
    ServedGUAMIList* getServedGUAMIListEncodePtr();

	ServedGUAMIItem* getServedGUAMIListItemDecodedPtr(unsigned int index) const;
	ServedGUAMIItem* getServedGUAMIListItemEncodedPtr(unsigned int index);

	void setNumofServedGUAMIListItem_v(unsigned int num);
	unsigned int getNumofServedGUAMIListItem() const;
	
    /////////////////////////////////////////////////////////////
    /// PLMNSupport List List
    /////////////////////////////////////////////////////////////
    PLMNSupportList * getPlmnSupportListDecodedPtr() const;
    PLMNSupportList * getPlmnSupportListEncodePtr();

    PLMNSupportItem * getPlmnSupportItemDecodedPtr(unsigned int index) const;
    PLMNSupportItem * getPlmnSupportItemEncodePtr(unsigned int index);

    void setNumofPlmnSupportListItem_m_v(unsigned int num);
    unsigned int getNumofPlmnSupportListItem() const;


	SliceSupportItem* getSliceSupportListDecodedPtr(PLMNSupportItem* plmnSupportItemPtr,
													unsigned int indSlices) const;
    SliceSupportItem* getSliceSupportListEncodePtr(         unsigned int indPlmn, 
                                                  unsigned int indSlices);

	void setNumofSlices_v(unsigned int num, unsigned int indPlmn);
	unsigned int getNumofSlices(unsigned int indPlmn) const;

    /////////////////////////////////////////////////////////////
    /// RelativeAMFCapacity
    /////////////////////////////////////////////////////////////
    RelativeAMFCapacity* getRelAmfCapacityDecodedPtr() const;
    RelativeAMFCapacity* getRelAmfCapacityEncodePtr();

    /////////////////////////////////////////////////////////////
    /// Ue Retetion Information
    /////////////////////////////////////////////////////////////
    UERetentionInformation * getUeRetentionInfoDecodedPtr() const;
	UERetentionInformation * getUeRetentionInfoEncodePtr();
	bool isUeRetentionInfoPresent() const;

  private:
  
    typedef NGSetupResponseIe ProtocolIeType;

    //OSS structures
    NGSetupResponse *NGSetupResponse_mp; //for encode

    //AMF Name
    AMFName amfName_m;  //for encode
    char amfNameValue_ma[maxAMFNameValue_c];
	char backupAmfNameValue_ma[maxAMFNameValue_c];
	
    AMFName *amfName_mp;  //for decode

	//Served GUAMI List
    ServedGUAMIList	servedGUAMIList_m;
	struct ServedGUAMIList_ servedGUAMIList_ma[maxnoofServedGUAMIs_c];
	unsigned int numOfSrvGUAMIListItem_m;
	AmfIdentifier guamiAmfIdentifier_ma[maxnoofServedGUAMIs_c];

	ServedGUAMIList	*servedGUAMIList_mp; //for decode	

	//Relative AMF Capacity
	RelativeAMFCapacity relAmfCapacity_m;
	RelativeAMFCapacity* relAmfCapacity_mp;

    //Plmn Support List
    PLMNSupportList plmnSupportList_m;
	struct PLMNSupportList_ plmnSupportList_ma[maxnoofPLMNs_c];
	unsigned int numOfPlmnSupportListItem_m;

    //SliceSupportList
    unsigned int numOfSliceSupportItem_m[maxnoofPLMNs_c];
    struct SliceSupportList_ sliceSupportItem_ma[maxnoofBPLMNs_c][maxnoofSlices_c];

	PLMNSupportList *plmnSupportList_mp; //for decode

	//UE Retention Information
	UERetentionInformation  UeRetentionInfo_m;
	UERetentionInformation* UeRetentionInfo_mp;
	bool  UeRetentionInfoPresent_m;
};

#endif // __ngSetupResponsePdu_h__INCLUDED__

