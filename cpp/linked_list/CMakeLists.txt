CMAKE_MINIMUM_REQUIRED(VERSION 2.8.7 FATAL_ERROR)

#SET(CMAKE_CXX_FLAGS "-std=c++11 -stdlib=libc++ -Wall")
SET(CMAKE_CXX_FLAGS "-std=c++11 -Wall")

PROJECT(LinkedList)

ADD_SUBDIRECTORY(test)

INCLUDE_DIRECTORIES(${LinkedList_SOURCE_DIR}/header)

ADD_LIBRARY(linkedlist SHARED src/linked_list.cpp)
