#ifndef __ngApOssIe_h__INCLUDED__
#define __ngApOssIe_h__INCLUDED__

#include "ngApOssPdu.h"

class Ie
{
  public:

    ////////////////////////////////////////////////////////////////////////
    /// @brief Ie()
    ///
    /// constructor
    ///
    ////////////////////////////////////////////////////////////////////////
    Ie()
    {
    }

    ////////////////////////////////////////////////////////////////////////
    /// @brief ~Ie()
    ///
    /// destructor
    ///
    ////////////////////////////////////////////////////////////////////////
    virtual ~ Ie()
    {
    }

    ////////////////////////////////////////////////////////////////////////
    /// @brief reset_v
    ///
    /// This fucntion sets the primary and the secondary buffers.
    ///
    /// @param - None
    ///
    /// @warning - None.
    ///
    /// @returns - None
    ///
    ////////////////////////////////////////////////////////////////////////
    virtual void reset_v() = 0;

    ////////////////////////////////////////////////////////////////////////
    /// @brief encode IEs to the ProtocolIEs
    ///
    /// This is a blank implementation of the fucntion left for the derived 
    /// calsses to be defined as needed.
    ///
    /// @param - None.
    ///
    /// @warning - All mandatory IEs should be set before calling this method.
    ///
    /// @returns - Return Code - 0 = success.
    ///
    ////////////////////////////////////////////////////////////////////////
    virtual unsigned int encodeIe(ProtocolIE_Container encodeIe_p) = 0;

    ////////////////////////////////////////////////////////////////////////
    /// @brief get the data number for the decoded data
    ///
    /// This function sets the number to the decoded data.
    ///
    /// @param - None.
    ///
    /// @warning - None.
    ///
    /// @returns - None.
    ///
    ////////////////////////////////////////////////////////////////////////
    virtual void decodedIe_v(void *decodedValue_p) = 0;
};

#endif // __Ie_h__INCLUDED__
