define  COMPILe_TIME_CMD
  sBeginTime=$(shell date "+%s"); \
  BeginTime=$(shell date "+%H:%M:%S"); \
  echo INFO:Begin compile time $${BeginTime}; \
  EndTime=`date +%H:%M:%S`;\
  echo "INFO:End  compile  time $${EndTime}"; \
  sEndTime=`date +%s` ;\
   timespan=`echo $${sEndTime}-$${sBeginTime}|bc`; \
  thetime=`echo $${timespan}/3600|bc`:`echo $${timespan}%3600/60|bc`:`echo $${timespan}%60|bc`; \
   echo "INFO:Compiling  time is $${thetime}";
endef


define create_common_directory
  directory:=$(OUTPUTDIR) $(ARCHDIR) $(LIBDIR) $(BINDIR) $(SRCOBJDIR)
  $$(shell mkdir -p $$(directory))
  $$(shell echo "make directory success: $$(directory)")
endef

#-----------------FUNCTION-----initialization_for_build:setup  varaiable for building------------------------------------------
define   initialization_for_build
  #common variable:
  other_goals:= help  
  OBJECT_FILE_EXT=.o
  BUILD_CONFIG_FILE_NAME:=make.mk

   #prject name setup: must be the  pathern of $(PRJNAME)/src 
   ifeq ("$$(words $$(subst /src, /src,$$(abspath .)))","3")   
      #path include two "src" word
      PRJNAME:=$$(notdir $$(word 2,$$(subst /src, /src,$$(abspath .))))
   else 
       #path include one  "src" word
       PRJNAME:=$$(notdir $$(subst /src,,$$(filter %src,$$(subst src,src ,$$(abspath .)))))
   endif
   $$(warning PRJNAME=$$(PRJNAME))   
   #some common path setup:
    TOPDIR:=$$(firstword $$(subst /$$(PRJNAME)/src, ,$$(abspath ./)))
    #project's  source path must be the style of:  xxx/$(PRJNAME)/src 
    SRCDIR:=$$(TOPDIR)/$$(PRJNAME)/src
    OUTPUTDIR:=$$(TOPDIR)/$$(PRJNAME)/Dos

   ifeq  ("$$(PRJNAME)", "")
      ifneq ("$$(MAKECMDGOALS)","help") 
           $$(error you must running nb command  in path of style: $$$$(PRJNAME)/src!) 
      endif
   endif

   #support c/c++ source type list:
   C_PLUS_PLUS_TYPES:=%.cc %.cpp %.cxx
   C_TYPES:=%.c

   #common compiling output path:
   LIBDIR:=$$(OUTPUTDIR)/lib
   BINDIR:=$$(OUTPUTDIR)/bin
   SRCOBJDIR:=$$(OUTPUTDIR)/obj
    
    #create common directory:
    $$(eval $$(call create_common_directory))
endef

#*****************************************************************************************************************
#include  prjconfig.mk file:
#if prjconfig.mk file is  in project source path ,use it, if no, using default prjconfig.mk file
define load_prjconfig
   #include project config(for compiler and its  parameter setup) file:
   PRJCONFIG:=$$(wildcard $$(SRCDIR)/prjconfig.mk)
   $$(warning PRJCONFIG=$$(PRJCONFIG))
   ifneq ("$$(MAKECMDGOALS)","clean")
   ifneq ("$$(MAKECMDGOALS)","subclean")
       ifneq ("$$(wildcard $$(SRCDIR)/prjconfig.mk)","")
          #use project src path config file:
          include $$(SRCDIR)/prjconfig.mk
       else
          #use default prject config file:
          include  $$(dir $$(MAKEFILE))/prjconfig.mk
          PRJCONFIG:=$$(dir $$(MAKEFILE))/prjconfig.mk
       endif
   endif
   endif 
endef
#****************************************************************************************************************


#**************************************Create  library  linking  rules function***************************************
#--$1: lib name,
#--$2: dependent lib name.
#--$3: Object file  path.
#--$4: src file path.
#--$5: rules.mk file full name
define create_lib_rule
  #step-1:Get source file  full name:
  #get lib's  source file list(full name): 
  lsources:=$$(wildcard $$(addprefix  $4,$$($1_SRCS)))
  lsources:=$$(filter  $(C_PLUS_PLUS_TYPES) $(C_TYPES), $$(lsources))
  #convert relative path to absolute path:
  lsources:=$$(foreach src,$$(lsources),$$(abspath $$(dir $$(src))/$$(notdir $$(src))))


  
  #step-2:Get lib's dependent file(object file:*.o) from source file list
  #get *.o file list from src file(*.cc,*.c, *.cpp ,and so on) 
  depobjs:=$$(subst $(SRCDIR),$(SRCOBJDIR),$$(addsuffix $(OBJECT_FILE_EXT),$$(basename $$(lsources))))
  #step-3:adding dependent lib file:
   deplibs:=$$(filter $(allgenlibs),$2)
   deplibs_fullname:=$$(addprefix $(LIBDIR)/lib,$$(addsuffix .so,$$(deplibs)))
   depfiles:=$$(depobjs)  $$(deplibs_fullname)
   
   #attention:goal full name not including version info:
   goal:=$(LIBDIR)/lib$1.so
   #linking  library falg:-shared
   cmdline:=$(g++) -shared
   #attention:linking output goal including version: 
   lib_version:=$($1_VERSION)
   cmdline+=-o  $$(goal)$$(if $$(lib_version),.$$(lib_version),) $$(depobjs) 
   #adding linking need lib name:
   cmdline+=$$(if $2,$$(addprefix -l,$2),) $(LLPATH) $(LINKFLAGS)

   #step-3:save rule to file:
   $$(shell echo "#linking $(LIBDIR)/lib$1.so$$(if $$(lib_version),.$$(lib_version),)  library  rule:" >> $5)
   $$(shell echo "$$(goal):$$(depfiles)"                         >> $5)
   $$(shell echo "	@echo Linking library lib$1.so$$(if $$(lib_version),.$$(lib_version),)  from $4 " >> $5)
   $$(shell echo "	cd $4 && $$(cmdline) "      >> $5)
   ifneq ("$$(lib_version)","")
      $$(shell echo "	cd $(LIBDIR) &&  ln -sf lib$1.so.$$(lib_version)  lib$1.so && echo" >> $5)
   else
      $$(shell echo "	@echo  "          >> $5)
   endif
   $$(shell echo " "                      >> $5)
   
   #step4:create source file compiling rules file::
    $$(foreach src,$$(lsources),$$(eval $$(call create_obj_rule,$4,$3,$$(src),$1,$5)))
endef 

#**************************************compiling  source to object  file  rules function***************************************
#--$1: src file path:()
#--$2: obj file path
#--$3: src file full  name (including absolute path),source file may not in the same path in $1, it may be in ralative path in $1!
#--$4: program name or  lib name (according it to get source compiling flags)
#--$5: rules.mk file full name
define create_obj_rule
  #rule goal and dependent setup:
   goal:=$$(addsuffix $$(OBJECT_FILE_EXT),$$(subst  $$(SRCDIR),$$(SRCOBJDIR),$$(basename $3)))
  #not  create the same rules:
  ifeq ("$$(filter $$(goal),$$(make.mk.objs))","")
      make.mk.objs+=$$(goal)
      #select compiler and compiling flags, according src file suffix(.cc,*.cpp/.c)      
      cmdline:=$$(if $$(filter $$(C_PLUS_PLUS_TYPES),$3),$$(g++) $$(CXXFLAGS),$$(gcc) $$(CCFLAGS))
     
      #adding current path to  compiling search header file path:
      cmdline+=-I$1
      
      #compiling FLAGS  for C or C++ file:
      cmdline+=$$($4_CXXFLAGS) $$($4_CCFLAGS) $$($4_CFLAGS) $$($4_CPPFLAGS)
     
      #adding current src path for  header file search :
      cmdline+=$$(IPATH)
      
      #adding compiling obj file and src file:
      #source file using non-absolutly path:
      cmdline+=-o $$(goal) -c $$(subst $$(TOPDIR),".",$3)
     
      #object file's path not equal to obj file path, it may be other subdirectory,need to create:
      ifneq ("$$(dir $$(goal))","$2")
          cmdline:=mkdir -p $$(dir $$(goal))  && cd $$(TOPDIR)  && $$(cmdline)
      else
          cmdline:=cd $$(TOPDIR)  && $$(cmdline)         
      endif

      $$(shell echo "#source file compiling to *.o file rule:"                    >> $5)
      $$(shell echo "$$(goal):$3"                                                  >> $5)
      $$(shell echo "	@echo Compiling source file $$(notdir $3) in $$(dir $3)"    >> $5)
      $$(shell echo "	$$(cmdline) "      >> $5)
      $$(shell echo "   allobjs+=$$(goal)"                                         >> $5)
      $$(shell echo "         "                                                    >> $5)
  endif
endef

#**************************************linking   progrm file  rules function***************************************
#--$1: program name,
#--$2: program dependent lib name.
#--$3: Object file  path.
#--$4: src file path.
#--$5: rules.mk file full name
define create_program_rule
  #step-1:Get source file  full name:
  #get lib's  source file list(full name):    
  psources:=$$(wildcard $$(addprefix  $4,$$($1_SRCS)))
  psources:=$$(filter  $(C_PLUS_PLUS_TYPES) $(C_TYPES), $$(psources))
  #convert relative path to absolute path:
  psources:=$$(foreach src,$$(psources),$$(abspath $$(dir $$(src))/$$(notdir $$(src))))
  
  #step-2:Get lib's dependent file(object file:*.o) from source file list
  #get *.o file list from src file(*.cc,*.c, *.cpp ,and so on) 
  depobjs:=$$(subst $(SRCDIR),$(SRCOBJDIR),$$(addsuffix $(OBJECT_FILE_EXT),$$(basename $$(psources))))

   #get  dependent file list: *.Io(converting from *.cpp, *.cc,*.c,and so on.)
   #step3:adding dependent lib file:
   deplibs:=$$(filter $$(allgenlibs),$2)
   deplibs_fullname:=$$(addprefix $(LIBDIR)/lib,$$(addsuffix .so,$$(deplibs)))
   depfiles:=$$(depobjs) $$(deplibs_fullname)
   goal:=$(BINDIR)/$1

   #step4: lib rule command line create:
   cmdline:=$(g++) -o $$(goal) $$(depobjs) 
   #adding linking dep. lib name:
   cmdline+=$(ASNEEDED)
   llibs:=$$(if $2,$2,$$($1_LIBS))
   cmdline+=$$(addprefix -l,$$(llibs))
   #adding linking flags:
   cmdline+=$(PLPATH) $(LINKFLAGS) $($1_LINKFLAGS)

   #step5:save rule to file:
   $$(shell echo "Program linking  rule:"                               >> $5)
   $$(shell echo "$$(goal):$$(depfiles)"                                >> $5)
   $$(shell echo "	@echo Linking program $1 from $4"               >> $5)
   $$(shell echo "	cd $4 && $$(cmdline) && echo "                 >> $5)
   $$(shell echo " "                                                    >> $5)
   #step6--create source file compiling rules:
   $$(foreach src,$$(psources),$$(eval $$(call create_obj_rule,$4,$3,$$(src),$1,$5)))
endef

#-$1--make.mk file full name
#--function: filter not using build config file and sum buinding objects:
define precreate_rules
  PROGRAMS:=
  LIBRARIES:=
  SUBDIRS:=
  include $1

  #Create Filter make.mk list:
  compilingdirs:=$$(if $$(SUBDIRS),$$(SUBDIRS),$$(allcards_SUBDIRS))
  cur_subdirs:=$$(filter-out %svn,$$(notdir $$(shell find $$(dir  $1)  -maxdepth 1 -type d)))
  filter_dirs:=$$(filter-out  $$(compilingdirs),$$(cur_subdirs))

  FILTER_BUILD_CONFIG_FILE:=$$(if $$(filter_dirs),$$(addprefix  $$(dir $1),$$(addsuffix /%,$$(filter_dirs))),)
  #repeative value not adding:
  FILTER_BUILD_CONFIG_FILES+=$$(if $$(filter  $$(FILTER_BUILD_CONFIG_FILES),$$(FILTER_BUILD_CONFIG_FILE)),,$$(FILTER_BUILD_CONFIG_FILE))

ifeq ("$$(filter $$(FILTER_BUILD_CONFIG_FILES),$1)","")
  libs:=$$(LIBRARIES)
  #lib simple name:
  allgenlibs+=$$(libs)
  #lib full names:
  alllibs+=$$(addprefix $$(LIBDIR)/lib,$$(addsuffix .so,$$(libs)))
  programs:=$$(PROGRAMS)
  #program fullname:
  allprograms+=$$(addprefix $$(BINDIR)/,$$(programs))

  #allobjs not sum here:because the same source file can be compiling to defferent path, so sum in create_obj_rule file
endif
endef

#-$1--make.mk file full name
#--this function create temp rules_setup.mk file so as to create real rulse ,include *.cc/*.c compiling ,*.so linking ,program linking rules,and so on.
define create_rule_file
   #create rule template file:
   srcdir:=$(dir $1)
   #curdir:this var used by make.mk file calling $(wildcard ...) function: 
   curdir:=$(dir $1)
   objdir:=$(SRCOBJDIR)$(subst $(SRCDIR),,$(dir $1))
    
   #now create rules_setup.mk  file:
   $$(shell mkdir -p $$(objdir))
   rules.mk:=$$(abspath $$(objdir)/rules.mk)
   rules_setup.mk:=$$(abspath $$(objdir)/rules_setup.mk)
   
   $$(shell echo "#this file is create by makefile, not to edit it!" > $$(rules.mk))
   PROGRAMS:=
   LIBRARIES:=
   include $1
   libs:=$$(LIBRARIES)
   $$(shell echo "   include $1"   >$$(rules_setup.mk))
   ifneq ("$$(libs)","")
         $$(shell echo "   #create lib compiling and  linking rules:"   >>$$(rules_setup.mk))
         $$(shell echo '   libs:=$$(libs)' >>$$(rules_setup.mk))
         $$(shell echo '   $$$$(foreach lib,$$(strip $$(libs)),$$$$(eval $$$$(call create_lib_rule,$$$$(lib),$$$$($$$$(lib)_LIBS),$$(objdir),$$(srcdir),$$(rules.mk))))'  >>$$(rules_setup.mk))
   endif
   $$(shell echo "              "   >>$$(rules_setup.mk))
   programs:=$$(PROGRAMS)
   ifneq ("$$(programs)","")
         $$(shell echo "   #create program compiling and linking rules:"   >>$$(rules_setup.mk))
         $$(shell echo '   programs:=$$(programs)' >>$$(rules_setup.mk))
         $$(shell echo '   $$$$(foreach program,$$(strip $$(programs)),$$$$(eval $$$$(call create_program_rule,$$$$(program),$$$$($$$$(program)_LIBS),$$(objdir),$$(srcdir),$$(rules.mk))))'  >>$$(rules_setup.mk))
   endif
   include $$(rules_setup.mk)
endef


#---this function no input  parameter
define setup_rule_variable
  allobjs:=
  allprograms:=
  allrules:=
  export allgenlibs:=
  make.mk.objs:=
  rules_setup.mks:=
  #curent source path directory:
  export  curdir:=
  FILTER_BUILD_CONFIG_FILES:=
  alllibs:=
  
  ifeq ("$$(abspath .)","$$(SRCDIR)")
      #Running in $$(PRJNAME)/src directory,so get subpath from top BUILD CONFIG FILE
     compiling_info:=now running  compiling command  from $$(SRCDIR) path.
      include $$(SRCDIR)/$$(BUILD_CONFIG_FILE_NAME)
      #compiling_path is a subpath list, including $$(SRCDIR) path:
      compiling_path:=$$(addprefix $$(SRCDIR)/,$$(SUBDIRS)) $$(SRCDIR)
    else
      compiling_info:=now running compiling command  in  a  subdirectory:
      compiling_path:=$$(abspath .)
    endif
 endef

#---this function no input  parameter
define create_common_rule 
 .PHONY:  help clean subclean 
  help:
	@echo "nb              :Builing using default  prjconfig.mk file."
	@echo "nb help         :Display make help infomation."    
	@echo "nb clean        :Clean project output path:$$(PRJNAME)/output/$$(ARCH)."    
	@echo "nb subclean     :Clean  project src's subdirectory compiling output directory."    
  clean:
	rm -rf $$(OUTPUTDIR)
	@echo
 #subdirectory clean:only delete project subdirectory compiling files:
  subclean:
	rm -rf $$(SRCOBJDIR)$$(subst $$(SRCDIR),,$$(abspath .))/
endef


define main_fun
.PHONY:all  after_all
after_all:
   $$(eval $$(call initialization_for_build))
   ifeq ("$$(filter $$(MAKECMDGOALS),clean subclean help)","")
       $$(eval $$(call load_prjconfig))
   endif
   $$(eval $$(call create_common_rule))
      $$(eval $$(call setup_rule_variable))
      #get build config  file list:
      BUILD_CONFIG_FILES:=$$(sort $$(shell find $$(compiling_path)  -name '$$(BUILD_CONFIG_FILE_NAME)'))
      #Prepare for create rules:
      $$(foreach mk,$$(BUILD_CONFIG_FILES),$$(eval $$(call precreate_rules,$$(mk))))
      #filter not used build config  file:
      BUILD_CONFIG_FILES:=$$(filter-out $$(FILTER_BUILD_CONFIG_FILES),$$(BUILD_CONFIG_FILES))

      #get rules.mk file set:
      rules.mks:=$$(addsuffix rules.mk,$$(subst $$(SRCDIR),$$(SRCOBJDIR),$$(dir $$(BUILD_CONFIG_FILES))))
 
     #create rules.mk file:
      $$(foreach file,$$(BUILD_CONFIG_FILES),$$(eval $$(call create_rule_file,$$(file)))) 
    
      #load rules from file:
      include $$(rules.mks) 
      #nb command running in SRCDIR, so  making package after compiling:
      ifeq ("$$(abspath .)","$$(SRCDIR)")
         all:$$(allprograms) $$(alllibs) 
      else
      #nb command running in subpath of $(SRCDIR)
         all:$$(allprograms) $$(alllibs) $$(allobjs)
      endif

      after_all:all
	@echo INFO:linking  $$(words $$(allprograms)) program files.
	@echo INOF:linking  $$(words $$(alllibs)) libraries.
	@echo INFO:compiling $$(words $$(allobjs)) source files.	  
	@echo INFO:nb building sucess! && echo    
	@$$(COMPILe_TIME_CMD)
     
      
       
      #follow rules will used to make all source files compiling first,and then lib linking, and program file linking last!
      #after all lib linking finished, begin to link program file: 
      $$(allprograms):$$(alllibs)
      #after all source file compiling finished,begin to link lib file:
      $$(alllibs):$$(allobjs) 
   
      #if prjconfig.mk is updated, all source file must be re-compile:
      $$(allobjs):$$(PRJCONFIG)
endef

#**************************************************************************************
#the main function for compiling execution:
$(eval $(call main_fun))

