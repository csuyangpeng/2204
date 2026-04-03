#ifndef __ngapOssCtxt_h__INCLUDED__
#define __ngapOssCtxt_h__INCLUDED__

#include "ngapOss.h"

//////////////////////////////////////////////////////////////////////////
/// @class NgapOssCtxt --
///
/// This class provides the interfaces to the OSS Library for encoding and
/// decoding ngap messages
///
//////////////////////////////////////////////////////////////////////////
class NgapOssCtxt
{
   public:
      NgapOssCtxt();

      ////////////////////////////////////////////////////////////////////////
      /// @brief destructor
      ///
      /// @param - None
      ///
      /// @warning - None.
      ///
      /// @returns - None
      ///
      ////////////////////////////////////////////////////////////////////////
      ~NgapOssCtxt();
        
      ////////////////////////////////////////////////////////////////////////
      /// @brief instance
      ///
      /// This function gives access to the singleton object
      ///
      /// @param - None
      ///
      /// @warning - None.
      ///
      /// @returns - pointer to the singleton object.
      ///
      ////////////////////////////////////////////////////////////////////////
      static NgapOssCtxt* instance();
      
      ////////////////////////////////////////////////////////////////////////
      /// @brief instance
      ///
      /// This function resets the data members
      ///
      /// @param - None
      ///
      /// @warning - None.
      ///
      /// @returns - None.
      ///
      ////////////////////////////////////////////////////////////////////////      
      void reset_v();
    
      ////////////////////////////////////////////////////////////////////////
      /// @brief get the primary codec buffer
      ///
      /// This function is used to get the reference of the primaryBuffer_m
      ///
      /// @param - None
      ///
      /// @warning - None.
      ///
      /// @returns - reference to the OssBuf.
      ///
      ////////////////////////////////////////////////////////////////////////  
      OssBuf& getPrimaryCodecBuf();
      
      ////////////////////////////////////////////////////////////////////////
      /// @brief get the secondary codec buffer
      ///
      /// This function is used to get the reference of the secondaryBuffer_m
      ///
      /// @param - None
      ///
      /// @warning - None.
      ///
      /// @returns - reference to the OssBuf.
      ///
      //////////////////////////////////////////////////////////////////////// 
      OssBuf& getSecondaryCodecBuf();
    
      ////////////////////////////////////////////////////////////////////////
      /// @brief encode Oss to PER
      ///
      /// This function checks for the automatic or manual encoding and 
      /// subsequently calls the ossEncode.
      ///
      /// @param - pduNum - input - pdu number defined in NgApToed.h
      ///          ossInput_p - input - pointer to the NgAP_PDU structure
      ///          perOutput_p - input - pointer to where the encoded PER
      ///          data will be stored
      ///          automatic - input - automatic or manual encode : default 
      ///          is automatic
      ///
      /// @warning - None.
      ///
      /// @returns - return code 0 = success.
      ///
      //////////////////////////////////////////////////////////////////////// 
      unsigned int encodeOssToPer(int             pduNum,
                                  void*           ossInput_p, //NGAP_PDU*       ossInput_p, 
                                  OssBuf*         perOutput_p,
                                  bool            automatic = true);
			 
      ////////////////////////////////////////////////////////////////////////
      /// @brief decode PER to OSS
      ///
      /// This fucntion checks for the automastic or manual decoding and 
      /// subsequently calls the ossDecode.
      ///
      /// @param - pduNum - input - pdu number defined in NgApToed.h
      ///          ossInput_p - input - pointer to the NgAP_PDU structure
      ///          perOutput_p - input - pointer to where the decoded OSS
      ///          data will be stored
      ///          automatic - input - automatic or manual encode default is automatic
      ///
      /// @warning - None.
      ///
      /// @returns - return code 0 = success.
      ///
      //////////////////////////////////////////////////////////////////////// 	 
      unsigned int decodePerToOss(int*          pduNum_p,
                                  OssBuf*       perInput_p,
                                  OssBuf*       ossOutput_p, 
                                  bool          automatic = true);
				
   private :
    
      // OSS tool object
      OssGlobal automatic_m;
      OssGlobal manual_m ;
      
      // buffer used by OSS tool
      const static unsigned int ossTempBufferSize_c  = 64*1024;
      const static unsigned int ossCodecBufferSize_c = 16*1024;	 
      
      unsigned char primaryCodecBuffer_ma[ossCodecBufferSize_c];
      unsigned char secondaryCodecBuffer_ma[ossCodecBufferSize_c];
      unsigned char ossTempBuffer_ma[ossTempBufferSize_c];
      
      // The primaryBuffer_m is used for first round encoding/decoding.
      // For incremental decoding, the PER message is first decoded into the
      // primaryBuffer_m, then fields are decoded into the secondaryBuffer_m.
      OssBuf primaryBuffer_m;
      OssBuf secondaryBuffer_m;
      OssBuf ossTempBuffer_m;
      const static unsigned int manualEncodeNotSupported_c = 256;
      const static unsigned int successfulOperation_c = 0;
}; // end class NgapOssCtxt

#endif //__NgapOssCtxt_h_INCLUDED__
