package main

import "fmt"
import "math/rand"
import "strconv"

func smallestNumberFaster(number *string, removeCount int) *string {
    var smallestNumber *string
    
    if number != nil && removeCount >= 0 {
        smallestNumber = new(string)
        *smallestNumber = *number
        
        for i := 0; i < removeCount && i < len(*number); i++ {
            toRemoveIndex := leftMostLargestNumber(smallestNumber)
            *smallestNumber = (*smallestNumber)[0:toRemoveIndex] + (*smallestNumber)[toRemoveIndex + 1:]
        }
    }

    return smallestNumber
}

func leftMostLargestNumber(number *string) int {
    digit, err := strconv.Atoi(string((*number)[0]))
    if err != nil {
        return -1
    }

    index := 0
    
    for i := 1; i < len(*number); i++ {
        if newDigit, err := strconv.Atoi(string((*number)[i])); newDigit > digit && err == nil {
            digit = newDigit
            index = i
        }
    }
    
    return index
}

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

func testSmallestNumbers(number *string, removeCount int) {
    smallestNumber := smallestNumber(number, removeCount)
    smallestNumberFaster := smallestNumberFaster(number, removeCount)
    
    success := false
    if smallestNumber == nil && smallestNumberFaster == nil {
        success = true
    } else if smallestNumber != nil && smallestNumberFaster != nil {
        success = *smallestNumber == *smallestNumberFaster
    }
    
    if number == nil {
        fmt.Print(number)
    } else {
        fmt.Print(*number)
    }
    fmt.Print(", ", removeCount, ";\t")
    if smallestNumber == nil {
        fmt.Print(smallestNumber)
    } else {
        fmt.Print(*smallestNumber)
    }
    fmt.Println(";\tsuccess:", success)
}

func main() {
    var number *string
    testSmallestNumbers(number, 0)
    testSmallestNumbers(number, -1)

    number = new(string)
    
    for i := 0; i < 1; i++ {
        *number = strconv.Itoa(rand.Int())
        
        for j := -1; j <= len(*number); j++ {
            testSmallestNumbers(number, j)
        }
    }
    
    // *number = "12345"
    
    // testSmallestNumbers(number, -1)
    // testSmallestNumbers(number, 0)
    // testSmallestNumbers(number, 1)
    // testSmallestNumbers(number, 2)
    // testSmallestNumbers(number, 3)
    // testSmallestNumbers(number, 4)
    // testSmallestNumbers(number, 5)

    // *number = "12301"
    // testSmallestNumbers(number, 0)
    // testSmallestNumbers(number, 1)
    // testSmallestNumbers(number, 2)
    // testSmallestNumbers(number, 3)
    // testSmallestNumbers(number, 4)
    // testSmallestNumbers(number, 5)
    // testSmallestNumbers(number, 6)
}
