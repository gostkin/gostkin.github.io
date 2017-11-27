package main

import "unicode"

func RemoveEven(items []int) []int {
    ret := make([]int, 0)
    for _, item := range items {
        if item % 2 == 1 {
            ret = append(ret, item)
        }
    }
    return ret
}

func PowerGenerator(number int) func() int {
    ret := 1
    return func() int {
        ret *= number
        return ret
    }
}

func DifferentWordsCount(s string) int {
    set := make(map[string]bool)
    word := ""
    ret := 0
    for _, c := range (s + " ") {
        if !unicode.IsDigit(c) {
            word += string(unicode.ToLower(c))
        } else if word != "" {
            if !set[word] {
                ret += 1
            }

            set[word] = true
            word = ""
        }
    }

    return ret
}
