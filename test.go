package main

import (
	"fmt"
	"math/rand"
	"runtime"

	"sort"
	"time"
)

func testData() [][]int {
	now := time.Now()
	src := rand.NewSource(now.UnixNano())
	seed := rand.New(src)
	data := make([][]int, 20000)
	for i := 0; i < len(data); i++ {
		data[i] = make([]int, 20000)
		for j := 0; j < 20000; j++ {
			data[i][j] = seed.Intn(20000)
		}
	}
	return data
}

func test() {
	data := testData()
	ch := make(chan int)
	for i := 0; i < len(data); i++ {
		go func(ch chan int, data []int) {
			sort.Ints(data[:])
			ch <- 1
		}(ch, data[i][:])
	}
	for i := 0; i < len(data); i++ {
		<-ch
	}
}

func GoMaxProcs() {
	st := time.Now()
	test()
	fmt.Println(time.Since(st))
	runtime.GOMAXPROCS(2)
	st = time.Now()
	test()
	fmt.Println(time.Since(st))
	runtime.GOMAXPROCS(3)
	st = time.Now()
	test()
	fmt.Println(time.Since(st))
	runtime.GOMAXPROCS(4)
	st = time.Now()
	test()
	fmt.Println(time.Since(st))

	fmt.Println("Hello World!")
}
