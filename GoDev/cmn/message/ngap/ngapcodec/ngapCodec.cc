#include <iostream>

#include "ngapCodec.h"

NgapCodec::NgapCodec(){
	
	buf_m.length = 0;
	for(int i=0;i<sizeof(buf_m.value);i++){
		buf_m.value[i] = 0;
	}	
}

NgapCodec::~NgapCodec(){

}

void NgapCodec::setEncBuffer(uchar* buf, int size){
   if(buf == NULL){
	   std::cout << "invalid buf pointer"<< std::endl;
	   return;
   }

   for(int i = 0; i<size; i++){
      buf_m.value[i] = buf[i];
   }

   buf_m.length = size;

   return;   
}

EncBuffer* NgapCodec::getEncBuffer(){
	return &buf_m;
}

bool NgapCodec::decodeHeader(NgapOssCtxt* ctx){

	OssBuf ossbuf;
	ossbuf.length = buf_m.length;
	ossbuf.value = buf_m.value;
	NgapOssPdu ngappdu;
    ngappdu.setPerBufRef_v(ossbuf);
	ngappdu.decodeNgapPdu(ctx);

	msgType_m = ngappdu.getMsgType();
	procedureCode_m = ngappdu.getProcedureCode();

	return true;
}

int NgapCodec::getMsgType(){
   return msgType_m;
}

int NgapCodec::getProcedureCode(){
   return procedureCode_m;
}

OssCtxt::OssCtxt(){
	ossCtxtPtr_m = new NgapOssCtxt();
}

OssCtxt::~OssCtxt(){
	delete ossCtxtPtr_m;
	ossCtxtPtr_m = NULL;
}
