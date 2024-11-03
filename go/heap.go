package main

import "fmt"

type Heap struct {
	array []int
}

func (h *Heap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (i * 2) + 1
}

func right(i int) int {
	return (i * 2) + 2
}

func (h *Heap) HeapifyUp(i int) {
	for h.array[parent(i)] < h.array[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

func (h *Heap) insert(val int) {
	h.array = append(h.array, val)
	h.HeapifyUp(len(h.array) - 1)
}

func (h *Heap) HeapifyDown(i int) {

	l, r := left(i), right(i)
	lastIndex := len(h.array) - 1

	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex { // parent has a left child only
			childToCompare = l
		} else if h.array[l] > h.array[r] { // left child > right chlid
			childToCompare = l
		} else { // right child > left  chlid
			childToCompare = r
		}

		if h.array[i] < h.array[childToCompare] {
			h.swap(i, childToCompare)
			i := childToCompare
			l, r = left(i), right(i)
		} else {
			return
		}
	}

}
func (h *Heap) Extract() int {
	n := len(h.array)
	extracted := h.array[0]
	if n == 0 {
		fmt.Println("Nothing can be extracted as length if heap is zero")
		return -1
	}

	h.array[0] = h.array[n-1]
	h.array = h.array[:n-1]
	h.HeapifyDown(0)

	return extracted
}

func main() {
	h := Heap{[]int{}}
	vars := []int{10, 27, 57, 67, 78, 25}
	for _, v := range vars {
		h.insert(v)
		fmt.Println(h)
	}
	fmt.Println("============")
	for i := 0; i < 3; i++ {
		h.Extract()
		fmt.Println(h)
	}

}
