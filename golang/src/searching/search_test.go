
package search

import (
    "testing"
    "fmt"
)

func Test_LinearSearch(t *testing.T) {
    searchTest(t, LinearSearch)
}

func Test_BinarySearch(t *testing.T) {
    searchTest(t, BinarySearch)
}


func searchTest(t *testing.T, fn func ([]int, int)(bool, int)) {
    testData := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
    testDataOneItem := []int{1}
    testDataEmpty := []int{}

    found ,pos := fn(testData, 1)
    if found == false || pos != 0 {
        fmt.Printf("Failed: First position \n")
        t.Fail()
    }

    found ,pos = fn(testData, 20)
    if found == false || pos != 19 {
        fmt.Printf("Failed: Last position\n")
        t.Fail()
    }

    found ,pos = fn(testData, 2)
    if found == false || pos != 1 {
        fmt.Printf("Failed: Second position from start\n")
        t.Fail()
    }

    found ,pos = fn(testData, 19)
    if found == false || pos != 18 {
        fmt.Printf("Failed: Second position from last\n")
        t.Fail()
    }

    found ,pos = fn(testData, 21)
    if found == true || pos != 0 {
        fmt.Printf("Failed: Biggest failed search\n")
        t.Fail()
    }

    found ,pos = fn(testData, -1)
    if found == true && pos != 0 {
        fmt.Printf("Failed: Smallest failed search\n")
        t.Fail()
    }

    found ,pos = fn(testDataOneItem, 1)
    if found != true || pos != 0 {
        fmt.Printf("Failed: Single item success \n")
        t.Fail()
    }

    found ,pos = fn(testDataOneItem, 2)
    if found != false || pos != 0 {
        fmt.Printf("Failed: Single item failed larger\n")
        t.Fail()
    }

    found ,pos = fn(testDataOneItem, 0)
    if found != false || pos != 0 {
        fmt.Printf("Failed: Single item failed smaller\n")
        t.Fail()
    }

    found ,pos = fn(testDataEmpty, 0)
    if found != false || pos != 0 {
        fmt.Printf("Failed: Search an empty set\n")
        t.Fail()
    }
}



