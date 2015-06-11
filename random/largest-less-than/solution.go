// http://repl.it/sGp/11
// Largest Less Than
package main

import "fmt"
import "strconv"

type Tree struct {
    data int
    left *Tree
    right *Tree
}

func (t *Tree) String() string {
    var result string;
    
    result = strconv.Itoa(t.data)
    
    if t.left != nil || t.right != nil {
        result += "["
        
        if t.left != nil {
            result += t.left.String()
        } else {
            result += ""
        }
        
        result += ","
        
        if t.right != nil {
            result += t.right.String()
        } else {
            result += ""
        }
        
        result += "]"
    }
    
    return result
}

func largestLessThan(tree *Tree, data int) *int {
    current := tree
    var result *int
    
    for current != nil {
        if data > current.data {
            if result == nil {
                result = new(int)
            }
            
            *result = current.data
            current = current.right
        } else {
            current = current.left
        }
    }
    
    return result
}

func main() {
    root := &Tree{10, &Tree{5, &Tree{1, nil, nil}, &Tree{8, nil, nil}}, &Tree{15, nil, &Tree{20, nil, nil}}}
    
    fmt.Println(root.String());
    
    fmt.Println(1, largestLessThan(root, 1))
    fmt.Println(2, *largestLessThan(root, 2))
    fmt.Println(3, *largestLessThan(root, 3))
    fmt.Println(4, *largestLessThan(root, 4))
    fmt.Println(5, *largestLessThan(root, 5))
    fmt.Println(6, *largestLessThan(root, 6))
    fmt.Println(7, *largestLessThan(root, 7))
    fmt.Println(8, *largestLessThan(root, 8))
    fmt.Println(9, *largestLessThan(root, 9))
    fmt.Println(10, *largestLessThan(root, 10))
    fmt.Println(11, *largestLessThan(root, 11))
    fmt.Println(12, *largestLessThan(root, 12))
    fmt.Println(13, *largestLessThan(root, 13))
    fmt.Println(14, *largestLessThan(root, 14))
    fmt.Println(15, *largestLessThan(root, 15))
    fmt.Println(16, *largestLessThan(root, 16))
    fmt.Println(17, *largestLessThan(root, 17))
    fmt.Println(18, *largestLessThan(root, 18))
    fmt.Println(19, *largestLessThan(root, 19))
    fmt.Println(20, *largestLessThan(root, 20))
    fmt.Println(21, *largestLessThan(root, 21))
}
