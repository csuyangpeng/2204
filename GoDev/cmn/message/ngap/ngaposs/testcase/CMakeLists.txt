
set(OSSLIB "/data/lib/ossasn1/linux-glibc2.3-amd64/10.1.1/lib")
set(LIBPATH ${PROJECT_BINARY_DIR}/lib
            ${OSSLIB})

link_directories(${LIBPATH})

aux_source_directory(. DIR_SRCS)

set(EXECUTABLE_OUTPUT_PATH ${PROJECT_BINARY_DIR}/bin)

add_executable(SetupReqTest ${DIR_SRCS})

set(OSSINC "/data/lib/ossasn1/linux-glibc2.3-amd64/10.1.1/include")
set(HEADER_PATH ${PROJECT_SOURCE_DIR}/inc
				${OSSINC})
include_directories(${HEADER_PATH})

set(CMAKE_CXX_FLAGS " -DOSSPRINT -DOSSDEBUG=2 -fexceptions -fno-merge-constants -fno-strict-aliasing -I${OSSINC} -I{SRCINC} ")
set(CMAKE_BUILD_TYPE "Debug")
                                            
target_link_libraries(SetupReqTest ngApMessage ngApOssPdus ngApToed ngApMsgTypes libosstoed.a)
