
set(LIB_SRC ngapToed_x86.cc)

add_library(ngapToed STATIC ${LIB_SRC})

set(LIBRARY_OUTPUT_PATH ${PROJECT_BINARY_DIR}/lib)

set(OSSINC ${PROJECT_SOURCE_DIR}/include)
set(HEADER_PATH ${OSSINC})
include_directories(${HEADER_PATH})

set(CMAKE_CXX_FLAGS " -DOSSPRINT -DOSSDEBUG=2 -fexceptions -fno-merge-constants -fno-strict-aliasing -I${OSSINC} ")
set(CMAKE_BUILD_TYPE "Debug")
