/*****************************************************************************/
/* Copyright (C) 2015 OSS Nokalva, Inc.  All rights reserved.                */
/*****************************************************************************/

/* THIS FILE IS PROPRIETARY MATERIAL OF OSS NOKALVA, INC.
 * AND MAY BE USED ONLY BY DIRECT LICENSEES OF OSS NOKALVA, INC.
 * THIS FILE MAY NOT BE DISTRIBUTED.
 * THIS COPYRIGHT STATEMENT MAY NOT BE REMOVED. */

/*****************************************************************************/
/* FILE: @(#)osscag.h	17.1  13/12/19                       */
/*****************************************************************************/

#ifndef OSSCAG_H
#define OSSCAG_H

#define getmem runtime_getmem

#include "ossiaapi.h"

#undef getmem

/*****************************************************************************/
/*  OSS CAGL API return codes                                                */
/*****************************************************************************/

    /*************************************************************************/
    /*  ossCreateControlTable() error codes                                  */
    /*************************************************************************/
#define ASN1_COMPILER_DLL_NOT_LINKED 104 /* ASN.1 compiler DLL was not linked */
#define ASN1_COMPILING_FAILED        105 /* compiling of the input ASN.1
					  * notation failed */
#define TOO_MANY_CMD_LINE_OPTIONS    106 /* number of input command line options
					  * exceeds OSS internal limit */
    /*************************************************************************/
    /*  One CAGL DLL-specific error codes                                    */
    /*************************************************************************/

#define FAILED_TO_UNLOCK_DLL_8       116
#define FAILED_TO_UNLOCK_DLL_7       117
#define FAILED_TO_UNLOCK_DLL_6       118
#define FAILED_TO_UNLOCK_DLL_5       119
#define FAILED_TO_UNLOCK_DLL_4       120
#define FAILED_TO_UNLOCK_DLL_3       121
#define FAILED_TO_UNLOCK_DLL_2       122
#define FAILED_TO_UNLOCK_DLL_1       123
#define FAILED_TO_ACCESS_DLL_DATA    124
#define INVALID_DLL_CALLER_2         125
#define INVALID_DLL_CALLER_1         126

#define ERROR_READING_SIGNATURE_FILE 127
#define MISSING_FILE_NAME_IN_SIGNATURE_FILE_OPTION 128
#define OSS_CAGL_DLL_IS_NOT_UNLOCKED 129

/* Additional more specific than ASN1_COMPILING_FAILED codes that are
 * determined based on compiling errors. */
#define ASN1_COMPILING_LICENSE_ERROR      131
#define ASN1_COMPILING_INPUT_OUTPUT_ERROR 132

#ifdef __cplusplus
extern "C" {
#endif

/*****************************************************************************/
/*  CAGL type definitions                                                    */
/*****************************************************************************/

#ifdef _WIN32
typedef struct osscagTbl {
    int (DLL_ENTRY_FPTR *_System ossCreateControlTablep)(struct ossAsn1Global *,
						char *, char *, void **);
    void (DLL_ENTRY_FPTR *_System ossDeleteControlTablep)(
						struct ossAsn1Global *, void *);
} OssCagTbl;

/* The structure "cagWinParm" is used to store DLL-related information. */
typedef struct cagWinParm {
    OssCagTbl  *osscagTbl;      /* Runtime compiler DLL function table */
    HINSTANCE   hOsscagDLL;     /* Handle of runtime compiler DLL */
} CagWinParm;

PUBLIC void DLL_ENTRY ossLinkCAGL(struct ossAsn1Global *aworld);
#endif /* _WIN32 */

typedef void *MemInfoHolder[2 + 4*10];



typedef struct ossAsn1Global {
    void	*ctlTbl;
    void	*(CDECL_ENTRY_FPTR *mallocp)(size_t p);
    void	*(CDECL_ENTRY_FPTR *reallocp)(void *p, size_t s);
    void	(CDECL_ENTRY_FPTR *freep)(void *p);
    unsigned long parmFlags;
    unsigned int  isExternal:1,
		  was_canceled:1,
		  mi_reset_in_main:1;
    int           specialErrorCode;
    MemInfoHolder mih;
	/* used for storing DLL- & library NLMs-related parameters */
    char	*cmd_line;
#ifdef _WIN32
    void	*bdata;
    CagWinParm	wp;
#endif /* _WIN32 */
} OssAsn1Global;

/*****************************************************************************/
/*  OSS CAGL API functions                                                   */
/*****************************************************************************/

#if defined(_WIN32)
/* Support for one CAGL DLL protection */
PUBLIC int DLL_ENTRY_FDEF ossUnlockDLL(struct ossAsn1Global *aworld,
					void *blobData);
#endif /* _WIN32 */

/* Support for ASN.1 compiling at runtime */
PUBLIC struct ossAsn1Global *DLL_ENTRY ossGetAsn1CompilerHandle(
				void *(CDECL_ENTRY_FPTR *userMalloc)(size_t size),
				void  (CDECL_ENTRY_FPTR *userFree)(void *buf),
				void *(CDECL_ENTRY_FPTR *userRealloc)(void *buf, size_t size));
PUBLIC void DLL_ENTRY ossDeleteAsn1CompilerHandle(struct ossAsn1Global *aworld);
PUBLIC int DLL_ENTRY  ossCreateControlTable(struct ossAsn1Global *aworld,
				char *cmd_line, char *err_file,
				void **ctbl_hdl);
PUBLIC void DLL_ENTRY ossDeleteControlTable(struct ossAsn1Global *aworld,
                                void *ctbl_hdl);

#ifdef __cplusplus
}
#endif /* __cplusplus */

#endif /* OSSCAG_H */
