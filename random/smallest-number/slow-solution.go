// http://repl.it/spO/9
package main

import "fmt"
import "math/rand"
import "strconv"

func smallestNumber(number *string, removeCount int) *string {
    var smallestNumber *string
    
    if number != nil && removeCount >= 0 {
        smallestNumber = new(string)
        smallestNumberHelper(number, removeCount, smallestNumber)
    }

    return smallestNumber
}

func smallestNumberHelper(number *string, removeCount int, smallestNumber *string) {
    if removeCount == 0 {
        smaller := false
        
        if *smallestNumber == "" {
            smaller = true
        } else {
            currentSmallestNumber, err1 := strconv.Atoi(*smallestNumber)
            currentNumber, err2 := strconv.Atoi(*number)
            
            smaller = err1 == nil && err2 == nil && currentSmallestNumber > currentNumber
        }
        
        if smaller {
            *smallestNumber = *number
        }
    } else {
        for i := 0; i < len(*number); i++ {
            nextNumber := (*number)[0:i] + (*number)[i + 1:]
            smallestNumberHelper(&nextNumber, removeCount - 1, smallestNumber)
        }
    }
}

func testSmallestNumber(number *string, removeCount int) {
    smallestNumber := smallestNumber(number, removeCount)

    if number == nil {
        fmt.Print(number)
    } else {
        fmt.Print(*number)
    }
    fmt.Print(", ", removeCount, ";\t")
    if smallestNumber == nil {
        fmt.Println(smallestNumber)
    } else {
        fmt.Println(*smallestNumber)
    }
}

func main() {
    var number *string
    testSmallestNumber(number, 0)
    testSmallestNumber(number, -1)

    number = new(string)
    
    for i := 0; i < 1; i++ {
        *number = strconv.Itoa(rand.Int() % 100000)
        
        for j := -1; j <= len(*number) + 1; j++ {
            testSmallestNumber(number, j)
        }
    }
}
