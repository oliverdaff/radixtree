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







## Tests
The tests can be invoked with `go test`

## License
MIT © Oliver Daff