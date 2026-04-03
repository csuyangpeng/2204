#ifndef __NgapCodec_H__
#define __NgapCodec_H__

#include <ngapOssCtxt.h>
#include <ngapOssPdu.h>

typedef unsigned char uchar;

typedef struct EncBuffer {
    int     length;
    uchar   value[1024];
} EncBuffer;

class OssCtxt 
{
public:
    OssCtxt();
    ~OssCtxt();
    NgapOssCtxt* ossCtxtPtr_m;
};

class NgapCodec
{
public:
    NgapCodec();
    ~NgapCodec();

    void setEncBuffer(uchar* buf, int size);
    EncBuffer* getEncBuffer();
    
    bool decodeHeader(NgapOssCtxt* ctx);
    int getMsgType();
    int getProcedureCode();

private:

    int msgType_m;
    int procedureCode_m;
    EncBuffer buf_m;
};

#endif 