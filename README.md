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

__Check if key stored with value__

Return true if key was found in the trie.

```go
trie = trie.NewRadixTree()
trie.Put("abc", 1)
deleted, _ = trie.Contains("abc")
```

__Longest Prefix__

Find the key with longest prefix in the trie for the query string.

```go
trie = trie.NewRadixTree()
trie.Put("abc", 1)
key, _ = trie.LongestPrefixOf("abcdef")
```

__Keys For Prefix__

Return all the keys in the trie that starts with this prefix.

```go
trie = trie.NewRadixTree()
trie.Put("www.test.com", 1)
trie.Put("www.example.com", 1)
key_channel = trie.KeyWithPrefix("www.")
```

__All Keys__

Return all the keys in the trie.

```go
trie = trie.NewRadixTree()
trie.Put("www.test.com", 1)
trie.Put("www.example.com", 1)
key_channel = trie.Keys()
```


__All Key Value Pairs__

Return all the key value pairs.
```go
trie = trie.NewRadixTree()
trie.Put("www.test.com", 1)
trie.Put("www.example.com", 1)
key_channel = trie.Items()
```



## Tests
The tests can be invoked with `go test`

## License
MIT © Oliver Daff