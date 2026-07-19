package main

import (
	"bufio"
	. "fmt"
	"os"
)

var in = bufio.NewReaderSize(os.Stdin, 1<<20)
var out = bufio.NewWriterSize(os.Stdout, 1<<20)

type bit []int

func (tr bit) add(i, v int) {
	for ; i < len(tr); i += i & -i {
		tr[i] += v
	}
}

func (tr bit) pre(i int) (res int) {
	for ; i > 0; i -= i & -i {
		res += tr[i]
	}
	return res
}

func (tr bit) query(l, r int) (res int) {
	return tr.pre(r) - tr.pre(l-1)
}

func solve() {
	var n int
	const mod = 998244353
	Fscan(in, &n)
	inc := make(bit, n+1)
	dec := make(bit, n+1)
	sig := make(bit, n+1)
	for i := 1; i <= n; i++ {
		var x int
		Fscan(in, &x)
		sig.add(x, 1)
		inc.add(x, (inc.pre(x-1)+dec.pre(x-1)+1)%mod)
		dec.add(x, (((inc.query(x+1, n)+dec.query(x+1, n))%mod-sig.query(x+1, n))%mod+mod)%mod)
	}
	Println(dec.pre(n) % mod)
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
