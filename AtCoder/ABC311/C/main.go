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
	sc := bufio.NewScanner(reader)
	sc.Scan()
	sc.Scan()
	a := strings.Split(sc.Text(), " ")
	table := make(map[int]int)
	next := 0
	cnt := 1
	seq := []string{}
	for {
		crnt := next
		seq = append(seq, strconv.Itoa(crnt+1))
		table[crnt] = cnt
		cnt++

		next, _ = strconv.Atoi(a[crnt])
		next--
		// 既出の数
		if v, ok := table[next]; ok {
			len := cnt - v
			return strconv.Itoa(len) + "\n" + strings.Join(seq[v-1:cnt-1], " ")
		}
	}
}
