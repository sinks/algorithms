// linked_list.h -- a linked list - template
#ifndef LINKED_LIST_H_
#define LINKED_LIST_H_

#include <cstddef>
#include <iostream>


template <class Type>
class Element {
    private:
        Element<Type> *m_next;
        Type *m_data;

    public:
        Element();
        //~Element();
        
        // Getters
        Type *data();
        Element<Type> *next();
        
        // Setters
        void setData(Type &data);
        void setNext(Element &element);
};

template <class Type>
Element<Type>::Element() {
    this->m_next = nullptr;
    this->m_data = nullptr;
}

template <class Type>
Type *Element<Type>::data() {
    return this->m_data;
}

template <class Type>
Element<Type> *Element<Type>::next() {
    return this->m_next;
}

template <class Type>
void Element<Type>::setData(Type &data) {
    this->m_data = &data;
}

template <class Type>
void Element<Type>::setNext(Element<Type> &element) {
    this->m_next = &element;
}


/*!
 * Linked List Template implementation
 */
template <class Type>
class LinkedList {
    private:
        
        Element<Type> *head;
        int count;

    public:
        // constructor
        LinkedList();
        // destructor
        ~LinkedList();
        
        // insert elements
        Element<Type> *push(Type &data);
        //Element<Type> insert_after(Element &element, Type &data);
        //Element<Type> insert_before(Element &element, Type &data);

        // remove elements 
        Type *pop();
        //Type *remove(Element &element);
        //Type *remove_after(Element &element);
        //Type *remove_before(Element &element);
        //void remove_matching(Type &data);
        
        // find
        //Element<Type> *find(Element<Type> &after=nullptr);

        
        // is list empty
        bool is_empty() const;
        // return amount of items in the list
        int size() const;
};


template <class Type>
LinkedList<Type>::LinkedList() {
    this->head = nullptr;
    this->count = 0;
}

template <class Type>
LinkedList<Type>::~LinkedList() {
    Element<Type> *element = this->head;
    Element<Type> *next = nullptr;
    while(element != nullptr) {
        next = element->next();
        delete element;
        element = next;
    }
}

/*!
 * push insert Type at the head
 * @param data The data to be inserted 
 * @return pointer to the element just inserted
 */
template <class Type>
Element<Type> *LinkedList<Type>::push(Type &data) {
    Element<Type> *element = new Element<Type>();
    element->setNext(*(this->head));
    element->setData(data);

    this->head = element;    
    this->count++;

    return element;
}

/*!
 * pop - remove the head element
 * @return the data that was inserted
 */
template <class Type>
Type *LinkedList<Type>::pop() {
    
    if(this->head != nullptr) {
        Element<Type> *next = this->head->next();
        Type *data = this->head->data();
        delete this->head;
        this->head = next;
        this->count--;
        return data;
    }
    else {
        return nullptr;
    }

}

/*!
 * is_empty
 * @return true if the list is empty
 */
template <class Type>
bool LinkedList<Type>::is_empty() const {
    return (this->head == nullptr);
}

/*!
 *  size
 *  @return elements in the list
 */
template <class Type>
int LinkedList<Type>::size() const {
    return this->count;
}


#endif

