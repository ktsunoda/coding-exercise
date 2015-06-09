// http://repl.it/re2/8
// https://www.hackerrank.com/challenges/largest-rectangle
package main

import "bufio"
// import "strconv"
import "fmt"
import "os"

func main() {
    // scanner := bufio.NewScanner(os.Stdin)
    // scanner.Split(bufio.ScanWords)
    // scanner.Scan()
    // count, err := strconv.Atoi(scanner.Text())
    // if err != nil {
    //     return
    // }
    
    // var buildings []int
    
    // for scanner.Scan() {
    //     building, err := strconv.Atoi(scanner.Text())
    //     if err != nil {
    //         return
    //     }
    //     buildings = append(buildings, building)
    // }

    // count := 5
    // buildings := []int{1, 2, 3, 4, 5}
    count := 10
    buildings := []int{8979, 4570, 6436, 5083, 7780, 3269, 5400, 7579, 2324, 2116}
    
    result := 0
    currentResult := 0

    for j := 0; j < count; j++ {
        leftLocation := j
        rightLocation := j
        
        for ; leftLocation > 0 && buildings[leftLocation - 1] >= buildings[j]; leftLocation-- {
        }
        
        for ; rightLocation < count - 1 && buildings[rightLocation + 1] >= buildings[j]; rightLocation++ {
        }
        
        currentResult = (rightLocation - leftLocation + 1) * buildings[j]
        
        if currentResult > result {
            result = currentResult
        }
    }
    
    writer := bufio.NewWriter(os.Stdout)
    fmt.Fprint(writer, result)
    writer.Flush()
}
