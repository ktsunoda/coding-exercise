// http://repl.it/rwW/18
// http://www.careercup.com/question?id=6313112158339072
package main

import "fmt"
import "strconv"

type List struct {
    data int
    next *List
}

func (l *List) String() string {
    var value string
    
    for current := l; current != nil; current = current.next {
        value += strconv.Itoa(current.data) + "->"
    }
    
    return value
}

func SwapPairs(list *List) *List {
    if list == nil || list.next == nil {
        return list
    }
    
    temp := list.next
    list.next = SwapPairs(temp.next)
    temp.next = list
    
    return temp
}

func main() {
    list := &List{0, &List{1, &List{2, &List{3, &List{4, nil}}}}}
    
    fmt.Println(list)
    fmt.Println(SwapPairs(list))
}
