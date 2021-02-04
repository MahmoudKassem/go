package main

import "fmt"

func main() {
    fmt.Printf("%v -> %v\n", "([])[]({})", isBracketingValid("([])[]({})"))

    fmt.Printf("%v -> %v\n", "([)]", isBracketingValid("([)]"))

    fmt.Printf("%v -> %v\n", "((()", isBracketingValid("((()"))

    fmt.Printf("%v -> %v\n", "", isBracketingValid(""))
}

type stack interface {
    push(interface{})
    pop() interface{}
    empty() bool
}

type runeStack []rune

func (runeStack *runeStack) push(element interface{}) {
    rune := element.(rune)

    *runeStack = append(*runeStack, rune)
}

func (runeStack *runeStack) pop() (element interface{}) {
    length := len(*runeStack)

    if length > 0 {
        element = (*runeStack)[length - 1]

        *runeStack = (*runeStack)[ : length - 1]

        return
    }

    return
}

func (runeStack *runeStack) empty() (empty bool) {
    return len(*runeStack) == 0
}

func isBracketingValid(bracketString string) (valid bool) {
    bracketStack := new(runeStack)

    validBracketPairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, currentBracket := range bracketString {
        if currentBracket == '(' || currentBracket == '[' || currentBracket == '{' {
            bracketStack.push(currentBracket)
        } else if currentBracket == ')' || currentBracket == ']' || currentBracket == '}' {
            if bracketStack.empty() {
              return
            }

            previousBracket := bracketStack.pop().(rune)
            if previousBracket != validBracketPairs[currentBracket] {
                return
            }
        }
    }

    return bracketStack.empty()
}
