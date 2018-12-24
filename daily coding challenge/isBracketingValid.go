package main

import "fmt"

func main() {
    fmt.Printf("%v -> %v\n", "([])[]({})", isBracketingValid("([])[]({})"))

    fmt.Printf("%v -> %v\n", "([)]", isBracketingValid("([)]"))

    fmt.Printf("%v -> %v\n", "((()", isBracketingValid("((()"))

    fmt.Printf("%v -> %v\n", "", isBracketingValid(""))
}

type stack interface {
    push()
    pop() interface{}
    empty() bool
}

type runeStack []rune

func (runeStack *runeStack) push(rune rune) {
    *runeStack = append(*runeStack, rune)
}

func (runeStack *runeStack) pop() interface{} {
    runeSlice := *runeStack

    length := len(runeSlice)

    if length > 0 {
        rune := runeSlice[length - 1]

        runeSlice = runeSlice[ : length - 1]

        *runeStack = runeSlice

        return rune
    }

    return nil
}

func (runeStack *runeStack)empty() (empty bool) {
    return len(*runeStack) == 0
}

func isBracketingValid(bracketString string) (valid bool) {
    bracketStack := runeStack{}

    validBracketPairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, currentBracket := range bracketString {
        if currentBracket == '(' || currentBracket == '[' || currentBracket == '{' {
            bracketStack.push(currentBracket)
        } else if currentBracket == ')' || currentBracket == ']' || currentBracket == '}' {
            previousBracket := bracketStack.pop().(rune)

            if previousBracket != validBracketPairs[currentBracket] {
                return
            }
        }
    }

    if !bracketStack.empty() {
        return
    }

    return true
}
