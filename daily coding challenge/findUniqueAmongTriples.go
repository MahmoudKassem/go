package main

import (
    "fmt"
    "math"
)

func main() {
    uniqueElement := findUniqueAmongTriples([]int32{6, 1, 3, 3, 3, 6, 6})
    fmt.Printf("%v -> %v\n", []int32{6, 1, 3, 3, 3, 6, 6}, uniqueElement)

    uniqueElement = findUniqueAmongTriples([]int32{13, 19, 13, 13})
    fmt.Printf("%v -> %v\n", []int32{13, 19, 13, 13}, uniqueElement)

    uniqueElement = findUniqueAmongTriples([]int32{10, -13, -19, -13, -13, 10, 10})
    fmt.Printf("%v -> %v\n", []int32{10, -13, -19, -13, -13, 10, 10}, uniqueElement)
}

func findUniqueAmongTriples(list []int32) (uniqueElement int32) {
    for bitPosition := 0; bitPosition < 32; bitPosition++ {
        onesCount := 0

        oneAtBitPosition := int32(1 << uint32(bitPosition))

        for _, element := range list {
            if element & oneAtBitPosition > 0 {
                onesCount++
            }
        }

        if onesCount % 3 != 0 {
            uniqueElement |= oneAtBitPosition
        }
    }

    if uniqueElement & int32(1 << uint32(30)) > 0 {
        uniqueElement += int32(math.Pow(2, 31) + 1)
    }

    return
}
