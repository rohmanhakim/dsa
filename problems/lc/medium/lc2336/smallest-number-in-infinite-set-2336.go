package lc2336

import (
	"sort"
)

type SmallestInfiniteSet struct {
	set []int
}

func Constructor() SmallestInfiniteSet {
	set := make([]int, 1000)
	for i := range 1000 {
		set[i] = i + 1
	}
	return SmallestInfiniteSet{
		set: set,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	smallest := this.set[0]
	this.set = this.set[1:]
	return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	idx := sort.Search(len(this.set), func(i int) bool {
		return this.set[i] >= num
	})
	if this.set[idx] == num {
		return
	}
	this.set = append(this.set, 0)
	copy(this.set[idx+1:], this.set[idx:])
	this.set[idx] = num
}

func SmallestNumberInInfiniteSet() {

	obj := Constructor()

	// obj.AddBack(2)
	// obj.PopSmallest() // 1
	// obj.PopSmallest() // 2
	// obj.PopSmallest() // 3
	// obj.AddBack(1)
	// obj.PopSmallest() // 1
	// obj.PopSmallest() // 4
	// obj.PopSmallest() // 5

	// obj.AddBack(2)
	// obj.PopSmallest() // 1
	// obj.PopSmallest() // 2
	// obj.AddBack(3)
	// obj.PopSmallest() // 3
	// obj.AddBack(2)
	// obj.PopSmallest() // 2
	// obj.PopSmallest() // 4

	// obj.PopSmallest() // 1
	// obj.AddBack(1)
	// obj.PopSmallest() // 1
	// obj.PopSmallest() // 2
	// obj.PopSmallest() // 3
	// obj.AddBack(2)
	// obj.AddBack(3)
	// obj.PopSmallest() // 2
	// obj.PopSmallest() // 3
	//
	obj.AddBack(2)
	obj.PopSmallest() // 1
	obj.PopSmallest() // 2
	obj.PopSmallest() // 3
	obj.AddBack(1)
	obj.PopSmallest() // 1
	obj.PopSmallest() // 4
	obj.PopSmallest() // 5
}
