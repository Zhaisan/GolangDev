package main

import "fmt"
import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    var walk func(t *tree.Tree)
    walk = func (t *tree.Tree) {
        if (t == nil) {
            return
        }
        walk(t.Left)
        ch <- t.Value
        walk(t.Right)
    }
    walk(t)
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
	ch2 := make(chan int)

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    for {
        node1, check1 := <- ch1
        node2, check2 := <- ch2

        if node1 != node2 || check1 != check2 {
            return false
        }

        if !check1 {
            break
        }
    }

    return true
}

func Tester() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

func main() {
    Tester()
}