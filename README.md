# go-shuffle-slices

Shuffle slices as you sort them.

Example:
```go
package main

import (
    "github.com/dolmen/go-shuffle-slice/shuffle"
    "fmt"
    "sort"
)

func main() {
    slice := []string{"foo", "bar", "baz"}
    shuffle.Shuffle(sort.StringSlice(slice))
    fmt.Println(slice)
}
```
