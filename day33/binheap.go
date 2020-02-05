package day33

// BinaryHeap represents min- or max-heap
type BinaryHeap struct {
	arr []int
	cmp func(int, int) bool
}

// NewMinHeap builds initial min-heap
func NewMinHeap() BinaryHeap {
	return BinaryHeap{[]int{}, minCmp}
}

// NewMaxHeap builds initial max-heap
func NewMaxHeap() BinaryHeap {
	return BinaryHeap{[]int{}, maxCmp}
}

// Push inserts new key into binary heap
func (h *BinaryHeap) Push(v int) {
	idx := len(h.arr)
	arr := append(h.arr, v)

	for {
		parentIdx := parent(idx)
		if idx == 0 || !h.cmp(arr[idx], arr[parentIdx]) {
			break
		}
		arr[idx], arr[parentIdx] = arr[parentIdx], arr[idx]
		idx = parentIdx
	}

	h.arr = arr
}

// Size returns size of a binaary tree
func (h *BinaryHeap) Size() int {
	return len(h.arr)
}

// Get returns root of a binary heap
func (h *BinaryHeap) Get() int {
	return h.arr[0]
}

// Pop removes root from a binary heap
func (h *BinaryHeap) Pop() int {
	min := h.arr[0]
	end := len(h.arr) - 1

	h.arr[0] = h.arr[end]
	updated := h.arr[:end]
	heapify(updated, 0, h.cmp)

	h.arr = updated

	return min
}

func minCmp(fst int, snd int) bool {
	return fst < snd
}

func maxCmp(fst int, snd int) bool {
	return fst > snd
}

func heapify(in []int, idx int, cmp func(int, int) bool) {
	smallest := idx

	l := left(idx)
	if l < len(in) && cmp(in[l], in[smallest]) {
		smallest = l
	}
	r := right(idx)
	if r < len(in) && cmp(in[r], in[smallest]) {
		smallest = r
	}
	if smallest != idx {
		in[idx], in[smallest] = in[smallest], in[idx]
		heapify(in, smallest, cmp)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}
