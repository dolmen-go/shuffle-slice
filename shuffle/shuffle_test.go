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

package shuffle

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// shuffle.Interface is a subset of sort.Interface
var _ Interface = sort.Interface(nil)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func fillAZ(az []string) {
	b := []byte{65}
	for i := 0; i < 26; i++ {
		az[i] = string(b)
		b[0]++
	}
}

func makeAZ() []string {
	var az [26]string
	fillAZ(az[:])
	return az[:]
}

func TestShuffle(t *testing.T) {
	const checks = 10
	shuffled := 0
	for i := 0; i < checks; i++ {
		az := makeAZ()
		t.Logf("Before: %v", az)
		Shuffle(sort.StringSlice(az))
		t.Logf(" After: %v", az)
		if !sort.IsSorted(sort.StringSlice(az)) {
			shuffled++
		}
	}
	// We can not ensure that the suffled result is not the same as the
	// original
	// But <checks> times in a row is really suspicious
	if shuffled == 0 {
		t.Fatal("Shuffle never shuffled!")
	}
}

func ExampleShuffle() {
	af := sort.StringSlice([]string{"A", "B", "C", "D", "E", "F"})
	fmt.Println("Original: %v", af)

	Shuffle(af)
	fmt.Println("Shuffled: %v", af)

	sort.Sort(af)
	fmt.Println("  Sorted: %v", af)
}
