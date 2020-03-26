package knslice

import (
	"container/heap"
	"math"
	"sort"
)

const infinite = math.MaxInt64

func merge(k [][]int, n int) []int {
	var nodes = make([]node, len(k))
	for i, sl := range k {
		nodes[i] = node{
			val:      sl[0],
			slice:    i,
			nextElem: 1,
		}
	}
	h := minheap(nodes)
	heap.Init(&h)

	out := make([]int, len(k)*n)
	for i := 0; i < len(k)*n; i++ {
		root := heap.Pop(&h).(node)
		out[i] = root.val

		if root.nextElem < n {
			root.val = k[root.slice][root.nextElem]
			root.nextElem++
		} else {
			root.val = infinite
		}

		heap.Push(&h, root)
		heap.Fix(&h, 0)
	}

	return out
}

func merge_dumb(k [][]int, n int) []int {
	var out = make([]int, 0, len(k)*n)
	for _, sl := range k {
		for _, v := range sl {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return i > j })
	return out
}
