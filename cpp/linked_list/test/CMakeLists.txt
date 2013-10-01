CMAKE_MINIMUM_REQUIRED(VERSION 2.8.7 FATAL_ERROR)
PROJECT(LinkedListTest)

ExternalProject_Get_Property(googletest source_dir)
ExternalProject_Get_Property(googletest binary_dir)

INCLUDE_DIRECTORIES(${source_dir}/include)
INCLUDE_DIRECTORIES(${LinkedList_SOURCE_DIR}/header)
LINK_DIRECTORIES(${LinkedList_BINARY_DIR})

ADD_EXECUTABLE(LinkedListTest ${PROJECT_SOURCE_DIR}/linked_list_test.cpp)

ADD_DEPENDENCIES(LinkedListTest googletest linkedlist)


if(UNIX)
    set(PTHREAD "-pthread")
endif()

TARGET_LINK_LIBRARIES(LinkedListTest
    ${LinkedList_BINARY_DIR}/${CMAKE_FIND_LIBRARY_PREFIXES}linkedlist.so
    ${binary_dir}/${CMAKE_FIND_LIBRARY_PREFIXES}gtest.a
    ${PTHREAD})
