CMAKE_MINIMUM_REQUIRED(VERSION 2.8.7 FATAL_ERROR)
if(APPLE)
  SET(CMAKE_C_COMPILER_INIT "/opt/local/bin/clang-mp-3.3")
  SET(CMAKE_CXX_COMPILER_INIT "/opt/local/bin/clang++-mp-3.3")
  set(CMAKE_C_COMPILER "/opt/local/bin/clang-mp-3.3")  
  set(CMAKE_CXX_COMPILER "/opt/local/bin/clang++-mp-3.3")
elseif(UNIX)
  SET(CMAKE_C_COMPILER_INIT "clang")
  SET(CMAKE_CXX_COMPILER_INIT "clang++")
  set(CMAKE_C_COMPILER "clang")
  set(CMAKE_CXX_COMPILER "clang++")
endif()

PROJECT(Algorithms)

INCLUDE(ExternalProject)

SET_DIRECTORY_PROPERTIES(PROPERTIES EP_PREFIX ${CMAKE_BINARY_DIR}/external)


add_custom_target(doc
${DOXYGEN_EXECUTABLE} ${CMAKE_CURRENT_BINARY_DIR}/Doxyfile
WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}
COMMENT "Generating API documentation with Doxygen" VERBATIM
)

option(GTEST_UPDATE "GTEST_UPDATE" OFF)
if(GTEST_UPDATE)
ExternalProject_Add(
    googletest
    SVN_REPOSITORY  http://googletest.googlecode.com/svn/trunk/
    #URL http://googletest.googlecode.com/files/gtest-1.6.0.zip    
    #TIMEOUT 10
    # Force separate output paths for debug and release builds to allow easy
    # identification of correct lib in subsequent TARGET_LINK_LIBRARIES commands
    CMAKE_ARGS -Dgtest_force_shared_crt=ON
    # Disable install step
    INSTALL_COMMAND ""
    # Wrap download, configure and build steps in a script to log output
    LOG_DOWNLOAD ON
    LOG_CONFIGURE ON
    LOG_BUILD ON
)
else()
ExternalProject_Add(
    googletest
    SVN_REPOSITORY  http://googletest.googlecode.com/svn/trunk/
    #URL http://googletest.googlecode.com/files/gtest-1.6.0.zip    
    #TIMEOUT 10
    # Force separate output paths for debug and release builds to allow easy
    # identification of correct lib in subsequent TARGET_LINK_LIBRARIES commands
    CMAKE_ARGS -Dgtest_force_shared_crt=ON
    # Disable install step
    INSTALL_COMMAND ""
    # Disble update step
    UPDATE_COMMAND ""
    # Wrap download, configure and build steps in a script to log output
    LOG_DOWNLOAD ON
    LOG_CONFIGURE ON
    LOG_BUILD ON
)
endif()



ADD_SUBDIRECTORY(linked_list)

