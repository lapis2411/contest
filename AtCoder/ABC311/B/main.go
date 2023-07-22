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
	cond := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(cond[0])
	d, _ := strconv.Atoi(cond[1])
	plan := make([][]rune, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		plan[i] = []rune(sc.Text())
	}
	max := 0
	cnt := 0
	for i := 0; i < d; i++ {
		can := true
		for j := 0; j < n; j++ {
			if plan[j][i] == 'x' {
				can = false
				break
			}
		}
		if can {
			cnt++
			if cnt > max {
				max = cnt
			}
		} else {
			cnt = 0
		}
	}
	return strconv.Itoa(max)
}
