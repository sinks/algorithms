
package sorting

import (
    "testing"
    "fmt"
)

func Test_BubbleSort(t *testing.T) {
    sortSmallTest(t, BubbleSort)
}

func Test_InsertionSort(t *testing.T) {
    sortSmallTest(t, InsertionSort)
}

func Test_merge(t *testing.T) {
    one := []int{1,5}
    two := []int{4,7}

    sorted := merge(one, two)

    if len(sorted) != 4 {
        fmt.Printf("Failed merge: Merge failed to combine\n")
        t.FailNow()
    }

    if sorted[0] != 1 {
        t.Fail()
    }

    if sorted[1] != 4 {
        t.Fail()
    }

    three := []int{1,6,9}
    four := []int{4,5,8,10}

    sorted = merge(three, four)

    if len(sorted) != 7 {
        t.FailNow()
    }

    if sorted[0] != 1 {
        t.Fail()
    }

    if sorted[6] != 10 {
        t.Fail()
    }
}


func sortSmallTest(t *testing.T, fn func ([]int)) {

    unsorted := []int{1,9,2,8,3,7,4,6,5}

    fn(unsorted)

    for i := 0; i < len(unsorted); i++ {
        if unsorted[i] != (i + 1) {
            fmt.Printf("Failed Small Test: Position %d\n", i)
            t.Fail()
        }
    }
}


