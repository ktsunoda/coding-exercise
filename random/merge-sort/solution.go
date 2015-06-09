// http://repl.it/r80/2
package main

import "fmt"

func MergeSort(data []int) []int {
    if (len(data) <= 1) {
        return data
    }
    
    middle := len(data) / 2
    
    left := MergeSort(data[:middle])
    right := MergeSort(data[middle:])
    
    fmt.Println(left, right)
    
    var result []int
    
    for len(left) > 0 || len(right) > 0 {
        mergeLeft := false
        mergeRight := false
        if len(left) > 0 && len(right) > 0 {
            if left[0] < right[0] {
                mergeLeft = true
            } else {
                mergeRight = true
            }
        } else if len(left) > 0 {
            mergeLeft = true
        } else {
            mergeRight = true
        }

        if mergeLeft {
            result = append(result, left[0])
            left = left[1:]
        } else if mergeRight {
            result = append(result, right[0])
            right = right[1:]
        }
    }
    
    return result
}

func main() {
    data := []int{0, 6, 7, 1, 4, 2, 3, 9, 8, 5}
    
    fmt.Println(data)
    fmt.Println(MergeSort(data))
}

