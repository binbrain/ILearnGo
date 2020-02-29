package bst

import (
    "testing"
)

func isBST (r *Node, min, max int) bool {
    if r == nil {
        return true
    }
    if r.value < min && r.value > max {
        return false
    }

    return isBST(r.left, min, r.value) && isBST(r.right, r.value, max)
}

func TestBST(t *testing.T) {
    t.Run("Basic BST test", func(t *testing.T) {
        tree := &Node{value:5}
        tree.Insert(6)
        tree.Insert(8)
        tree.Insert(13)
        tree.Insert(14)
        tree.Insert(7)
        tree.Insert(2)

        got := isBST(tree, 1, 15)

        if got != true {
            t.Errorf("got %t want given %t", got, true)
        }
    })
}

func TestFind(t *testing.T) {
    t.Run("BST Find test", func(t *testing.T) {
        tree := &Node{value:5}
        tree.Insert(6)
        tree.Insert(8)
        tree.Insert(13)
        tree.Insert(14)
        tree.Insert(7)
        tree.Insert(2)

        got := isBST(tree, 1, 15)

        if got == true {
            t.Errorf("got %t want given %t", got, true)
        }
    })
}
