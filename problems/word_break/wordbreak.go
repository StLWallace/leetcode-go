/* https://leetcode.com/problems/word-break/
Not Accepted
I believe this logic works but the implementation is too slow
I read some solutions and the better approach is to use dynamic programming to compare substrings to the word dict
*/

package wordbreak

import (
	"fmt"
	"strings"
)

// Main function
// Steps:
// 1. Check the word dict to see if there are any characters in s not there
//    This can more quickly eliminate some cases
// 2. Use the same approach to eliminate options from the word dict that are linear combos of other words
// 3. Use recursion to see if the beginning of s is one of the words
func wordBreak(s string, wordDict []string) bool {

    if ! hasNeededChars(s, wordDict) {
        return false
    }

    reduced_word_dict := reduceWordDict(wordDict)

    is_valid := false
    is_valid = tryBreaks(s, reduced_word_dict)
    return is_valid
}

// Tries to recursively break off the beginning of s by matching it to a word in the dict
// If the current word in the dict matches the beginning of s, it continues the process with the remaining s until:
// 1. s is completely accounted for
// 2. a piece of s doesn't match any words
func tryBreaks(s string, wordDict []string) bool {
    is_valid := false
    for i := 0; i < len(wordDict); i++ {
        if is_valid {
            break
        }
        word := wordDict[i]
        //fmt.Printf("word: %s\n", word)
        starts_with_word := checkBeginsWithWord(word, s)
        if starts_with_word {
            new_s := removeWord(word, s)
            //fmt.Printf("new_s: %s\n", new_s)
            if len(new_s) == 0 {
                is_valid = true
            } else {
                is_valid = wordBreak(new_s, wordDict)
            }
        }
    }
    return is_valid
}

// Take a word out of the front of a string
// Args:
// word - the word to remove
// s - the string to remove the word from
// returns: s without the initial characters from word
// This would be better to just take len(word) instead of word as an arg
func removeWord(word string, s string) string {
    start := len(word)
    new_s := s[start:]
    return new_s
}

// Returns true if s begins with word, false otherwise
func checkBeginsWithWord(word string, s string) bool {
    end := len(word)
    if len(word) > len(s) {
        return false
    } else if s[0:end] == word {
        return true
    } else {
        return false
    }
}

// Returns a slice of unique characters in a string
func uniqueChars(input string) []string {
    chars := []string{}
    for _, c := range(input) {
        s := string(c)
        if ! sliceContains(chars, s) {
            chars = append(chars, s)
        }
    }
    return chars
}

// Returns true if c is an element of s, false otherwise
func sliceContains(s []string, c string) bool {
    for _, e := range(s) {
        if string(e) == c {
            return true
        }
    }
    return false
}

// Checks if a slice of words contains all the unique characters in s
func hasNeededChars(s string, w []string) bool {
    w_string := strings.Join(w, "")
    w_chars := uniqueChars(w_string)
    s_chars := uniqueChars(s)
    for _, e := range(s_chars) {
        if ! sliceContains(w_chars, string(e)) {
            return false
        }
    }
    return true
}

// Removes elements from a slice of strings that are linear combinations of other items
func reduceWordDict(wordDict []string) []string {
    new_word_dict := []string{}
    for i := 0; i < len(wordDict); i++ {
        this_s := wordDict[i]
        fmt.Printf("this_s: %s\n", this_s)
        var this_word_dict []string
        this_word_dict = append(this_word_dict, wordDict[:i]...)
        this_word_dict = append(this_word_dict, wordDict[i+1:]...)
        fmt.Printf("this_word_dict: %s\n", this_word_dict)
        if !tryBreaks(this_s, this_word_dict) {
            new_word_dict = append(new_word_dict, this_s)
        }
    }
    return new_word_dict
}