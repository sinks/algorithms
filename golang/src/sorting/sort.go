
package sorting

func BubbleSort(unsorted []int) {
    //sorted 
    for i := 0 ; i < len(unsorted) ; i++ {
        // moving
        for j := 0 ; j < len(unsorted) - i - 1 ; j++ {
            if unsorted[j] > unsorted[j+1] {
                temp := unsorted[j]
                unsorted[j] = unsorted[j+1]
                unsorted[j+1] = temp
            }
        }
    }
}

func InsertionSort(unsorted []int) {
    // current item to insert and division to 
    for i := 1; i < len(unsorted); i++ {
        found := false
        //search sorted section for insertion
        for j := i - 1; !found && j > 0; j-- {
            if unsorted[j+1] < unsorted[j] {
                // swap values
                temp := unsorted[j+1]
                unsorted[j+1] = unsorted[j]
                unsorted[j] = temp
            } else {
                found = true
            }
        }
    }
}


func MergeSort(unsorted []int) {
    // merge part of size 1
    
}

func merge(sortedA []int, sortedB []int) ([]int) {
    merged := []int{}

    aI := 0
    bI := 0

    for ; len(merged) != (len(sortedA) + len(sortedB)); {
        if (aI < len(sortedA)) && ((bI >= len(sortedB)) || sortedA[aI] <= sortedB[bI]) {
            // append and move pA
            merged = append(merged, sortedA[aI])
            aI++
        } else if (bI < len(sortedB)) && ((aI >= len(sortedA))  || sortedB[bI] < sortedA[aI]) {
            // append and move pB
            merged = append(merged, sortedB[bI])
            bI++
        }
    }

    return merged
}

func QuickSort(unsortedSlice []int) {

}

func CountSort(unsortedSlice []int) {

}

func TimSort(unsortedSlice []int) {

}

