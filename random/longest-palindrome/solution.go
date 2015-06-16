// http://repl.it/snC/10
package main

import "fmt"

func longestPalindrome(text *string) *string {
    var palindrome *string

    if text != nil {
        palindrome = new(string)
        
        for i := 0; i < len(*text); i++ {
            currentPalindrome := longestPalindromeMiddle(text, i)
            
            if len(*currentPalindrome) > len(*palindrome) {
                palindrome = currentPalindrome
            }
        }
    }
    
    return palindrome
}

func longestPalindromeMiddle(text *string, middle int) *string {
    oddPalindrome := longestPalindromeLeftRight(text, middle, middle)
    evenPalindrome := longestPalindromeLeftRight(text, middle, middle + 1)
    
    if len(*oddPalindrome) > len(*evenPalindrome) {
        return oddPalindrome
    }
    
    return evenPalindrome
}

func longestPalindromeLeftRight(text *string, left int, right int) *string {
    palindrome := ""

    for left >= 0 && right < len(*text) && (*text)[left] == (*text)[right] {
        left--
        right++
    }
    
    left++
    right--
    
    if left <= right {
        palindrome = (*text)[left:right + 1]
    }
    
    return &palindrome
}

func sp(s string) *string { return &s }

func main() {
    texts := []*string{
        nil, 
        sp("a"),
        sp("abc"),
        sp("abab"),
        sp("ababcb"),
        sp("ababcba"),
        sp("abbac"),
        sp("abbacaacabb"),
        sp("abcacbaXXYabcacbY"),
        sp("abcacbaXXYabcacbYabbaabbaabba"),
    }
    
    answers := []*string{
        nil,
        sp("a"),
        sp("a"),
        sp("aba"),
        sp("aba"),
        sp("abcba"),
        sp("abba"),
        sp("bbacaacabb"),
        sp("abcacba"),
        sp("abbaabbaabba"),
    }
    
    for i := 0; i < len(texts); i++ {
        text, answer := texts[i], answers[i]
        
        if text == nil {
            fmt.Println(text, answer, longestPalindrome(text) == answer)
        } else {
            fmt.Println(*text, *answer, *longestPalindrome(text) == *answer)
        }
    }
}
