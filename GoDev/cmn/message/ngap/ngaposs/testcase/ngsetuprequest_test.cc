#include <iostream>

#include <ngSetupRequestMsg.h>
#include <ngApOssPdu.h>


using namespace std;

unsigned char gNgSetupRequestMessage[] = {
0x00,0x15,0x00,0x48,0x00,0x00,0x04,0x00,0x66,0x00,0x26,0x01,0x00,0x01,0x02,0x03,0x10,0x01,0x02,0x03,0x00,0x00,0x00,0x08,0x01,0x02,0x03,0x00,0x00,0x00,0x08,0x01,
0x02,0x03,0x10,0x01,0x02,0x03,0x00,0x00,0x00,0x08,0x01,0x02,0x03,0x00,0x00,0x00,0x08,0x00,0x1B,0x00,0x08,0x00,0x01,0x02,0x03,0x10,0x11,0x22,0x11,0x00,0x52,0x40,
0x06,0x01,0x80,0x51,0x52,0x53,0x54,0x00,0x15,0x40,0x01,0x00};

int main()
{
    cout << "Start..." << endl;

    NGSetupRequestMsg ngsetupRequestMsg;

	unsigned char plmnValue_a[]={0x01, 0x02, 0x03};
	ngsetupRequestMsg.gGNBId_m.pLMNIdentity.length = sizeof(plmnValue_a);
	memcpy(ngsetupRequestMsg.gGNBId_m.pLMNIdentity.value,plmnValue_a,sizeof(plmnValue_a));
    unsigned char gnbidValue_a[]={0x11,0x22,0x11};
	ngsetupRequestMsg.gGNBId_m.gNB_ID.length = 24;//sizeof(gnbidValue_a);
	memcpy(ngsetupRequestMsg.gGNBId_m.gNB_ID.value, gnbidValue_a,sizeof(gnbidValue_a));

    ngsetupRequestMsg.pagingDrx_m = v32;

	ngsetupRequestMsg.gNbNamePresent_m = true;
	char rNameValue_a[] = {0x51, 0x52, 0x53, 0x54};
    ngsetupRequestMsg.gNbName_m.nameLen = sizeof(rNameValue_a);
	memcpy(ngsetupRequestMsg.gNbName_m.value, rNameValue_a, sizeof(rNameValue_a));

	ngsetupRequestMsg.numofSupportedTAItems_m = 2;
	ngsetupRequestMsg.numofBPlmnItem_m[0] = 2;
	ngsetupRequestMsg.numofBPlmnItem_m[1] = 2;
	ngsetupRequestMsg.numofSliceItem_m[0][0]=1;
	ngsetupRequestMsg.numofSliceItem_m[0][1]=1;
	ngsetupRequestMsg.numofSliceItem_m[1][0]=1;
	ngsetupRequestMsg.numofSliceItem_m[1][1]=1;

    unsigned char tacValue_a[] = {0x01, 0x02,0x03}; 
	//unsigned char plmnValue_a[]={0x01, 0x02, 0x03};
	unsigned char sDValue_a[]={0x01, 0x02, 0x03};
	unsigned char sSTValue_a[]={0x01};

    TAC tacValue;
	tacValue.length = sizeof(tacValue_a);
	memcpy(tacValue.value, tacValue_a, sizeof(tacValue_a));

	PLMNIdentity plmnValue;
	plmnValue.length = sizeof(plmnValue_a);
	memcpy(plmnValue.value, plmnValue_a, sizeof(plmnValue_a));
	
	SD sdValue;
	sdValue.length = sizeof(sDValue_a);
	memcpy(sdValue.value, sDValue_a, sizeof(sDValue_a));

	SST sstValue;
	sstValue.length = sizeof(sSTValue_a);
	memcpy(sstValue.value, sSTValue_a, sizeof(sSTValue_a));

    for(int i=0; i< ngsetupRequestMsg.numofSupportedTAItems_m ; i++)
    {
        ngsetupRequestMsg.tacList_m[i] = tacValue;
		for(int j=0; j< ngsetupRequestMsg.numofBPlmnItem_m[i]; j++)
		{
		    ngsetupRequestMsg.plmnList_m[i][j] = plmnValue;
			for(int k=0;k< ngsetupRequestMsg.numofSliceItem_m[i][j];k++)
			{
			   ngsetupRequestMsg.nssaiSDList_m[i][j][k] = sdValue;
			   ngsetupRequestMsg.nssaiSSTList_m[i][j][k] = sstValue;
			}
		}
    }

	cout << "Initial message data"<<endl;
	ngsetupRequestMsg.display();


    NgApOssCtxt  ctx_data, *ctx = &ctx_data;
	

	MsgBuffer buffer;
	buffer = ngsetupRequestMsg.encode(ctx);
	
    cout << "encode, length = " << buffer.length << endl;

    cout << "Start decode the message"<< endl;
//	NGSetupRequestMsg ngsetupRequestMsgDecode;
//	MsgBuffer decBuffer;
//	decBuffer.length = sizeof(gNgSetupRequestMessage);
//	decBuffer.value= gNgSetupRequestMessage;
//    int rt = ngsetupRequestMsgDecode.decodePerToOss(ctx,decBuffer);
//	ngsetupRequestMsgDecode.decode();
//    ngsetupRequestMsg.display();


	MsgBuffer decBuffer;
	decBuffer.length = sizeof(gNgSetupRequestMessage);
	decBuffer.value= gNgSetupRequestMessage;

	OssBuf ossbuf;
	ossbuf.length = decBuffer.length;
	ossbuf.value = decBuffer.value;
	NgApOssPdu ngappdu;
    ngappdu.setPerBufRef_v(ossbuf);
	ngappdu.decodeNgApPdu(ctx);

	cout <<"Msg Type: " << ngappdu.getMsgType() << endl;
	cout <<"Msg Procedure code: "<< ngappdu.getProcedureCode() << endl;

	NGSetupRequestMsg ngsetupRequestMsgDec;
	ngsetupRequestMsgDec.decodePerToOss(ctx,decBuffer);
	ngsetupRequestMsgDec.decode();
	ngsetupRequestMsgDec.display();
	
    cout << "Finished."<< endl;
	return 0;
}
