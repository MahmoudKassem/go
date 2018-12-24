package main

import "fmt"

func main() {
    binaryTree1 := &binaryTree{
        value: 0,
        left: &binaryTree{
            value: 1,
            left:  nil,
            right: nil,
        },
        right: &binaryTree{
            value: 0,
            left: &binaryTree{
                value: 1,
                left: &binaryTree{
                    value: 1,
                    left:  nil,
                    right: nil,
                },
                right: &binaryTree{
                    value: 1,
                    left:  nil,
                    right: nil,
                },
            },
            right: &binaryTree{
                value: 0,
                left:  nil,
                right: nil,
            },
        },
    }

    printBinaryTree(binaryTree1, 0)
    fmt.Printf(" -> %v\n\n", countUnivalbinaryTrees(binaryTree1))

    binaryTree2 := &binaryTree{
        value: 1,
        left: &binaryTree{
            value: 1,
            left: &binaryTree{
                value: 1,
                left:  nil,
                right: nil,
            },
            right: &binaryTree{
                value: 1,
                left:  nil,
                right: nil,
            },
        },
        right: &binaryTree{
            value: 1,
            left: &binaryTree{
                value: 1,
                left:  nil,
                right: nil,
            },
            right: &binaryTree{
                value: 1,
                left:  nil,
                right: nil,
            },
        },
    }

    printBinaryTree(binaryTree2, 0)
    fmt.Printf(" -> %v\n\n", countUnivalbinaryTrees(binaryTree2))

    binaryTree3 := &binaryTree{
        value: 1,
        left: &binaryTree{
            value: 1,
            left: &binaryTree{
                value: 1,
                left:  nil,
                right: nil,
            },
            right: nil,
        },
        right: &binaryTree{
            value: 1,
            left: &binaryTree{
                value: 1,
                left:  nil,
                right: nil,
            },
            right: &binaryTree{
                value: 1,
                left:  nil,
                right: nil,
            },
        },
    }

    printBinaryTree(binaryTree3, 0)
    fmt.Printf(" -> %v\n\n", countUnivalbinaryTrees(binaryTree3))
}

type binaryTree struct {
    value interface{}
    left  *binaryTree
    right *binaryTree
}

func printTabs(numberOfTabs int)  {
    for count := 0; count < numberOfTabs; count++ {
        fmt.Printf("\t")
    }
}

func printBinaryTree(binaryTree *binaryTree, numberOfTabs int) {
    fmt.Printf("binaryTree {\n")

    numberOfTabs++

    printTabs(numberOfTabs)

    fmt.Printf("value: %v\n", binaryTree.value)

    printTabs(numberOfTabs)

    if binaryTree.left != nil {
        fmt.Printf("left: ")

        printBinaryTree(binaryTree.left, numberOfTabs + 1)

        fmt.Printf("\n")
    } else {
        fmt.Printf("left: nil\n")
    }

    printTabs(numberOfTabs)

    if binaryTree.right != nil {
        fmt.Printf("right: ")

        printBinaryTree(binaryTree.right, numberOfTabs + 1)

        fmt.Printf("\n")
    } else {
        fmt.Printf("right: nil\n")
    }

    numberOfTabs--

    printTabs(numberOfTabs)

    fmt.Printf("}")
}

func countUnivalbinaryTrees(binaryTree *binaryTree) (count int) {
    traverserMessages := traverser(binaryTree)

    for traverserMessage := range traverserMessages {
        count += traverserMessage
    }

    return
}

func traverser(binaryTree *binaryTree) <-chan int {
    traverserMessages := make(chan int)

    go func() {
        traverse(binaryTree, traverserMessages)

        defer close(traverserMessages)
    }()

    return traverserMessages
}

func traverse(binaryTree *binaryTree, traverserMessages chan int) {
    if binaryTree == nil {
        return
    }

    if binaryTree.left == nil && binaryTree.right == nil ||
       binaryTree.left != nil && binaryTree.right == nil ||
       binaryTree.left == nil && binaryTree.right != nil ||
       binaryTree.left != nil && binaryTree.right != nil && binaryTree.left.value == binaryTree.right.value {
        traverserMessages <- 1
    }

    traverse(binaryTree.left, traverserMessages)

    traverse(binaryTree.right, traverserMessages)
}
