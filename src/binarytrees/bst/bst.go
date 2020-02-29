package bst

import "fmt"


type Node struct {
    value int
    right, left *Node
}


func (r *Node) FindMin () {
    // returm min
}

func (r *Node) FindMax () {
    // return max
}

func (r *Node) Insert (value int) {
    // Sorted BinarySearchTree Insert (unbalanced Insert)
    if r.value > value {
        if r.left == nil {
            r.left = &Node{value:value}
        } else {
            r.left.Insert(value)
        }
    } else {
        if r.right == nil {
            r.right = &Node{value:value}
        } else {
            r.right.Insert(value)
        }
    }
}

func (r *Node) Find (value int) *Node {
    if r.value == value {
        return r.left.Find(value)
    }
    if r.value < value {
        return r.right.Find(value)
    }
    if r.value > value {
        return r.left.Find(value)
    }

    return nil
}

func (r *Node) Delete (value int) {
}

func depthFirstPrint (n *Node) {
    fmt.Println(n.value)
    if n.left != nil {
        depthFirstPrint(n.left)
    }
    if n.right != nil {
        depthFirstPrint(n.right)
    }
}
