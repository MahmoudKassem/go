package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    list1 := []interface{}{1, 2, 3, 4}
    list2 := []interface{}{"a", "b", "c", "d"}
    list3 := []interface{}{}

    fmt.Println("#21 insert an element at a given position into a list")
    fmt.Printf("%v, %v, %v -> %v\n", list1, 5, 3, insertAt(list1, 5, 3))
    fmt.Printf("%v, %v, %v -> %v\n", list2, "e", 3, insertAt(list2, "e", 3))
    fmt.Printf("%v, %v, %v -> %v\n\n", list3, 5, 3, insertAt(list3, 5, 3))

    fmt.Println("#22 create a list containing all integers within a given range")
    fmt.Printf("%v, %v -> %v\n", 3, 8, myRange(3, 8))
    fmt.Printf("%v, %v -> %v\n", -8, -3, myRange(-8, -3))
    fmt.Printf("%v, %v -> %v\n\n", 8, 3, myRange(8, 3))

    fmt.Println("#23 extract a given number of randomly selected elements from a list")
    fmt.Printf("%v, %v -> %v\n", list1, 3, randomSelection(list1,3))
    fmt.Printf("%v, %v -> %v\n", list2, 3, randomSelection(list2,3))
    fmt.Printf("%v, %v -> %v\n\n", list3, 3, randomSelection(list3,3))

    fmt.Println("#24 draw N different random numbers from the set 1..M")
    fmt.Printf("%v, %v -> %v\n", 6, 49, lotto(6,49))
    fmt.Printf("%v, %v -> %v\n", 0, 49, lotto(0,49))
    fmt.Printf("%v, %v -> %v\n", 6, 0, lotto(6,0))
    fmt.Printf("%v, %v -> %v\n\n", 49, 6, lotto(49,6))

    fmt.Println("#25 generate a random permutation of the elements of a list")
    fmt.Printf("%v -> %v\n", list1, randomPermutation(list1))
    fmt.Printf("%v -> %v\n", list2, randomPermutation(list2))
    fmt.Printf("%v -> %v\n\n", list3, randomPermutation(list3))
}

func insertAt(list []interface{}, valueToBeInserted interface{}, position int) (updatedList []interface{}) {
    for index := 0; index < len(list); index++ {
        if index + 1 == position {
            updatedList = append(updatedList, valueToBeInserted, list[index])
        } else {
            updatedList = append(updatedList, list[index])
        }
    }

    return
}

func myRange(lowerBound, higherBound int) (resultList []int) {
    for lowerBound <= higherBound {
        resultList = append(resultList, lowerBound)

        lowerBound++
    }

    return
}

func removeAt(list []interface{}, position int) (resultList []interface{}, removedElement interface{}) {
    for index := 0; index < len(list); index++ {
        if (index + 1) == position {
            removedElement = list[index]
        } else {
            resultList = append(resultList, list[index])
        }
    }

    return
}

func randomSelection(list []interface{}, numberOfElements int) (resultList []interface{}) {
    randomNumberGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

    for numberOfElements > 0 {
        length := len(list)

        if length == 0 {
            return
        }

        randomNumber := randomNumberGenerator.Intn(length) + 1

        reducedList, randomElement := removeAt(list, randomNumber)

        list = reducedList

        resultList = append(resultList, randomElement)

        numberOfElements--
    }

    return
}

func lotto(numberOfElements, higherBound int) (resultList []int) {
    if higherBound < numberOfElements || higherBound <= 0 {
        return
    }

    numberInterval := myRange(1, higherBound)

    list := make([]interface{}, higherBound)

    for index, number := range numberInterval {
        list[index] = number
    }

    randomNumbers := randomSelection(list, numberOfElements)

    resultList = make([]int, numberOfElements)

    for index, randomNumber := range randomNumbers {
        resultList[index] = randomNumber.(int)
    }

    return
}

func randomPermutation(list []interface{}) (permutatedList []interface{}) {
    return randomSelection(list, len(list))
}
