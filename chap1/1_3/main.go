package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	args1()
	args2()
	args3()
}

func args1() {
	tt1 := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	tt2 := time.Since(tt1)
	fmt.Println(tt2)
}

func args2() {
	tt1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	tt2 := time.Since(tt1)
	fmt.Println(tt2)
}

func args3() {
	tt1 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	tt2 := time.Since(tt1)
	fmt.Println(tt2)
}
