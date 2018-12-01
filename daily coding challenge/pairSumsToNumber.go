package main

import "fmt"

func main() {
    a, b, _:= pairSumsToNumber(17, []int{15, 7, 3, 10})
    fmt.Printf("%v %v -> (%v %v)\n", 17, []int{15, 7, 3, 10}, a, b)

    _, _, found := pairSumsToNumber(2, []int{0, 1})
    fmt.Printf("%v %v -> %v\n", 2, []int{0, 1}, found)

    a, b, _ = pairSumsToNumber(0, []int{-1, 0, 1})
    fmt.Printf("%v %v -> (%v %v)\n", 0, []int{-1, 0, 1}, a, b)

    a, b, _ = pairSumsToNumber(10, []int{5, 0, 5})
    fmt.Printf("%v %v -> (%v %v)\n", 10, []int{5, 0, 5}, a, b)
}

func pairSumsToNumber(number int, list []int) (a int, b int, found bool) {
    differenceMap := map[int]int{}

    for _, value := range list {
        difference, found := differenceMap[value]

        if found {
            return value, difference, true
        }

        differenceMap[number - value] = value
    }

    return
}
