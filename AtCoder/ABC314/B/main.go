package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type InputStrings string

func main() {
	fmt.Println(Do(os.Stdin))
}

func Do(reader io.Reader) string {
	in := bufio.NewReader(reader)
	var n int
	fmt.Fscan(in, &n)

	cs := make([]int, n)
	as := make([]map[int]struct{}, n)

	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(in, &c)
		cs[i] = c
		as[i] = make(map[int]struct{})
		for j := 0; j < c; j++ {
			var a int
			fmt.Fscan(in, &a)
			as[i][a] = struct{}{}
		}
	}
	var t int
	fmt.Fscan(in, &t)

	type tags struct {
		Number int
		Length int
	}
	min := 37
	r := []string{}
	for i := 0; i < n; i++ {
		_, ok := as[i][t]
		if ok {
			if cs[i] < min {
				r = []string{strconv.Itoa(i + 1)}
				min = cs[i]
			} else if cs[i] == min {
				r = append(r, strconv.Itoa(i+1))
			}
		}
	}
	// r := []tags{}
	// for i := 0; i < n; i++ {
	// 	_, ok := as[i][t]
	// 	if ok {
	// 		r = append(r, tags{Number: i, Length: cs[i]})
	// 	}
	// }
	// sort.Slice(r, func(i, j int) bool { return r[i].Length < r[j].Length })
	if len(r) == 0 {
		return "0"
	}
	str := strconv.Itoa(len(r)) + "\n"
	str += strings.Join(r, " ")
	return str

}
