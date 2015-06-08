// http://repl.it/r5R/12
// http://www.careercup.com/question?id=5160350129324032
package main

import "fmt"

func main() {
    rotatedSorted := []int{7, 8, 9, 10, 1, 2, 3, 4, 5}
    fmt.Println(find(7, rotatedSorted))
    fmt.Println(find(8, rotatedSorted))
    fmt.Println(find(9, rotatedSorted))
    fmt.Println(find(10, rotatedSorted))
    fmt.Println(find(1, rotatedSorted))
    fmt.Println(find(2, rotatedSorted))
    fmt.Println(find(3, rotatedSorted))
    fmt.Println(find(4, rotatedSorted))
    fmt.Println(find(5, rotatedSorted))
    fmt.Println(find(6, rotatedSorted))
}

func find(search int, array []int) int {
    currentFront := 0
    currentBack := len(array)
    var currentIndex int

    for {
        currentIndex = (currentBack + currentFront) / 2
        // fmt.Println(currentFront, currentBack, currentIndex, array[currentFront:currentBack])
        if array[currentIndex] == search {
            return currentIndex
        } else if currentFront == currentBack - 1 {
            return -1
        }
        
        if array[currentIndex] <= array[currentBack - 1] {
            // The last half is sorted
            if search > array[currentIndex] && search <= array[currentBack - 1] {
                currentFront = currentIndex
            } else {
                currentBack = currentIndex
            }
        } else {
            // The first half is sorted
            if search >= array[currentFront] && search < array[currentIndex] {
                currentBack = currentIndex
            } else {
                currentFront = currentIndex
            }
        }
    }
    
    return -1
}

