

package pattern_matching


func shiftAnd(pattern string, text string) {

}

func shiftOr(pattern string, text string) {

}

// Horspool use horspool algorithm for pattern matching
func Horspool(pattern string, text *string) ([]int)  {
    matches := []int{}

    if len(pattern) == 0 || len(*text) == 0 {
        return matches
    }

    shiftMap := horspoolShiftMap(pattern)

    // find match
    for textI := 0; textI < len(*text) - len(pattern) + 1; {

        // loop over pattern
        match := false
        for pPos, tPos := 0, textI; pPos < len(pattern); pPos++ {
            if pattern[pPos] == (*text)[tPos] {
                match = true
            } else {
                match = false
                pPos = len(pattern)
            }
            tPos++
        }

        if match {
            matches = append(matches, textI)
        }

        // shift
        shiftAmount, found := shiftMap[(*text)[textI + len(pattern) - 1]]
        if found {
            textI += shiftAmount
        } else {
            textI += len(pattern)
        }
    }

    return matches
}

func horspoolShiftMap(pattern string) (map[uint8]int) {

    // create shift map; backwards n-1...1
    // baabc  =   b:1 a:2   *:len(pattern)

    shiftMap := map[uint8]int{}

    for shiftI, shift := len(pattern) - 2, 1; shiftI >= 0 ;  {
        _, there := shiftMap[pattern[shiftI]]
        if !there {
            shiftMap[pattern[shiftI] ] = shift
        }
        shiftI--
        shift++
    }
    return shiftMap
}



