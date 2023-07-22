package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type InputStrings string

func main() {
	// test := "1 2\na12\n"
	// reader := strings.NewReader(test)
	fmt.Println(Do(os.Stdin))
}

func Do(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	condition := sc.Text()
	fmt.Println(condition)
	for {
		if ok := sc.Scan(); ok == false {
			break
		}
		i := sc.Text()
	}
}
