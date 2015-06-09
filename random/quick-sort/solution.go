// http://repl.it/r9r/1
package main

import "fmt"

func QuickSort(data []int) []int {
    if len(data) <= 1 {
        return data
    }

    var less []int
    var pivots []int
    var more []int

    pivotIndex := len(data) / 2
    pivot := data[pivotIndex]
    
    for i := 0; i < len(data); i++ {
        if data[i] < pivot {
            less = append(less, data[i])
        } else if data[i] > pivot {
            more = append(more, data[i])
        } else {
            pivots = append(pivots, data[i])
        }
    }
    
    return append(append(QuickSort(less), pivots...), QuickSort(more)...)
}


func main() {
    data := []int{0, 5, 1, 7, 4, 8, 2, 3, 6, 9}
    
    fmt.Println(data)
    data = QuickSort(data)
    fmt.Println(data)
}

