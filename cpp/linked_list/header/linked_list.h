// linked_list.h -- a linked list - template
#ifndef LINKED_LIST_H_
#define LINKED_LIST_H_

#include<cstddef>

template <class Type>
class LinkedList {
    private:
        struct node {
            node *next;
            node *prev;
            Type *data;
        };
        
        node *head;
        node *tail;

    public:
        // constructor
        LinkedList();
        // destructor
        //~LinkedList();

        // insert at end
        void insert(Type data);
        // remove 
        //void remove(Type data);
};

template <class Type>
LinkedList<Type>::LinkedList() {
    this->head = nullptr;
    this->tail = 0;//nullptr;
}

template <class Type>
void LinkedList<Type>::insert(Type data) {
    node dataNode;
    dataNode.next = 0;
    dataNode.prev = 0;
    dataNode.data = &data;
}


#endif

