package main

import "fmt"

func main() {
    list1 := []interface{}{}
    list2 := []interface{}{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}
    list3 := []interface{}{1, 1, 1, 1, 2, 3, 3, 1, 1, 4, 5, 5, 5, 5}

    fmt.Println("#11 run-length encoding of a list(modified)")
    fmt.Printf("%v -> %v\n", list1, encodeModified(list1))
    fmt.Printf("%v -> %v\n", list2, encodeModified(list2))
    fmt.Printf("%v -> %v\n\n", list3, encodeModified(list3))

    fmt.Println("#12 decode a run-length encoded list")
    fmt.Printf("%v -> %v\n", encodeModified(list1), decodeModified(encodeModified(list1)))
    fmt.Printf("%v -> %v\n", encodeModified(list2), decodeModified(encodeModified(list2)))
    fmt.Printf("%v -> %v\n\n", encodeModified(list3), decodeModified(encodeModified(list3)))

    fmt.Println("#13 run-length encoding of a list(direct solution)")
    fmt.Printf("%v -> %v\n", list1, encodeDirect(list1))
    fmt.Printf("%v -> %v\n", list2, encodeDirect(list2))
    fmt.Printf("%v -> %v\n\n", list3, encodeDirect(list3))

    fmt.Println("#14 duplicate the elements of a list")
    fmt.Printf("%v -> %v\n", list1, dupli(list1))
    fmt.Printf("%v -> %v\n", list2, dupli(list2))
    fmt.Printf("%v -> %v\n\n", list3, dupli(list3))

    fmt.Println("#15 replicate the elements of a list a given number of times")
    fmt.Printf("%v, %v -> %v\n", list1, 3, repli(list1, 3))
    fmt.Printf("%v, %v -> %v\n", list2, 3, repli(list2, 3))
    fmt.Printf("%v, %v -> %v\n\n", list3, 3, repli(list3, 3))

    fmt.Println("#16 drop every N'th element from a list")
    fmt.Printf("%v, %v -> %v\n", list1, 5, myDrop(list1, 5))
    fmt.Printf("%v, %v -> %v\n", list2, 5, myDrop(list2, 5))
    fmt.Printf("%v, %v -> %v\n\n", list3, 5, myDrop(list3, 5))

    fmt.Println("#17 split a list into two parts; the length of the first part is given")
    firstPart, secondPart := split(list1, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n", list1, 5, firstPart, secondPart)
    firstPart, secondPart = split(list2, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n", list2, 5, firstPart, secondPart)
    firstPart, secondPart = split(list3, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n\n", list3, 5, firstPart, secondPart)

    fmt.Println("#18 extract a slice from a list")
    fmt.Printf("%v, %v, %v -> %v\n", list1, 3, 8, slice(list1, 3, 8))
    fmt.Printf("%v, %v, %v -> %v\n", list2, 3, 8, slice(list2, 3, 8))
    fmt.Printf("%v, %v, %v -> %v\n\n", list3, 3, 8, slice(list3, 3, 8))

    fmt.Println("#19 rotate a list N places to the left")
    fmt.Printf("%v, %v -> %v\n", list1, 3, rotate(list1, 3))
    fmt.Printf("%v, %v -> %v\n", list2, 3, rotate(list2, 3))
    fmt.Printf("%v, %v -> %v\n", list3, -3, rotate(list3, -3))
    fmt.Printf("%v, %v -> %v\n\n", list3, 3, rotate(list3, 3))

    fmt.Println("#20 remove the K'th element from a list")
    resultList, removedElement := removeAt(list1, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n", list1, 5, resultList, removedElement)
    resultList, removedElement = removeAt(list2, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n", list2, 5, resultList, removedElement)
    resultList, removedElement = removeAt(list3, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n\n", list3, 5, resultList, removedElement)
}

func pack(list []interface{}) (packedList []interface{}) {
    var currentValue interface{}

    packedSublist := []interface{}{}

    for _, element := range list {
        if currentValue != element {
            currentValue = element

            if len(packedSublist) > 0 {
                packedList = append(packedList, packedSublist)
            }

            packedSublist = []interface{}{currentValue}
        } else {
            packedSublist = append(packedSublist, element)
        }
    }

    if len(packedSublist) > 0 {
        packedList = append(packedList, packedSublist)
    }

    return
}

func encodeModified(list []interface{}) (encodedList []interface{}) {
    packedList := pack(list)

    for _, packedSubList := range packedList {
        packedSubList := packedSubList.([]interface{})

        lengthOfSublist := len(packedSubList)

        if lengthOfSublist > 1 {
            encodedList = append(encodedList, [2]interface{}{lengthOfSublist, packedSubList[0]})
        } else {
            encodedList = append(encodedList, packedSubList[0])
        }
    }

    return
}

func decodeModified(encodedList []interface{}) (decodedList []interface{}) {
    for _, encodedElement := range encodedList {
        encodedPair, isPair := encodedElement.([2]interface{})

        if isPair {
            length := encodedPair[0].(int)

            for length > 0 {
                decodedList = append(decodedList, encodedPair[1])

                length--
            }

        } else {
            decodedList = append(decodedList, encodedElement)
        }
    }

    return
}

func encodeDirect(list []interface{}) (encodedList []interface{}) {
    var currentValue interface{}

    runLength := 1

    for _, element := range list {
        if currentValue != element {
            if runLength > 1 {
                encodedList = append(encodedList, [2]interface{}{runLength, currentValue})

                runLength = 1
            } else if currentValue != nil {
                encodedList = append(encodedList, currentValue)
            }

            currentValue = element
        } else {
            runLength++
        }
    }

    if runLength > 1 {
        encodedList = append(encodedList, [2]interface{}{runLength, currentValue})
    } else if currentValue != nil{
        encodedList = append(encodedList, currentValue)
    }

    return
}

func dupli(list []interface{}) (duplicatedList []interface{}) {
    for _, element := range list {
        duplicatedList = append(duplicatedList, element, element)
    }

    return
}

func repli(list []interface{}, numberOfReplications int) (replicatedList []interface{}) {
    times := 1

    for _, element := range list {
        for times <= numberOfReplications {
            replicatedList = append(replicatedList, element)

            times++
        }

        times = 1
    }

    return
}

func myDrop(list []interface{}, number int) (reducedList []interface{}) {
    count := 1

    for _, element := range list {
        if count == number {
            count = 1
        } else {
            reducedList = append(reducedList, element)

            count++
        }
    }

    return
}

func split(list []interface{}, lengthOfFirstPart int)  (firstPart, secondPart []interface{}) {
    for index := 0; index < len(list); index++ {
        if index < lengthOfFirstPart {
            firstPart = append(firstPart, list[index])
        } else {
            secondPart = append(secondPart, list[index])
        }
    }

    return
}

func slice(list []interface{}, lowerBound, higherBound int) (slicedList []interface{}) {
    for index := 0; index < len(list); index++ {
        if (index + 1) >= lowerBound && (index + 1) <= higherBound {
            slicedList = append(slicedList, list[index])
        }
    }

    return
}

func rotate(list []interface{}, positions int) (rotatedList []interface{}) {
    if positions == 0 || len(list) == 0 {
        return list
    }

    firstPart, secondPart := []interface{}{}, []interface{}{}

    if positions > 0 {
        firstPart, secondPart = split(list, positions)
    } else {
        firstPart, secondPart = split(list, (len(list) + positions))
    }

    rotatedList = append(append(rotatedList, secondPart...), firstPart...)

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
