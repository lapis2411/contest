package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const PI = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"

func main() {
	fmt.Println(Do(os.Stdin))
}

func Do(reader io.Reader) string {
	in := bufio.NewReader(reader)
	var n int
	fmt.Fscan(in, &n)
	rPI := []rune(PI)
	ret := rPI[0 : n+2]
	return string(ret)
}
