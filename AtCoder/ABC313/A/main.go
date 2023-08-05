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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	ps := strings.Split(sc.Text(), " ")
	max := 0
	for i := 1; i < n; i++ {
		pi, _ := strconv.Atoi(ps[i])
		if max < pi {
			max = pi
		}
	}
	p0, _ := strconv.Atoi(ps[0])
	dif := max - p0 + 1
	if dif <= 0 {
		return "0"
	}
	return strconv.Itoa(dif)

}
