package knslice

import "fmt"

type (
	node struct {
		val      int
		slice    int
		nextElem int
	}
	minheap []node
)

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h minheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x interface{}) {
	*h = append(*h, x.(node))
}

func (h *minheap) Pop() interface{} {
	old := *h
	oldLen := len(old)
	elem := old[oldLen-1]
	*h = old[0 : oldLen-1]
	return elem
}

func (h minheap) print() {
	fmt.Printf("%+v\n", h)
}
