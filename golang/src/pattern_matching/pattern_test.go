
package pattern_matching

import "testing"
import "io/ioutil"
import "bytes"

func Test_horspoolShiftMap(t *testing.T) {

    shiftMap := horspoolShiftMap("ababc")
    if shiftMap['b'] != 1 {
        t.Fail()
    }

    if shiftMap['a'] != 2 {
        t.Fail()
    }
}


func Test_Horspool(t *testing.T) {

    text := "abc"
    matches := Horspool("abc", &text)

    if len(matches) != 1 && matches[0] != 0 {
        t.Fail()
    }

    text = "a"
    matches = Horspool("abcc", &text)

    if len(matches) != 0 {
        t.Fail()
    }

    matches = Horspool("", &text)

    if len(matches) != 0 {
        t.Fail()
    }

    text ="aabcaabcab"
    matches = Horspool("abc", &text)

    if len(matches) != 2 && matches[0] != 1 && matches[1] != 5 {
        t.Fail()
    }

    test, _ := ioutil.ReadFile("test_files/advs.txt")
    buffer := bytes.NewBuffer(test)
    testString := buffer.String()

    matches = Horspool("photograph", &testString)
    if len(matches) == 0 {
        t.Fail()
    }
    for _ ,textI := range matches {

        if testString[textI:textI+10] != "photograph" {
            t.Fail()
        }
    }
}

// Benchmark_Horspool loads a text file (adventures of sherlock holmes) and 
// searches for photograph
func Benchmark_Horspool(b *testing.B) {
    b.StopTimer()
    test, _ := ioutil.ReadFile("test_files/advs.txt")
    buffer := bytes.NewBuffer(test)
    testString := buffer.String()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        Horspool("photograph", &testString)
    }
}
