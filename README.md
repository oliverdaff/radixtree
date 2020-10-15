# radixtree

A data structure that represents a space-optimized trie
(prefix tree) in which each node that is the only child is merged with its parent.

This implementation will store strings with values, the trie allow nil values to be stored but not nil.

# API

__Create new__

Create a new RadixTree

```go
trie = trie.NewRadixTree()
```

__Put key value__

Put key value into the trie

```go
trie = trie.NewRadixTree()
trie.Put("abc", 1)
```

__Get value__

Retrieve the value stored in the trie for the key else else nil.

```go
trie = trie.NewRadixTree()
trie.Put("abc", 1)
result, _ = trie.Get("abc")
```

__Delete value__

Delete the value stored in the trie for the key, returns true
if a key was found and the value was deleted else returns false.

```go
trie = trie.NewRadixTree()
trie.Put("abc", 1)
deleted, _ = trie.Delete("abc")
```





## Tests
The tests can be invoked with `go test`

## License
MIT © Oliver Daff