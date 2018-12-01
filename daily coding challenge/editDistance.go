package main

import "fmt"

func main() {
    distance := editDistance("kitten", "sitting")
    fmt.Printf("%v, %v -> %v\n", "kitten", "sitting", distance)

    distance = editDistance("test", "tset")
    fmt.Printf("%v, %v -> %v\n", "test", "tset", distance)

    distance = editDistance("giraf", "farig")
    fmt.Printf("%v, %v -> %v\n", "giraf", "farig", distance)
}

func editDistance(string1, string2 string) (distance int) {
    length1 := len(string1)
    length2 := len(string2)

    if length1 > length2{
        distance = length1 - length2

        for index := range string2 {
            if string2[index] != string1[index] {
                distance++
            }
        }

    } else {
        distance = length2 - length1

        for index := range string1 {
            if string1[index] != string2[index] {
                distance++
            }
        }
    }

    return
}
