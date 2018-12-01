package main

import "fmt"

func main() {
    list1 := []interface{}{1, 2, 3, 4}
    list2 := []interface{}{"a", "b", "c", "d"}
    list3 := []interface{}{}
    list4 := []interface{}{1}
    list5 := []interface{}{1, 3, 1}
    list6 := []interface{}{"a", "b", "c", "c", "b", "a"}
    list7 := []interface{}{1, 2, []interface{}{3, []interface{}{4, 5}}, 6}
    list8 := []interface{}{"a", "b", []interface{}{"c", []interface{}{"d", "e"}}, "f"}
    list9 := []interface{}{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}
    list10 := []interface{}{1, 1, 1, 1, 2, 3, 3, 1, 1, 4, 5, 5, 5, 5}

    fmt.Println("#1 last element of a list")
    fmt.Printf("%v -> %v\n", list1, myLast(list1))
    fmt.Printf("%v -> %v\n", list2, myLast(list2))
    fmt.Printf("%v -> %v\n\n", list3, myLast(list3))

    fmt.Println("#2 last but one element of a list")
    fmt.Printf("%v -> %v\n", list1, lastButOne(list1))
    fmt.Printf("%v -> %v\n", list2, lastButOne(list2))
    fmt.Printf("%v -> %v\n\n", list3, lastButOne(list3))

    fmt.Println("#3 k'th element of a list")
    fmt.Printf("%v, %v -> %v\n", list1, 2, elementAt(list1, 2))
    fmt.Printf("%v, %v -> %v\n", list2, 2, elementAt(list2, 2))
    fmt.Printf("%v, %v -> %v\n\n", list3, 2, elementAt(list3, 2))

    fmt.Println("#4 number of elements in a list")
    fmt.Printf("%v -> %v\n", list1, myLength(list1))
    fmt.Printf("%v -> %v\n", list2, myLength(list2))
    fmt.Printf("%v -> %v\n\n", list3, myLength(list3))

    fmt.Println("#5 reverse a list")
    fmt.Printf("%v -> %v\n", list1, myReverse(list1))
    fmt.Printf("%v -> %v\n", list2, myReverse(list2))
    fmt.Printf("%v -> %v\n\n", list3, myReverse(list3))

    fmt.Println("#6 is list a palindrome")
    fmt.Printf("%v -> %v\n", list1, isPalindrom(list1))
    fmt.Printf("%v -> %v\n", list2, isPalindrom(list2))
    fmt.Printf("%v -> %v\n", list3, isPalindrom(list3))
    fmt.Printf("%v -> %v\n", list4, isPalindrom(list4))
    fmt.Printf("%v -> %v\n", list5, isPalindrom(list5))
    fmt.Printf("%v -> %v\n\n", list6, isPalindrom(list6))

    fmt.Println("#7 flatten a nested list")
    fmt.Printf("%v -> %v\n", list1, myFlatten(list1))
    fmt.Printf("%v -> %v\n", list2, myFlatten(list2))
    fmt.Printf("%v -> %v\n", list3, myFlatten(list3))
    fmt.Printf("%v -> %v\n", list7, myFlatten(list7))
    fmt.Printf("%v -> %v\n\n", list8, myFlatten(list8))

    fmt.Println("#8 remove consecutive duplicates in a list")
    fmt.Printf("%v -> %v\n", list1, compress(list1))
    fmt.Printf("%v -> %v\n", list2, compress(list2))
    fmt.Printf("%v -> %v\n", list3, compress(list3))
    fmt.Printf("%v -> %v\n", list9, compress(list9))
    fmt.Printf("%v -> %v\n\n", list10, compress(list10))

    fmt.Println("#9 pack consecutive duplicates of elements into sublists")
    fmt.Printf("%v -> %v\n", list1, pack(list1))
    fmt.Printf("%v -> %v\n", list2, pack(list2))
    fmt.Printf("%v -> %v\n", list3, pack(list3))
    fmt.Printf("%v -> %v\n", list9, pack(list9))
    fmt.Printf("%v -> %v\n\n", list10, pack(list10))

    fmt.Println("#10 run-length encoding of a list")
    fmt.Printf("%v -> %v\n", list1, encode(list1))
    fmt.Printf("%v -> %v\n", list2, encode(list2))
    fmt.Printf("%v -> %v\n", list3, encode(list3))
    fmt.Printf("%v -> %v\n", list9, encode(list9))
    fmt.Printf("%v -> %v\n\n", list10, encode(list10))
}

func myLast(list []interface{}) (last interface{}) {
    for index := range list {
        last = list[index]
    }

    return
}

func lastButOne(list []interface{}) (lastButOne interface{}) {
    for index := range list {
        if index > 1 {
            lastButOne = list[index - 1]
        }
    }

    return
}

func elementAt(list []interface{}, position int) (elementAt interface{}) {
    for index := range list {
        if index == position - 1 {
            elementAt = list[index]
        }
    }

    return
}

func myLength(list []interface{}) (myLength int) {
    for index := range list {
        myLength += index - (index - 1)
    }

    return
}

func myReverse(list []interface{}) (reversedList []interface{}) {
    for _, element := range list {
        reversedList = append([]interface{}{element}, reversedList...)
    }

    return
}

func isPalindrom(list []interface{}) (isPalindrom bool) {
    reversedList := myReverse(list)

    for index := range list {
        if list[index] != reversedList[index] {
            return
        }
    }

    return true
}

func myFlatten(nestedList []interface{}) (flattendedList []interface{}) {
    for _, element := range nestedList {
        list, isList := element.([]interface{})

        if isList {
            flattendedList = append(flattendedList, myFlatten(list)...)
        } else {
            flattendedList = append(flattendedList, element)
        }
    }

    return
}

func compress(list []interface{}) (compressedList []interface{}) {
    var currentValue interface{}

    for _, element := range list {
        if currentValue != element {
            compressedList = append(compressedList, element)

            currentValue = element
        }
    }

    return
}

func pack(list []interface{}) (packedList []interface{}) {
    var currentValue interface{}

    consecutiveDuplicatesList := []interface{}{}

    for _, element := range list {
        if currentValue != element {
            currentValue = element

            if myLength(consecutiveDuplicatesList) > 0 {
                packedList = append(packedList, consecutiveDuplicatesList)
            }

            consecutiveDuplicatesList = []interface{}{currentValue}
        } else {
            consecutiveDuplicatesList = append(consecutiveDuplicatesList, element)
        }
    }

    if myLength(consecutiveDuplicatesList) > 0 {
        packedList = append(packedList, consecutiveDuplicatesList)
    }

    return
}

func encode(list []interface{}) (encodedList []interface{}) {
    packedList := pack(list)

    for _, consecutiveDuplicatesList := range packedList {
        consecutiveDuplicatesList := consecutiveDuplicatesList.([]interface{})

        encodedList = append(encodedList, []interface{}{myLength(consecutiveDuplicatesList), consecutiveDuplicatesList[0]})
    }

    return
}
