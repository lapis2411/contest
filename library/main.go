package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type InputStrings string

func main() {
	test := "1 2\na12\n"
	reader := strings.NewReader(test)
	Do(reader)
}

func Do(reader io.Reader) {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	condition := sc.Text()
	fmt.Println(condition)
	var bufr []string
	for {
		if ok := sc.Scan(); ok == false {
			break
		}
		bufr = append(bufr, sc.Text())
	}
	for _, v := range bufr {
		fmt.Println(v)
	}
}
