package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	file1, err := os.OpenFile("./file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}

	file2, err := os.OpenFile("./file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}

	_, err = file1.WriteString("hello, file1\n")
	if err != nil {
		log.Println(err)
	}

	_, err = file2.WriteString("hello, file2\n")
	if err != nil {
		log.Println(err)
	}

	// この時点ではファイルを閉じていないので削除できない
	// err = os.Remove("./file")
	// if err != nil {
	// 	log.Println(err)
	// }

	sc1 := bufio.NewScanner(file1)
	sc1.Split(bufio.ScanLines)
	for sc1.Scan() {
		log.Println(sc1.Text())
	}

	sc2 := bufio.NewScanner(file2)
	for sc2.Scan() {
		log.Println(sc2.Text())
	}

	_ = file1.Close()
	_ = file2.Close()

	err = os.Remove("./file")
	if err != nil {
		log.Println(err)
	}
}
