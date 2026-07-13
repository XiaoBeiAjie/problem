package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReaderSize(os.Stdin, 1<<20)
var out = bufio.NewWriterSize(os.Stdout, 1<<20)

func solve() {
	var n, p, q, r int
	fmt.Fscan(in, &n, &p, &q, &r)

	w := make([]int, n + 1)
	for i := 1; i <= n; i ++  {
		fmt.Fscan(in, &w[i])
		w[i] = w[i] + w[i - 1]
	}

	o:
	for i, v := range w {
		for _, t := range []int{p, q, r} {
			v += t
			l, r := i, n
			check := func(mid int) bool {
				return w[mid] <= v
			}
			for l < r {
				mid := (l + r + 1) / 2
				if check(mid) {
					l = mid
				} else {
					r = mid - 1
				}
			}
			if w[l] != v {
				continue o
			}
		}
		fmt.Fprintln(out, "Yes")
		return 
	}
	fmt.Fprintln(out, "No")
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
