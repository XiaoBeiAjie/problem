package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var in = bufio.NewReaderSize(os.Stdin, 1<<20)
var out = bufio.NewWriterSize(os.Stdout, 1<<20)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func solve() {
	var n int
	fmt.Fscan(in, &n)
	w := make([]int, 0)
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < n; i ++ {
		var op int
		fmt.Fscan(in, &op)
		switch op {
			case 1: {
				var x int
				fmt.Fscan(in, &x)
				w = append(w, x)
			}
			case 2: {
				if h.Len() > 0 {
					x := heap.Pop(h)
					fmt.Println(x)
				} else {
					fmt.Println(w[0])
					w = w[1:]
				}
			}
			case 3: {
				for _, x := range w {
					heap.Push(h, x)
				}
				w = make([]int, 0)
			}
		}
	}
}

func main() {
    defer out.Flush()

    var T int
    // fmt.Fscan(in, &T)
    T = 1
    for ; T > 0; T-- {
        solve()
    }
}
