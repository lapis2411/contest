package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type InputStrings string

func main() {
	fmt.Println(Do(os.Stdin))
}

func Do(reader io.Reader) string {
	in := bufio.NewReader(reader)
	var n, A int
	fmt.Fscan(in, &n)

	// sum := 0
	favg := 0.0
	ns := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A)
		ns[i] = A
		favg += float64(ns[i]) / float64(n)
	}
	// sc := bufio.NewScanner(reader)
	// sc.Scan()
	// n, _ := strconv.Atoi(sc.Text())
	// sc.Scan()
	// as := strings.Split(sc.Text(), " ")
	// ns := make([]int, n)
	// sum := 0
	// favg := 0.0
	// for i, v := range as {
	// 	ns[i], _ = strconv.Atoi(v)
	// 	sum += ns[i]
	// 	favg += float64(ns[i]) / float64(n)
	// }
	// avg := int(math.Round(float64(sum) / float64(n)))
	// avg := sum / n
	avg := int(favg)
	diff_s := 0
	diff := 0
	for _, v := range ns {
		d := 0
		if v > avg {
			d = v - (avg + 1)
			diff += d
		} else {
			d = avg - v
			diff -= d
		}
		diff_s += d
	}

	ret := diff_s / 2
	if diff != 0 {
		ad := math.Ceil(math.Abs(float64(diff)) / 2)
		ret += int(ad)
	}
	return strconv.Itoa(ret)

}
