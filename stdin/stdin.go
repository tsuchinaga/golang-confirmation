package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 1行目にいくつのデータが渡されるかのn、2行目にスペース区切りでn個のデータ、n個のデータの合計を出すプログラムを比較

func main() {
	// FmtScan()
	ScanWords()
}

// 10.txt       0.494s
// 100.txt      0.521s
// 1000.txt     0.501s
// 10000.txt    0.512s
// 100000.txt   0.798s
// 1000000.txt  3.110s
// 2000000.txt  5.749s
// 4000000.txt 11.146s
func FmtScan() {
	var n, sum int
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		_, _ = fmt.Scan(&a)
		sum += a
	}
	fmt.Println(sum)
}

// 10.txt       0.484s
// 100.txt      0.505s
// 1000.txt     0.503s
// 10000.txt    0.498s
// 100000.txt   0.519s
// 1000000.txt  0.519s
// 2000000.txt  0.549s
// 4000000.txt  0.612s
func ScanWords() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sum := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sum += a
	}

	fmt.Println(sum)
}
