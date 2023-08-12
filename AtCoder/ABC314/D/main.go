package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type InputStrings string

func main() {
	fmt.Fprint(os.Stdout, Do(os.Stdin))
}

func Do(reader io.Reader) string {
	in := bufio.NewReader(reader)
	var n, q int
	var s string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &q)
	rs := []rune(s)
	buf := map[int]rune{} // i文字目はruneという情報　t が　2, 3なら書き込みしてクリアする
	lis := 0              // 最後の大文字化、小文字化命令
	for i := 0; i < q; i++ {
		var t, x int
		var c byte
		fmt.Fscan(in, &t)
		fmt.Fscan(in, &x)
		fmt.Fscan(in, &c)

		if t == 1 {
			r := []rune(c)
			buf[x-1] = r[0]
			continue
		}
		lis = t
		for i, v := range buf {
			rs[i] = v
		}
		buf = map[int]rune{}
	}

	r := string(rs)
	if lis != 0 {
		if lis == 2 {
			r = strings.ToLower(r)
		} else {
			r = strings.ToUpper(r)
		}
		t := []rune(r)
		for i, v := range buf {
			t[i] = v
		}
		r = string(t)
	}

	return r

}
