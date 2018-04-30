/*
Copyright 2015 Olivier Mengué

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package shuffle allows to shuffle a slice using the sort.Interface.
//
// See also https://golang.org/pkg/math/rand/#Shuffle (go 1.10).
package shuffle

import "math/rand"

// Subset of sort.Interface
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Shuffle a slice that implements the shuffle.Interface (a subset of
// sort.Interface)
//
// Note: it may happen that the result satisfies sort.IsSorted(data)
//
// See also https://golang.org/pkg/math/rand/#Shuffle (go 1.10).
func Shuffle(data Interface) {
	n := data.Len()
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		j := rand.Intn(i + 1)
		data.Swap(i, j)
	}
}
