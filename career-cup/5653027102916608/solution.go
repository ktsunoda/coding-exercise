// http://repl.it/ruB/4
// http://www.careercup.com/question?id=5653027102916608
package main

import "fmt"
import "regexp"
import "unicode"

func isPalindrome(sentence string) bool {
    begin := 0
    end := len(sentence) - 1
    
    re := regexp.MustCompile("[[:alnum:]]")
    for begin < end {
        for ; !re.MatchString(sentence[begin:begin + 1]); begin++ {
            
        }
        
        for ; !re.MatchString(sentence[end:end + 1]); end-- {
        }
        
        if unicode.ToLower(rune(sentence[begin])) != unicode.ToLower(rune(sentence[end])) {
            return false
        }
        
        begin++
        end--
    }

    return true
}

func main() {
    fmt.Println(isPalindrome("A man, a plan, a canal, Panama!"))
}
