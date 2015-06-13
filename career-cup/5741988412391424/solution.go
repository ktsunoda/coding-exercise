// http://repl.it/sHE/6
// http://www.careercup.com/question?id=5741988412391424
package main

import "fmt"

func permutate(toPermutate []string) {
    permutateHelper("", toPermutate)    
}

func permutateHelper(prefix string, toPermutate [] string) {
    if len(toPermutate) > 0 {
        current := toPermutate[0]

        for i := 0; i < len(current); i++ {
            permutateHelper(fmt.Sprint(prefix, string(current[i])), toPermutate[1:])
        }
    } else {
        fmt.Println(prefix)
    }
}

func main() {
    permutate([]string{"red", "fox", "jumped"})
}
