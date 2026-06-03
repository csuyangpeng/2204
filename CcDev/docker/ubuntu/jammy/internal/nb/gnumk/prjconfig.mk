#real comiling parameter setup, next var must be setup: 
g++:=g++
gcc:=gcc
#compiling flags setup:
CXXFLAGS :=-g -O2 -Wall  -Wno-error -fPIC -pthread
CCFLAGS:=-g -O2 -Wall  -Wno-error -fPIC  -pthread

#linking flags setup:
LINKFLAGS:=  -L$(LIBDIR)
#compiling includ header file path setup:
IPATH:=-I. 
#lib linking  path setup:
LLPATH:=-L$(LIBDIR)  
#Program linking path setup:
PLPATH:=-L$(LIBDIR)



