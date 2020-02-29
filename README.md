# ILearnGo

This is my personal playground for learning Go and brushing up on my data structures and algorithms at the same time. Its probably useful to nobody but me.

## Setup

Requires docker. To build run...

```
make build
```

## Test

Testing can be done at the module level and submodule level. Also simple to run specific tests.

```
make test app=fib
```

To run a specific test, for example a BST test called *Find* contained in the submodule *bst*.

```
make t1 t=Find app=binarytrees/bst
```
