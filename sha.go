package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	var part []byte
	var err error
	var file *os.File
	var num int
	var shalist [64]string
	if file, err = os.Open("hops.log"); err != nil {
		return
	}
	reader := bufio.NewReader(file)
	num = 0
	for {
		if part, _, err = reader.ReadLine(); err != nil {
			break
		}
		if err == io.EOF {
			err = nil
		}
		hasher := sha1.New()
		hasher.Write(part)
		shalist[num] = string(hasher.Sum(nil))
		fmt.Printf("%s %x\n", part, shalist[num])
	}
}
