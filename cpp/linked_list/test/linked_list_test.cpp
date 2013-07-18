
#include "linked_list.h"
#include "gtest/gtest.h"
#include <string>

TEST(LinkedList, Construction) {
    LinkedList<int> list;
}

TEST(LinkedList, Insert) {
    LinkedList<int> list;
    int test1 = 12;
    list.insert(test1);
}


int main(int argc, char **argv) {
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
