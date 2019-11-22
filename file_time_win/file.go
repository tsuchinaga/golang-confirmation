package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	fmt.Println("こんにちわーるど")

	fi, err := os.Stat("target_file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	if sys, ok := fi.Sys().(*syscall.Win32FileAttributeData); ok {
		fmt.Println("sys.CreationTime", sys.CreationTime, time.Unix(0, sys.CreationTime.Nanoseconds()))
		fmt.Println("sys.FileAttributes", sys.FileAttributes)
		fmt.Println("sys.FileSizeHigh", sys.FileSizeHigh)
		fmt.Println("sys.LastAccessTime", sys.LastAccessTime, time.Unix(0, sys.LastAccessTime.Nanoseconds()))
		fmt.Println("sys.LastWriteTime", sys.LastWriteTime, time.Unix(0, sys.LastWriteTime.Nanoseconds()))
	}
}
