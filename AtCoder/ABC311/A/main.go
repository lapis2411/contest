package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type InputStrings string

func main() {
	fmt.Println(Do(os.Stdin))
}

func Do(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	sc.Scan()
	str := sc.Text()
	eA := false
	eB := false
	eC := false
	for i, v := range str {
		if v == 'A' {
			eA = true
		} else if v == 'B' {
			eB = true
		} else if v == 'C' {
			eC = true
		}
		if eA && eB && eC {
			return strconv.Itoa(i + 1)
		}
	}
	return ""
}
