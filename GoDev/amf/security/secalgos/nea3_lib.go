/* -----------------------------------------------------------------------
 * code extracted from ETSI / SAGE specification of the 3GPP Confidentiality and Integrity Algorithms 128-EEA3 & 128-EIA3.
 * Document 2: ZUC Specification. Version 1.6 from 28th June 2011, appendix A.
 * https://www.gsma.com/security/wp-content/uploads/2019/05/eea3eia3zucv16.pdf
 * code extracted from ETSI / SAGE specification of the 3GPP Confidentiality and Integrity Algorithms 128-EEA3 & 128-EIA3.
 * Document 1: 128-EEA3 and 128-EIA3 Specification. Version 1.7 from the 30th December 2011, annex 1.
 * https://www.gsma.com/security/wp-content/uploads/2019/05/EEA3_EIA3_specification_v1_8.pdf
 * (warning: only link to version 1.9 exists)
 *
 * All updated ZUC specifications maybe found on the GSMA website:
 * https://www.gsma.com/security/security-algorithms/
 *-----------------------------------------------------------------------*/

/*---------------------------------------------
 * ZUC / EEA3 / EIA3 : LTE security algorithm
 *--------------------------------------------*/
// EEA3
package secalgos

/*
#include <stdio.h>
#include <stdlib.h>


typedef unsigned char u8;
typedef unsigned int u32;

void myprint(char* s) {
	printf("%s\n", s);
	printf("%ld\n",sizeof(u32));
}

// the state registers of LFSR
u32 LFSR_S0;
u32 LFSR_S1;
u32 LFSR_S2;
u32 LFSR_S3;
u32 LFSR_S4;
u32 LFSR_S5;
u32 LFSR_S6;
u32 LFSR_S7;
u32 LFSR_S8;
u32 LFSR_S9;
u32 LFSR_S10;
u32 LFSR_S11;
u32 LFSR_S12;
u32 LFSR_S13;
u32 LFSR_S14;
u32 LFSR_S15;
//the registers of F
u32 F_R1;
u32 F_R2;

//the outputs of BitReorganization
u32 BRC_X0;
u32 BRC_X1;
u32 BRC_X2;
u32 BRC_X3;

// the s-boxes
u8 S0[256] = {
0x3e,0x72,0x5b,0x47,0xca,0xe0,0x00,0x33,0x04,0xd1,0x54,0x98,0x09,0xb9,0x6d,0xcb,
0x7b,0x1b,0xf9,0x32,0xaf,0x9d,0x6a,0xa5,0xb8,0x2d,0xfc,0x1d,0x08,0x53,0x03,0x90,
0x4d,0x4e,0x84,0x99,0xe4,0xce,0xd9,0x91,0xdd,0xb6,0x85,0x48,0x8b,0x29,0x6e,0xac,
0xcd,0xc1,0xf8,0x1e,0x73,0x43,0x69,0xc6,0xb5,0xbd,0xfd,0x39,0x63,0x20,0xd4,0x38,
0x76,0x7d,0xb2,0xa7,0xcf,0xed,0x57,0xc5,0xf3,0x2c,0xbb,0x14,0x21,0x06,0x55,0x9b,
0xe3,0xef,0x5e,0x31,0x4f,0x7f,0x5a,0xa4,0x0d,0x82,0x51,0x49,0x5f,0xba,0x58,0x1c,
0x4a,0x16,0xd5,0x17,0xa8,0x92,0x24,0x1f,0x8c,0xff,0xd8,0xae,0x2e,0x01,0xd3,0xad,
0x3b,0x4b,0xda,0x46,0xeb,0xc9,0xde,0x9a,0x8f,0x87,0xd7,0x3a,0x80,0x6f,0x2f,0xc8,
0xb1,0xb4,0x37,0xf7,0x0a,0x22,0x13,0x28,0x7c,0xcc,0x3c,0x89,0xc7,0xc3,0x96,0x56,
0x07,0xbf,0x7e,0xf0,0x0b,0x2b,0x97,0x52,0x35,0x41,0x79,0x61,0xa6,0x4c,0x10,0xfe,
0xbc,0x26,0x95,0x88,0x8a,0xb0,0xa3,0xfb,0xc0,0x18,0x94,0xf2,0xe1,0xe5,0xe9,0x5d,
0xd0,0xdc,0x11,0x66,0x64,0x5c,0xec,0x59,0x42,0x75,0x12,0xf5,0x74,0x9c,0xaa,0x23,
0x0e,0x86,0xab,0xbe,0x2a,0x02,0xe7,0x67,0xe6,0x44,0xa2,0x6c,0xc2,0x93,0x9f,0xf1,
0xf6,0xfa,0x36,0xd2,0x50,0x68,0x9e,0x62,0x71,0x15,0x3d,0xd6,0x40,0xc4,0xe2,0x0f,
0x8e,0x83,0x77,0x6b,0x25,0x05,0x3f,0x0c,0x30,0xea,0x70,0xb7,0xa1,0xe8,0xa9,0x65,
0x8d,0x27,0x1a,0xdb,0x81,0xb3,0xa0,0xf4,0x45,0x7a,0x19,0xdf,0xee,0x78,0x34,0x60
};

u8 S1[256] =  {
0x55,0xc2,0x63,0x71,0x3b,0xc8,0x47,0x86,0x9f,0x3c,0xda,0x5b,0x29,0xaa,0xfd,0x77,
0x8c,0xc5,0x94,0x0c,0xa6,0x1a,0x13,0x00,0xe3,0xa8,0x16,0x72,0x40,0xf9,0xf8,0x42,
0x44,0x26,0x68,0x96,0x81,0xd9,0x45,0x3e,0x10,0x76,0xc6,0xa7,0x8b,0x39,0x43,0xe1,
0x3a,0xb5,0x56,0x2a,0xc0,0x6d,0xb3,0x05,0x22,0x66,0xbf,0xdc,0x0b,0xfa,0x62,0x48,
0xdd,0x20,0x11,0x06,0x36,0xc9,0xc1,0xcf,0xf6,0x27,0x52,0xbb,0x69,0xf5,0xd4,0x87,
0x7f,0x84,0x4c,0xd2,0x9c,0x57,0xa4,0xbc,0x4f,0x9a,0xdf,0xfe,0xd6,0x8d,0x7a,0xeb,
0x2b,0x53,0xd8,0x5c,0xa1,0x14,0x17,0xfb,0x23,0xd5,0x7d,0x30,0x67,0x73,0x08,0x09,
0xee,0xb7,0x70,0x3f,0x61,0xb2,0x19,0x8e,0x4e,0xe5,0x4b,0x93,0x8f,0x5d,0xdb,0xa9,
0xad,0xf1,0xae,0x2e,0xcb,0x0d,0xfc,0xf4,0x2d,0x46,0x6e,0x1d,0x97,0xe8,0xd1,0xe9,
0x4d,0x37,0xa5,0x75,0x5e,0x83,0x9e,0xab,0x82,0x9d,0xb9,0x1c,0xe0,0xcd,0x49,0x89,
0x01,0xb6,0xbd,0x58,0x24,0xa2,0x5f,0x38,0x78,0x99,0x15,0x90,0x50,0xb8,0x95,0xe4,
0xd0,0x91,0xc7,0xce,0xed,0x0f,0xb4,0x6f,0xa0,0xcc,0xf0,0x02,0x4a,0x79,0xc3,0xde,
0xa3,0xef,0xea,0x51,0xe6,0x6b,0x18,0xec,0x1b,0x2c,0x80,0xf7,0x74,0xe7,0xff,0x21,
0x5a,0x6a,0x54,0x1e,0x41,0x31,0x92,0x35,0xc4,0x33,0x07,0x0a,0xba,0x7e,0x0e,0x34,
0x88,0xb1,0x98,0x7c,0xf3,0x3d,0x60,0x6c,0x7b,0xca,0xd3,0x1f,0x32,0x65,0x04,0x28,
0x64,0xbe,0x85,0x9b,0x2f,0x59,0x8a,0xd7,0xb0,0x25,0xac,0xaf,0x12,0x03,0xe2,0xf2
};
//the constants D
u32 EK_d[16] = {
0x44D7, 0x26BC, 0x626B, 0x135E, 0x5789, 0x35E2, 0x7135, 0x09AF,
0x4D78, 0x2F13, 0x6BC4, 0x1AF1, 0x5E26, 0x3C4D, 0x789A, 0x47AC
};
// c = a + b mod (2^31 - E1)
u32 AddM(u32 a, u32 b)
{
	u32 c = a + b;
	return (c & 0x7FFFFFFF) + (c >> 31);
}

//LFSR with initialization mode
#define MulByPow2(x, k) ((((x) << k) | ((x) >> (31 - k))) & 0x7FFFFFFF)

void LFSRWithInitialisationMode(u32 u)
{
	u32 f, v;
	f = LFSR_S0;

	v = MulByPow2(LFSR_S0, 8);
	f = AddM(f, v);
	v = MulByPow2(LFSR_S4, 20);
	f = AddM(f, v);
	v = MulByPow2(LFSR_S10, 21);
	f = AddM(f, v);
	v = MulByPow2(LFSR_S13, 17);
	f = AddM(f, v);
	v = MulByPow2(LFSR_S15, 15);
	f = AddM(f, v);

	f = AddM(f, u);

    LFSR_S0 = LFSR_S1;
	LFSR_S1 = LFSR_S2;
	LFSR_S2 = LFSR_S3;
	LFSR_S3 = LFSR_S4;
	LFSR_S4 = LFSR_S5;
	LFSR_S5 = LFSR_S6;
	LFSR_S6 = LFSR_S7;
	LFSR_S7 = LFSR_S8;
	LFSR_S8 = LFSR_S9;
	LFSR_S9 = LFSR_S10;
	LFSR_S10 = LFSR_S11;
	LFSR_S11 = LFSR_S12;
	LFSR_S12 = LFSR_S13;
	LFSR_S13 = LFSR_S14;
	LFSR_S14 = LFSR_S15;
	LFSR_S15 = f;
}

void LFSRWithWorkMode(void)
{
    u32 f, v;
	f = LFSR_S0;

	v = MulByPow2(LFSR_S0, 8);
	f = AddM(f, v);
	v = MulByPow2(LFSR_S4, 20);
	f = AddM(f, v);
	v = MulByPow2(LFSR_S10, 21);
	f = AddM(f, v);
	v = MulByPow2(LFSR_S13, 17);
	f = AddM(f, v);
	v = MulByPow2(LFSR_S15, 15);
	f = AddM(f, v);

    LFSR_S0 = LFSR_S1;
	LFSR_S1 = LFSR_S2;
	LFSR_S2 = LFSR_S3;
	LFSR_S3 = LFSR_S4;
	LFSR_S4 = LFSR_S5;
	LFSR_S5 = LFSR_S6;
	LFSR_S6 = LFSR_S7;
	LFSR_S7 = LFSR_S8;
	LFSR_S8 = LFSR_S9;
	LFSR_S9 = LFSR_S10;
	LFSR_S10 = LFSR_S11;
	LFSR_S11 = LFSR_S12;
	LFSR_S12 = LFSR_S13;
	LFSR_S13 = LFSR_S14;
	LFSR_S14 = LFSR_S15;
	LFSR_S15 = f;
}

void BitReorganization(void)
{
    BRC_X0 = ((LFSR_S15 & 0x7FFF8000) << 1) | (LFSR_S14 & 0xFFFF);
	BRC_X1 = ((LFSR_S11 & 0xFFFF) << 16) | (LFSR_S9 >> 15);
	BRC_X2 = ((LFSR_S7 & 0xFFFF) << 16) | (LFSR_S5 >> 15);
	BRC_X3 = ((LFSR_S2 & 0xFFFF) << 16) | (LFSR_S0 >> 15);
}

#define ROT(a, k) (((a) << k) | ((a) >> (32 - k)))

u32 L1(u32 X)
{
    return (X ^ ROT(X, 2) ^ ROT(X, 10) ^ ROT(X, 18) ^ ROT(X, 24));
}

u32 L2(u32 X)
{
    return (X ^ ROT(X, 8) ^ ROT(X, 14) ^ ROT(X, 22) ^ ROT(X, 30));
}

#define MAKEU32(a, b, c, d) (((u32)(a) << 24) | ((u32)(b) << 16) | ((u32)(c) << 8) | ((u32)(d)))

u32 F(void)
{
    u32 W, W1, W2, u, v;

    W  = (BRC_X0 ^ F_R1) + F_R2;
    W1 = F_R1 + BRC_X1;
    W2 = F_R2 ^ BRC_X2;

    u = L1((W1 << 16) | (W2 >> 16));
    v = L2((W2 << 16) | (W1 >> 16));

    F_R1 = MAKEU32(S0[u >> 24], S1[(u >> 16) & 0xFF],
    S0[(u >> 8) & 0xFF], S1[u & 0xFF]);
    F_R2 = MAKEU32(S0[v >> 24], S1[(v >> 16) & 0xFF],
    S0[(v >> 8) & 0xFF], S1[v & 0xFF]);

    return W;
}

#define MAKEU31(a, b, c) (((u32)(a) << 23) | ((u32)(b) << 8) | (u32)(c))

 void Initialization(u8* k, u8* iv)
{
    u32 w, nCount;

    LFSR_S0 = MAKEU31(k[0], EK_d[0], iv[0]);
    LFSR_S1 = MAKEU31(k[1], EK_d[1], iv[1]);
    LFSR_S2 = MAKEU31(k[2], EK_d[2], iv[2]);
    LFSR_S3 = MAKEU31(k[3], EK_d[3], iv[3]);
    LFSR_S4 = MAKEU31(k[4], EK_d[4], iv[4]);
    LFSR_S5 = MAKEU31(k[5], EK_d[5], iv[5]);
    LFSR_S6 = MAKEU31(k[6], EK_d[6], iv[6]);
    LFSR_S7 = MAKEU31(k[7], EK_d[7], iv[7]);
    LFSR_S8 = MAKEU31(k[8], EK_d[8], iv[8]);
    LFSR_S9 = MAKEU31(k[9], EK_d[9], iv[9]);
    LFSR_S10 = MAKEU31(k[10], EK_d[10], iv[10]);
    LFSR_S11 = MAKEU31(k[11], EK_d[11], iv[11]);
    LFSR_S12 = MAKEU31(k[12], EK_d[12], iv[12]);
    LFSR_S13 = MAKEU31(k[13], EK_d[13], iv[13]);
    LFSR_S14 = MAKEU31(k[14], EK_d[14], iv[14]);
    LFSR_S15 = MAKEU31(k[15], EK_d[15], iv[15]);

    F_R1 = 0;
    F_R2 = 0;
    nCount = 32;
    while (nCount > 0)
    {
        BitReorganization();
        w = F();
        LFSRWithInitialisationMode(w >> 1);
        nCount --;
    }
}

 void GenerateKeystream(u32* pKeystream, u32 KeystreamLen)
{
    u32 i;
    BitReorganization();
    F();
    LFSRWithWorkMode();

    for (i = 0; i < KeystreamLen; i ++)
    {
        BitReorganization();
        pKeystream[i] = F() ^ BRC_X3;
        LFSRWithWorkMode();
    }
}

void ZUC(u8* k, u8* iv, u32* ks, u32 len)
{
    Initialization(k, iv);
    GenerateKeystream(ks, len);
}

 // EEA3

void EEA3(u8* CK, u32 COUNT, u32 BEARER, u32 DIRECTION,
u32 LENGTH, u32* M, u32* C)
{
    u32 *z, L, i;
    u8 	IV[16];
    u32 lastbits = (32-(LENGTH%32))%32;

    L 	= (LENGTH+31)/32;
    z 	= (u32 *) malloc(L*sizeof(u32));

    IV[0]	= (COUNT>>24) & 0xFF;
    IV[1]	= (COUNT>>16) & 0xFF;
    IV[2]	= (COUNT>>8)  & 0xFF;
    IV[3]	=  COUNT      & 0xFF;

    IV[4]	= ((BEARER << 3) | ((DIRECTION&1)<<2)) & 0xFC;
    IV[5]	= 0;
    IV[6]	= 0;
    IV[7]	= 0;

    IV[8]	= IV[0];
    IV[9]	= IV[1];
    IV[10]	= IV[2];
    IV[11]	= IV[3];

    IV[12]	= IV[4];
    IV[13]	= IV[5];
    IV[14]	= IV[6];
    IV[15]	= IV[7];
  //printf("---k:%x\n", CK[0]);
  //printf("---k:%x\n", CK[1]);
  //printf("---k:%x\n", CK[2]);
  //printf("---k:%x\n", CK[3]);
  //printf("---k:%x\n", CK[4]);
  //printf("---count:%x\n", COUNT);
  //printf("---BEARER:%x\n", BEARER);
  //printf("---DIRECTION:%x\n", DIRECTION);
  //printf("---LENGTH:%d\n", LENGTH);
  //
  //printf("---L:%d\n", L);
  //
  //
  //printf("---M:---\n");
  //printf("---M:%08x\n", M[0]);
  //printf("---M:%08x\n", M[1]);
  //printf("---M:%08x\n", M[2]);
  //printf("---M:%08x\n", M[3]);
  //printf("---M:%08x\n", M[4]);
  //printf("---M:%08x\n", M[5]);
  //printf("---M:%08x\n", M[6]);

    ZUC(CK, IV, z, L);

    //printf("C:");
    for (i=0; i<L; i++)
        {
            C[i] = M[i] ^ z[i];
            //printf("%x ", C[i]);

         }
        //printf("\n");
    // zero last bits of data in case its length is not word-aligned (32 bits)
    //   this is an addition to the C reference code, which did not handle it
    if (lastbits)
    i--;
    C[i] &= 0x100000000 - (u32)(1<<lastbits);

    //printf("%d\n", C[i]);

    free(z);
}


 // EIA3

//EIA3: LTE Integrity computation algorithm
//EIA3.c

u32 GET_WORD(u32 * DATA, u32 i)
{
    u32 WORD, ti;
    ti	= i % 32;
    if (ti == 0)
    WORD = DATA[i/32];
    else
    WORD = (DATA[i/32]<<ti) | (DATA[i/32+1]>>(32-ti));
    return WORD;
}

u8 GET_BIT(u32 * DATA, u32 i)
{
    return (DATA[i/32] & (1<<(31-(i%32)))) ? 1 : 0;
}

void EIA3(u8* IK, u32 COUNT, u32 BEARER, u32 DIRECTION,
u32 LENGTH, u32* M, u32* MAC)
{
    u32	*z, N, L, T, i;
    u8 IV[16];

    IV[0]	= (COUNT>>24) & 0xFF;
    IV[1]	= (COUNT>>16) & 0xFF;
    IV[2]	= (COUNT>>8) & 0xFF;
    IV[3]	= COUNT & 0xFF;

    IV[4]	= (BEARER << 3) & 0xF8;
    IV[5]	= IV[6] = IV[7] = 0;

    IV[8]	= ((COUNT>>24) & 0xFF) ^ ((DIRECTION&1)<<7);
    IV[9]	= (COUNT>>16) & 0xFF;
    IV[10]	= (COUNT>>8) & 0xFF;
    IV[11]	= COUNT & 0xFF;

    IV[12]	= IV[4];
    IV[13]	= IV[5];
    IV[14]	= IV[6] ^ ((DIRECTION&1)<<7);
    IV[15]	= IV[7];

    N	= LENGTH + 64;
    L	= (N + 31) / 32;
    z	= (u32 *) malloc(L*sizeof(u32));
    ZUC(IK, IV, z, L);

    T = 0;
    for (i=0; i<LENGTH; i++) {
        if (GET_BIT(M,i)) {
        T ^= GET_WORD(z,i);
        }
    }
    T ^= GET_WORD(z,LENGTH);

    *MAC = T ^ z[L-1];
    free(z);
}
// end of EIA3.c
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"unsafe"
)

func test1() {
	//使用C.CString创建的字符串需要手动释放。
	//cs := C.CString("Hello World\n")
	//fmt.Println("call C.myprint start")
	//C.myprint(cs) // 启用一个新的线程
	//C.free(unsafe.Pointer(cs))
	//fmt.Println("call C.myprint end")
	//	output
	/*call C.myprint start
	call C.myprint end
	Hello World*/

	fmt.Println(C.LFSR_S0)
	//void Initialization(u8* k, u8* iv);
	//k := []byte{100}
	//iv := []byte{0}
	//var ck *C.u8
	//var civ *C.u8
	//
	//ck = (*C.uchar)(unsafe.Pointer(&k[0]))
	//civ = (*C.uchar)(unsafe.Pointer(&iv[0]))

	//C.Initialization(ck, civ)
	//C.free(unsafe.Pointer(ck))
	//C.free(unsafe.Pointer(civ))
	fmt.Println(C.LFSR_S0)

	//void EEA3(
	// u8* CK, u32 COUNT,
	// u32 BEARER, u32 DIRECTION,
	// u32 LENGTH, u32* M, u32* C)

	// 4.4Test Set 2
	Key, _ := hex.DecodeString("e5bd3ea0eb55ade866c6ac58bd54302a")
	Count := uint32(0x56823)
	Bearer := uint32(0x18)
	Direction := uint32(0x1)
	Length := 800 //bits,Plaintext length
	// 有字节序问题
	//Plaintext, _ := hex.DecodeString("6cf65340735552ab0c9752fa6f9025fe0bd675d9005875b200000000")
	var Plaintext = []uint32{0x14a8ef69, 0x3d678507, 0xbbe7270a, 0x7f67ff50, 0x06c3525b, 0x9807e467, 0xc4e56000, 0xba338f5d,
		0x42955903, 0x67518222, 0x46c80d3b, 0x38f07f4b, 0xe2d8ff58, 0x05f51322, 0x29bde93b,
		0xbbdcaf38, 0x2bf1ee97, 0x2fbf9977, 0xbada8945, 0x847a2a6c, 0x9ad34a66, 0x7554e04d, 0x1f7fa2c3, 0x3241bd8f, 0x01ba220d}
	//CiphertextT, _ := hex.DecodeString("a6c85fc66afb8533aafc2518dfe784940ee1e4b030238cc800000000")
	Ciphertext := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	fmt.Println("Ciphertext len ", len(Ciphertext))

	CK := (*C.uchar)(unsafe.Pointer(&Key[0]))
	COUNT := (C.uint)(Count)
	BEARER := (C.uint)(Bearer)
	DIRECTION := (C.uint)(Direction)
	LENGTH := (C.uint)(Length)
	M := (*C.uint)(unsafe.Pointer(&Plaintext[0]))
	Cm := (*C.uint)(unsafe.Pointer(&Ciphertext[0]))
	fmt.Println("CK ", CK)
	fmt.Println("COUNT ", COUNT)
	fmt.Println("BEARER ", BEARER)
	fmt.Println("DIRECTION ", DIRECTION)
	fmt.Println("LENGTH ", LENGTH)
	fmt.Println("M ", *M)
	fmt.Println("Cm ", *Cm)

	//func C.GoBytes(unsafe.Pointer, C.int) []byte
	goM := C.GoBytes((unsafe.Pointer)(M), C.int(len(Plaintext)))
	fmt.Printf("M %x\n", goM)
	goCm := C.GoBytes((unsafe.Pointer)(Cm), C.int(len(Ciphertext)))
	fmt.Printf("M %x\n", goCm)
	C.EEA3(CK, COUNT, BEARER, DIRECTION, LENGTH, M, Cm)
	goCm2 := C.GoBytes((unsafe.Pointer)(Cm), C.int(len(Ciphertext)))
	fmt.Printf("-----加密 %x\n", goCm2)
	fmt.Printf("-----加密 %x\n", BytesToUint32Array(goCm2))
	testC, _ := hex.DecodeString("e0431d135cbea1de97fd1b5abf2c851d4f7b2d71ea1f9657a8af0832f433a4bcc709ad56bc587e416" +
		"688cf69743f35d178805e86fb2d201dfcf7cf3e0f193bbc4e202ae8fc50e3d013266f0fa6bcf2b23a475adf0da0a457d8ba5e9838f2d680017ba064")

	if !bytes.Equal(goCm2, testC) {
		panic("加密错误!")
	}
	// 解密
	text := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	textC := (*C.uint)(unsafe.Pointer(&text[0]))
	C.EEA3(CK, COUNT, BEARER, DIRECTION, LENGTH, Cm, textC)
	goCm3 := C.GoBytes((unsafe.Pointer)(textC), C.int(len(Ciphertext)))
	fmt.Printf("-----解密 %x\n", goCm3)
	fmt.Printf("-----解密 %x\n", BytesToUint32Array(goCm3))
	testM, _ := hex.DecodeString("69efa8140785673d0a27e7bb50ff677f5b52c30667e407980060e5c45d8f33ba03599542228251673b" +
		"0dc8464b7ff03858ffd8e22213f5053be9bd2938afdcbb97eef12b7799bf2f4589daba6c2a7a84664ad39a4de05475c3a27f1f8fbd41320d22ba01")

	if !bytes.Equal(goCm3, testM) {
		panic("解密错误!")
	}

}

//C 131d43e0 dea1be5c 5a1bfd97 1d852cbf 712d7b4f 57961fea 3208afa8 bca433f4 56ad09c7 417e58bc 69cf8866 d1353f74
//  865e8078 1d202dfb 3ecff7fc bc3b190f e82a204e d0e350fc 0f6f2613 b2f2bca6 df5a473a 57a4a00d 985ebad8 80d6f238 64a07b01
func BytesToUint32Array(m []byte) []uint32 {
	dst := make([]uint32, 0)

	// 4byte to uint32
	var length = len(m)
	var offset = 0
	for {

		if length >= 4 {
			// m 大端序，最少4个字节，BigEndian 指[]byte的字节序
			dst = append(dst, binary.BigEndian.Uint32(m[offset:]))
			length -= 4
			offset += 4
		}
		if length < 4 { // 4 byte alignment
			break
		}
	}
	return dst
}

func BytesToUint32ArrayV1(m []byte) ([]uint32, error) {
	dst := make([]uint32, 0)
	// m 大端序，4个字节对齐
	if len(m)%4 != 0 {
		err := fmt.Errorf("Invalid length")
		return nil, err
	}
	// 4byte to uint32
	var length = len(m)
	var offset = 0
	for {

		if length >= 4 {
			// m 大端序，最少4个字节
			dst = append(dst, binary.BigEndian.Uint32(m[offset:]))
			length -= 4
			offset += 4
		}
		if length < 4 { // 4 byte alignment
			break
		}
	}
	return dst, nil
}
func Uint32ArrayToBytesV1(m []uint32) ([]byte, error) {
	dst := make([]byte, 0)
	// m 大端序，4个字节对齐
	// uint32 to 4byte
	b := make([]byte, 4)
	for _, value := range m {
		binary.BigEndian.PutUint32(b, value)
		dst = append(dst, b...)
	}
	if len(dst) == 0 {
		err := fmt.Errorf("Invalid length")
		return nil, err
	}
	return dst, nil
}

// 32bits左对齐，右侧位补充0
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize) % blockSize
	padtext := bytes.Repeat([]byte{0}, padding) //用0去填充
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

// 统一使用大端字节序
// void EEA3(u8* CK, u32 COUNT, u32 BEARER, u32 DIRECTION,
// u32 LENGTH, u32* M, u32* C)
func Eea3(ck []byte, count uint32, bearer uint32, direction uint32,
	length uint32, m []uint32, cipher []byte) {

	CK := (*C.uchar)(unsafe.Pointer(&ck[0]))
	COUNT := (C.uint)(count)
	BEARER := (C.uint)(bearer)
	DIRECTION := (C.uint)(direction)
	LENGTH := (C.uint)(length)
	M := (*C.uint)(unsafe.Pointer(&m[0]))
	CiphertextC := (*C.uint)(unsafe.Pointer(&cipher[0]))
	C.EEA3(CK, COUNT, BEARER, DIRECTION, LENGTH, M, CiphertextC)

}
func Eea3V1(ck []byte, count uint32, bearer uint32, direction uint32,
	length uint32, m []uint32, ciper []uint32) {
	//	入参检查，切片不能为空

	CK := (*C.uchar)(unsafe.Pointer(&ck[0]))
	COUNT := (C.uint)(count)
	BEARER := (C.uint)(bearer)
	DIRECTION := (C.uint)(direction)
	LENGTH := (C.uint)(length)
	M := (*C.uint)(unsafe.Pointer(&m[0]))
	CiphertextC := (*C.uint)(unsafe.Pointer(&ciper[0]))
	C.EEA3(CK, COUNT, BEARER, DIRECTION, LENGTH, M, CiphertextC)
	//goCm3 := C.GoBytes((unsafe.Pointer)(CiphertextC), C.int(len(ciper)))
	//ciphertext = goCm3
	return
}

// 统一使用大端字节序
//void EIA3(u8* IK, u32 COUNT, u32 BEARER, u32 DIRECTION,
//				   u32 LENGTH, u32* M, u32* MAC)
func Eia3(ik []byte, count uint32, bearer uint32, direction uint32,
	length uint32, m []uint32, mac []byte) {

	CK := (*C.uchar)(unsafe.Pointer(&ik[0]))
	COUNT := (C.uint)(count)
	BEARER := (C.uint)(bearer)
	DIRECTION := (C.uint)(direction)
	LENGTH := (C.uint)(length)
	M := (*C.uint)(unsafe.Pointer(&m[0]))
	// 小端模式
	MAC := (*C.uint)(unsafe.Pointer(&mac[0]))
	C.EIA3(CK, COUNT, BEARER, DIRECTION, LENGTH, M, MAC)

}

// 仅返回值MAC格式不同
// length:bits,Plaintext length
func Eia3V1(ik []byte, count uint32, bearer uint32, direction uint32,
	length uint32, m []uint32, mac *uint32) {

	CK := (*C.uchar)(unsafe.Pointer(&ik[0]))
	COUNT := (C.uint)(count)
	BEARER := (C.uint)(bearer)
	DIRECTION := (C.uint)(direction)
	LENGTH := (C.uint)(length)
	M := (*C.uint)(unsafe.Pointer(&m[0]))
	// 小端模式
	MAC := (*C.uint)(unsafe.Pointer(mac))
	C.EIA3(CK, COUNT, BEARER, DIRECTION, LENGTH, M, MAC)

}
