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
	in := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(in[0])
	m, _ := strconv.Atoi(in[1])
	weaker := make([]int, n)
	stronger := make([]int, n)
	for i := 0; i < m; i++ {
		sc.Scan()
		ab := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(ab[0])
		b, _ := strconv.Atoi(ab[1])
		inda := a - 1
		stronger[inda] = 1
		indb := b - 1
		weaker[indb] = 1
	}
	conds := []int{}
	for i, v := range weaker {
		if v == 0 {
			conds = append(conds, i)
		}
	}
	if len(conds) != 1 {
		return "-1"
	}
	cond := conds[0]
	if stronger[cond] == 0 {
		return "-1"
	}
	res := cond + 1
	return strconv.Itoa(res)

}
