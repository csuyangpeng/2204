
aux_source_directory(. DIR_LIB_SRCS)

add_library(ngapMsgTypes STATIC ${DIR_LIB_SRCS})

set(LIBRARY_OUTPUT_PATH ${PROJECT_BINARY_DIR}/lib)


set(OSSINC ${PROJECT_SOURCE_DIR}/include)
set(HEADER_PATH ${PROJECT_SOURCE_DIR}/inc
				${OSSINC})

include_directories(${HEADER_PATH})


set(CMAKE_CXX_FLAGS " -DOSSPRINT -DOSSDEBUG=2 -fexceptions -fno-merge-constants -fno-strict-aliasing")
set(CMAKE_BUILD_TYPE "Debug")
