package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type InputStrings string

func main() {
	fmt.Println(Do(os.Stdin))
}

func Do(reader io.Reader) string {
	in := bufio.NewReader(reader)
	var n, m int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)
	var s string
	fmt.Fscan(in, &s)
	rs := []rune(s)

	ss := make([][]rune, m)
	dc := make([]int, n)
	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(in, &c)
		c--
		dc[i] = c
		if len(ss[c]) == 0 {
			ss[c] = make([]rune, 0)
		}
		ss[c] = append(ss[c], rs[i])
	}
	// for i := 0; i < m; i++ {
	// 	fmt.Println(string(ss[i]))
	// }

	for i := 0; i < m; i++ {
		l := len(ss[i])
		if l <= 1 {
			continue
		}
		// fmt.Println(string(ss[i][l-1:]))
		// fmt.Println(string(ss[i][:l-1]))
		ss[i] = append(ss[i][l-1:], ss[i][:l-1]...)
	}
	// for i := 0; i < m; i++ {
	// 	fmt.Println(string(ss[i]))
	// }
	ret := make([]rune, n)
	for i := 0; i < n; i++ {
		ci := dc[i]
		v := ss[ci][0]
		if len(ss[ci]) > 1 {
			ss[ci] = ss[ci][1:]
		}
		ret[i] = v
	}

	return string(ret)

}
