package main

import "fmt"

func main() {
    fmt.Printf("%v, %v -> %v\n", "kitten", "sitting", editDistance("kitten", "sitting"))

    fmt.Printf("%v, %v -> %v\n", "test", "tset", editDistance("test", "tset"))

    fmt.Printf("%v, %v -> %v\n", "giraf", "farig", editDistance("giraf", "farig"))
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
