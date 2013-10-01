
package search


// LinearSearch steps through a slice and finds the first occurence of search

func LinearSearch(intSlice []int, search int) (bool, int) {
    found := false
    position := 0

    for i := 0; !found && i < len(intSlice) ; i++ {
        if intSlice[i] == search {
            found = true
            position = i
        }
    }

    return found, position
}

// BinarySearch takes a sorted slice and a search term
// returns bool true for found and false for not found
// also returns the index
func BinarySearch(sortedData []int, search int) (bool, int) {

    low := 0
    high := len(sortedData) - 1

    for ;high >= low; {
        // get midway index
        mid := int((high - low) / 2) + low

        if sortedData[mid] < search {
            low = mid + 1
        } else if sortedData[mid] > search {
            high = mid - 1
        } else {
            return true, mid
        }
    }

    return false, 0
}



