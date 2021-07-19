package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
    list11 := []int{2, 0, 2}
    list12 := []interface{}{[]interface{}{1, 2, 3}, []interface{}{4, 5}, []interface{}{6, 7, 8}, []interface{}{4, 5}, []interface{}{9, 10, 11, 12}, []interface{}{13, 14}, []interface{}{15}}
    list13 := []interface{}{[]interface{}{"a", "b", "c"}, []interface{}{"d", "e"}, []interface{}{"f", "g", "h"}, []interface{}{"d", "e"}, []interface{}{"i", "j", "k", "l"}, []interface{}{"m", "n"}, []interface{}{"o"}}

    fmt.Println("#1 last element of a list")
    fmt.Printf("%v -> %v\n", list1, last(list1))
    fmt.Printf("%v -> %v\n", list2, last(list2))
    fmt.Printf("%v -> %v\n\n", list3, last(list3))

    fmt.Println("#2 last but one element of a list")
    fmt.Printf("%v -> %v\n", list1, lastButOne(list1))
    fmt.Printf("%v -> %v\n", list2, lastButOne(list2))
    fmt.Printf("%v -> %v\n\n", list3, lastButOne(list3))

    fmt.Println("#3 k'th element of a list")
    fmt.Printf("%v, %v -> %v\n", list1, 2, elementAt(list1, 2))
    fmt.Printf("%v, %v -> %v\n", list2, 2, elementAt(list2, 2))
    fmt.Printf("%v, %v -> %v\n\n", list3, 2, elementAt(list3, 2))

    fmt.Println("#4 number of elements in a list")
    fmt.Printf("%v -> %v\n", list1, length(list1))
    fmt.Printf("%v -> %v\n", list2, length(list2))
    fmt.Printf("%v -> %v\n\n", list3, length(list3))

    fmt.Println("#5 reverse a list")
    fmt.Printf("%v -> %v\n", list1, reverse(list1))
    fmt.Printf("%v -> %v\n", list2, reverse(list2))
    fmt.Printf("%v -> %v\n\n", list3, reverse(list3))

    fmt.Println("#6 is list a palindrome")
    fmt.Printf("%v -> %v\n", list1, isPalindrom(list1))
    fmt.Printf("%v -> %v\n", list2, isPalindrom(list2))
    fmt.Printf("%v -> %v\n", list3, isPalindrom(list3))
    fmt.Printf("%v -> %v\n", list4, isPalindrom(list4))
    fmt.Printf("%v -> %v\n", list5, isPalindrom(list5))
    fmt.Printf("%v -> %v\n\n", list6, isPalindrom(list6))

    fmt.Println("#7 flatten a nested list")
    fmt.Printf("%v -> %v\n", list1, flatten(list1))
    fmt.Printf("%v -> %v\n", list2, flatten(list2))
    fmt.Printf("%v -> %v\n", list3, flatten(list3))
    fmt.Printf("%v -> %v\n", list7, flatten(list7))
    fmt.Printf("%v -> %v\n\n", list8, flatten(list8))

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

    fmt.Println("#11 run-length encoding of a list(modified)")
    fmt.Printf("%v -> %v\n", list3, encodeModified(list3))
    fmt.Printf("%v -> %v\n", list9, encodeModified(list9))
    fmt.Printf("%v -> %v\n\n", list10, encodeModified(list10))

    fmt.Println("#12 decode a run-length encoded list")
    fmt.Printf("%v -> %v\n", encodeModified(list3), decodeModified(encodeModified(list3)))
    fmt.Printf("%v -> %v\n", encodeModified(list9), decodeModified(encodeModified(list9)))
    fmt.Printf("%v -> %v\n\n", encodeModified(list10), decodeModified(encodeModified(list10)))

    fmt.Println("#13 run-length encoding of a list(direct solution)")
    fmt.Printf("%v -> %v\n", list3, encodeDirect(list3))
    fmt.Printf("%v -> %v\n", list9, encodeDirect(list9))
    fmt.Printf("%v -> %v\n\n", list10, encodeDirect(list10))

    fmt.Println("#14 duplicate the elements of a list")
    fmt.Printf("%v -> %v\n", list3, dupli(list3))
    fmt.Printf("%v -> %v\n", list9, dupli(list9))
    fmt.Printf("%v -> %v\n\n", list10, dupli(list10))

    fmt.Println("#15 replicate the elements of a list a given number of times")
    fmt.Printf("%v, %v -> %v\n", list3, 3, repli(list3, 3))
    fmt.Printf("%v, %v -> %v\n", list9, 3, repli(list9, 3))
    fmt.Printf("%v, %v -> %v\n\n", list10, 3, repli(list10, 3))

    fmt.Println("#16 drop every N'th element from a list")
    fmt.Printf("%v, %v -> %v\n", list3, 5, drop(list3, 5))
    fmt.Printf("%v, %v -> %v\n", list9, 5, drop(list9, 5))
    fmt.Printf("%v, %v -> %v\n\n", list10, 5, drop(list10, 5))

    fmt.Println("#17 split a list into two parts; the length of the first part is given")
    firstPart, secondPart := split(list3, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n", list3, 5, firstPart, secondPart)
    firstPart, secondPart = split(list9, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n", list9, 5, firstPart, secondPart)
    firstPart, secondPart = split(list10, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n\n", list10, 5, firstPart, secondPart)

    fmt.Println("#18 extract a slice from a list")
    fmt.Printf("%v, %v, %v -> %v\n", list3, 3, 8, slice(list3, 3, 8))
    fmt.Printf("%v, %v, %v -> %v\n", list9, 3, 8, slice(list9, 3, 8))
    fmt.Printf("%v, %v, %v -> %v\n\n", list10, 3, 8, slice(list10, 3, 8))

    fmt.Println("#19 rotate a list N places to the left")
    fmt.Printf("%v, %v -> %v\n", list3, 3, rotate(list3, 3))
    fmt.Printf("%v, %v -> %v\n", list9, 3, rotate(list9, 3))
    fmt.Printf("%v, %v -> %v\n", list10, -3, rotate(list10, -3))
    fmt.Printf("%v, %v -> %v\n\n", list10, 3, rotate(list10, 3))

    fmt.Println("#20 remove the K'th element from a list")
    resultList, removedElement := removeAt(list3, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n", list3, 5, resultList, removedElement)
    resultList, removedElement = removeAt(list9, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n", list9, 5, resultList, removedElement)
    resultList, removedElement = removeAt(list10, 5)
    fmt.Printf("%v, %v -> (%v, %v)\n\n", list10, 5, resultList, removedElement)

    fmt.Println("#21 insert an element at a given position into a list")
    fmt.Printf("%v, %v, %v -> %v\n", list1, 5, 3, insertAt(list1, 5, 3))
    fmt.Printf("%v, %v, %v -> %v\n", list2, "e", 3, insertAt(list2, "e", 3))
    fmt.Printf("%v, %v, %v -> %v\n\n", list3, 5, 3, insertAt(list3, 5, 3))

    fmt.Println("#22 create a list containing all integers within a given range")
    fmt.Printf("%v, %v -> %v\n", 3, 8, range_(3, 8))
    fmt.Printf("%v, %v -> %v\n", -8, -3, range_(-8, -3))
    fmt.Printf("%v, %v -> %v\n\n", 8, 3, range_(8, 3))

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

    fmt.Println("#26 generate the combinations of K distinct objects chosen from the N elements of a list")
    fmt.Printf("%v, %v -> %v\n", list1, 3, combinations(list1, 3))
    fmt.Printf("%v, %v -> %v\n", list2, 3, combinations(list2, 3))
    fmt.Printf("%v, %v -> %v\n", list3, 3, combinations(list3, 3))
    fmt.Printf("%v, %v -> %v\n\n", list1, 0, combinations(list1, 0))

    fmt.Println("#27 group the elements of a set into disjoint subsets")
    fmt.Printf("%v, %v -> %v\n", list1, list11, groups(list1, list11))
    fmt.Printf("%v, %v -> %v\n", list2, list11, groups(list2, list11))
    fmt.Printf("%v, %v -> %v\n\n", list3, list11, groups(list3, list11))

    fmt.Println("#28 sorting a list of lists according to length of sublists")
    fmt.Printf("%v -> %v\n", list12, lsort(list12))
    fmt.Printf("%v -> %v\n", list13, lsort(list13))
    fmt.Printf("%v -> %v\n\n", list3, lsort(list3))
    fmt.Printf("%v -> %v\n", list12, lfsort(list12))
    fmt.Printf("%v -> %v\n", list13, lfsort(list13))
    fmt.Printf("%v -> %v\n", list3, lfsort(list3))
}

func last(list []interface{}) (last interface{}) {
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

func length(list []interface{}) (length int) {
    for index := range list {
        length += index - (index - 1)
    }

    return
}

func reverse(list []interface{}) (reversed []interface{}) {
    for index := length(list) - 1; index >= 0; index-- {
        reversed = append(reversed, list[index])
    }

    return
}

func isPalindrom(list []interface{}) (isPalindrom bool) {
    reversed := reverse(list)
    for index := range list {
        if list[index] != reversed[index] {
            return
        }
    }

    return true
}

func flatten(nested []interface{}) (flattened []interface{}) {
    for _, head := range nested {
        if list, isList := head.([]interface{}); isList {
            flattened = append(flattened, flatten(list)...)
        } else {
            flattened = append(flattened, head)
        }
    }

    return
}

func compress(list []interface{}) (compressed []interface{}) {
    var current interface{}
    for _, head := range list {
        if current != head {
            compressed = append(compressed, head)
            current = head
        }
    }

    return
}

func pack(list []interface{}) (packed []interface{}) {
    subList := []interface{}{}
    var current interface{}
    for _, head := range list {
        if head != current {
            if length(subList) > 0 {
                packed = append(packed, subList)
            }

            subList = []interface{}{head}
            current = head
        } else {
            subList = append(subList, current)
        }
    }

    if length(subList) > 0 {
        packed = append(packed, subList)
    }

    return
}

func encode(list []interface{}) (encoded []interface{}) {
    packed := pack(list)
    for _, head := range packed {
        subList := head.([]interface{})
        encoded = append(encoded, []interface{}{length(subList), subList[0]})
    }

    return
}

func encodeModified(list []interface{}) (encoded []interface{}) {
    packed := pack(list)
    for _, head := range packed {
        subList := head.([]interface{})
        subListLength := length(subList)
        if subListLength > 1 {
            encoded = append(encoded, [2]interface{}{subListLength, subList[0]})
        } else {
            encoded = append(encoded, subList[0])
        }
    }

    return
}

func decodeModified(encoded []interface{}) (decoded []interface{}) {
    for _, head := range encoded {
        if pair, isPair := head.([2]interface{}); isPair {
            for length := pair[0].(int); length > 0; length-- {
                decoded = append(decoded, pair[1])
            }
        } else {
            decoded = append(decoded, head)
        }
    }

    return
}

func encodeDirect(list []interface{}) (encoded []interface{}) {
    var current interface{}
    replications := 1
    for _, head := range list {
        if current != head {
            if replications > 1 {
                encoded = append(encoded, [2]interface{}{replications, current})
                replications = 1
            } else if current != nil {
                encoded = append(encoded, current)
            }

            current = head
        } else {
            replications++
        }
    }

    if replications > 1 {
        encoded = append(encoded, [2]interface{}{replications, current})
    } else if current != nil{
        encoded = append(encoded, current)
    }

    return
}

func dupli(list []interface{}) (duplicated []interface{}) {
    for _, head := range list {
        duplicated = append(duplicated, head, head)
    }

    return
}

func repli(list []interface{}, replications int) (replicated []interface{}) {
    for _, head := range list {
        times := 1
        for times <= replications {
            replicated = append(replicated, head)
            times++
        }
    }

    return
}

func drop(list []interface{}, frequency int) (reduced []interface{}) {
    count := 1
    for _, head := range list {
        if count == frequency {
            count = 1
        } else {
            reduced = append(reduced, head)
            count++
        }
    }

    return
}

func split(list []interface{}, position int) (left, right []interface{}) {
    index := 1
    for _, head := range list {
        if index <= position {
            left = append(left, head)
        } else {
            right = append(right, head)
        }

        index++
    }

    return
}

func slice(list []interface{}, start, end int) (slice []interface{}) {
    index := 1
    for _, head := range list {
        if index > end {
            return
        }

        if index >= start && index <= end {
            slice = append(slice, head)
        }

        index++
    }

    return
}

func rotate(list []interface{}, positions int) (rotated []interface{}) {
    if positions == 0 || length(list) == 0 {
        return list
    }

    left, right := []interface{}{}, []interface{}{}
    if positions > 0 {
        left, right = split(list, positions)
    } else {
        left, right = split(list, (length(list) + positions))
    }

    rotated = append(append(rotated, right...), left...)
    return
}

func removeAt(list []interface{}, position int) (reduced []interface{}, removed interface{}) {
    index := 1
    for _, head := range list {
        if index == position {
            removed = head
        } else {
            reduced = append(reduced, head)
        }

        index++
    }

    return
}

func insertAt(list []interface{}, value interface{}, position int) (increased []interface{}) {
    index := 1
    for _, head := range list {
        if index == position {
            increased = append(increased, value, head)
        } else {
            increased = append(increased, head)
        }

        index++
    }

    return
}

func range_(minimum, maximum int) (range_ []int) {
    for minimum <= maximum {
        range_ = append(range_, minimum)
        minimum++
    }

    return
}

func randomSelection(list []interface{}, draws int) (selection []interface{}) {
    var randomElement interface{}
    randomNumberGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
    length := length(list)
    for draws > 0 && length > 0 {
        randomNumber := randomNumberGenerator.Intn(length) + 1
        list, randomElement = removeAt(list, randomNumber)
        selection = append(selection, randomElement)
        length--
        draws--
    }

    return
}

func convertToList(intList []int) (list []interface{}) {
    for _, head := range intList {
        list = append(list, head)
    }

    return
}

func lotto(draws, maximum int) (lotto []int) {
    if maximum < draws || maximum <= 0 {
        return
    }

    range_ := range_(1, maximum)
    selection := randomSelection(convertToList(range_), draws)
    for _, head := range selection {
        lotto = append(lotto, head.(int))
    }

    return
}

func randomPermutation(list []interface{}) (permutation []interface{}) {
    return randomSelection(list, length(list))
}

func combinations(list []interface{}, draws int) (combinations_ [][]interface{}) {
    if draws <= 0 {
        combinations_ = append(combinations_, []interface{}{})
        return
    }

    if length(list) == 0 {
        return
    }

    head := list[0]
    tail := list[1:]
    withHead := [][]interface{}{}
    for _, combination := range combinations(tail, (draws - 1)) {
        combination = append([]interface{}{head}, combination...)
        withHead = append(withHead, combination)
    }

    withoutHead := combinations(tail, draws)
    combinations_ = append(withHead, withoutHead...)
    return
}

func isElement(list []interface{}, element interface{}) (isElement bool) {
    for _, head := range list {
        if element == head {
            return true
        }
    }

    return
}

func difference(list1, list2 []interface{}) (difference []interface{}) {
    for _, head := range list1 {
        if(!isElement(list2, head)) {
            difference = append(difference, head)
        }
    }

    return
}

func groups(list []interface{}, sizes []int) (groups_ [][][]interface{}){
    if length(convertToList(sizes)) == 0 {
        groups_ = append(groups_, [][]interface{}{})
        return
    }

    head := sizes[0]
    tail := sizes[1:]
    for _, combination := range combinations(list, head) {
        for _, group := range groups(difference(list, combination), tail) {
            groups_ = append(groups_, append([][]interface{}{combination}, group...))
        }
    }

    return
}

func pivot(sortBy func(a, b interface{}) bool, pivot interface{}, list []interface{}) (less, greaterOrEqual []interface{}) {
    for _, head := range list {
        if sortBy(head, pivot) {
            less = append(less, head)
        } else {
            greaterOrEqual = append(greaterOrEqual, head)
        }
    }

    return
}

func sortBy(sortBy_ func(a, b interface{}) bool, list []interface{}) (sorted []interface{}) {
    if length(list) == 0 {
        return
    }

    head := list[0]
    tail := list[1:]
    less, greaterOrEqual := pivot(sortBy_, head, tail)
    sortedLess := sortBy(sortBy_, less)
    sortedGreaterOrEqual := sortBy(sortBy_, greaterOrEqual)
    sorted = append(append(sortedLess, head), sortedGreaterOrEqual...)
    return
}

func lsort(list []interface{}) (sorted []interface{}) {
    sorted = sortBy(func(a, b interface{}) bool {
        listA := a.([]interface{})
        listB := b.([]interface{})
        return length(listA) < length(listB)
    }, list)

    return
}

func group(groupBy_ func(a, b interface{}) bool, key interface{}, list []interface{}) (group, rest []interface{}) {
    group = append(group, key)
    for _, head := range list {
        if groupBy_(head, key) {
            group = append(group, head)
        } else {
            rest = append(rest, head)
        }
    }

    return
}

func groupBy(groupBy_ func(a, b interface{}) bool, list []interface{}) (grouped []interface{}) {
    if length(list) == 0 {
        return
    }

    head := list[0]
    tail := list[1:]
    group, rest := group(groupBy_, head, tail)
    grouped = append(groupBy(groupBy_, rest), group)
    return
}

func lfsort(list []interface{}) (sorted []interface{}) {
    grouped := groupBy(func(a, b interface{}) bool {
        listA := a.([]interface{})
        listB := b.([]interface{})
        return length(listA) == length(listB)
    }, list)

    sortedGroups := lsort(reverse(grouped))
    for _, sortedGroup := range sortedGroups {
        groupMembers := sortedGroup.([]interface{})
        for _, groupMember := range groupMembers {
            sorted = append(sorted, groupMember)
        }
    }

    return
}
