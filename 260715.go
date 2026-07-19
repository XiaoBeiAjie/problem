package main

import (
	"bufio"
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

const mod = 998244353

func qmi(a, b int64) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

const N int = 2e5 + 10

var fac [N + 1]int64
var ifac [N + 1]int64

func init() {
	fac[0] = 1
	for i := 1; i <= N; i++ {
		fac[i] = fac[i-1] * int64(i) % mod
	}
	ifac[N] = qmi(fac[N], mod-2)
	for i := N; i > 0; i-- {
		ifac[i-1] = ifac[i] * int64(i) % mod
	}
}

func C(n, k int) int64 {
	return fac[n] * ifac[k] % mod * ifac[n-k] % mod
}

func solve() {
	var n, k int
	fmt.Fscan(in, &n, &k)
	w := make([]int64, n+1)
	a := make([]int64, n+1)
	ans := int64(1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		w[i] = a[i] % int64(k)
	}
	for i := 1; i <= n; i++ {
		res := int64(1)
		len := int64(1)
		div := int64(1)
		if w[i]*int64(2)%int64(k) == 0 {
			cnt := map[int64]int64{a[i]: 1}
			for j := i + 1; j <= n && w[j] == w[i]; j++ {
				len++
				res = res * len % mod
				cnt[a[j]]++
				div = div * cnt[a[j]] % mod
			}
			ans = ans * res % mod * qmi(div, mod-2) % mod
			i += int(len - 1)
		} else {
			cnt := int64(1)
			res := int64(1)
			len := int64(1)
			for j := i + 1; j <= n && (w[j] == w[i] || w[j] == (int64(k)-w[i])%int64(k)); j++ {
				len++
				res = res * len % mod
				if w[j] == w[i] {
					cnt++
				}
			}
			ans = ans * res % mod * ifac[cnt] % mod * ifac[len-cnt] % mod
			i += int(len - 1)
		}
	}
	fmt.Fprintln(out, ans)
}

func main() {
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)
	// T = 1
	for ; T > 0; T-- {
		solve()
	}
}
