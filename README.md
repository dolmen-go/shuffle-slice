# github.com/dolmen-go/shuffle-slice

Shuffle slices as you sort them.

Example:
```go
package main

import (
    "github.com/dolmen-go/shuffle-slice/shuffle"
    "fmt"
    "sort"
)

func main() {
    slice := []string{"foo", "bar", "baz"}
    shuffle.Shuffle(sort.StringSlice(slice))
    fmt.Println(slice)
}
```


Note: since go1.10, the package `math/rand` has a function [`Shuffle`](https://golang.org/pkg/math/rand/#Shuffle):

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    slice := []string{"foo", "bar", "baz"}
    rand.Shuffle(len(slice), func(i, j int) {
    	slice[i], slice[j] = slice[j], slice[i]
    })
    fmt.Println(slice)
}
```
