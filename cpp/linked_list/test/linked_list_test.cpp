
#include "linked_list.h"
#include "gtest/gtest.h"
#include <string>

TEST(LinkedList, Construction) {
    LinkedList<int> list;
}

TEST(LinkedList, Destruction) {
    LinkedList<int> list;
    
    int test1 = 12;
    list.push(test1);
}

TEST(LinkedList, Push) {
    LinkedList<int> list;
    
    int test1 = 12;
    Element<int> *test1Ele = list.push(test1);
    ASSERT_EQ(test1, *(test1Ele->data()));

    int test2 = 13;
    Element<int> *test2Ele = list.push(test2);
    ASSERT_EQ(test2, *(test2Ele->data()));
}

TEST(LinkedList, IsEmpty) {
    LinkedList<int> list;
    
    ASSERT_TRUE(list.is_empty());

    int test1 = 12;
    list.push(test1);
    
    ASSERT_FALSE(list.is_empty());

}

TEST(LinkedList, Size) {
    LinkedList<int> list;
    
    ASSERT_EQ(0, list.size());

    int test1 = 12;
    Element<int> *returnedElement = list.push(test1);
    
    ASSERT_EQ(1, list.size());
    
    list.pop();

    ASSERT_EQ(0, list.size());
    

}


int main(int argc, char **argv) {
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
