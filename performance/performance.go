package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))             //print memory usage in MiB
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc)) //print total memory allocated in MiB
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))               //print total system memory in MiB
	fmt.Printf("\tNumGC = %v\n", m.NumGC)                    //print number of garbage collections
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func hundredMillionMap() {
	printMemUsage()
	m := make(map[int]int, 100000000)
	start := time.Now()

	for i := 0; i < 100000000; i++ {
		m[i] = i
	}

	elapsed := time.Since(start)

	fmt.Println("Map initialized with one hundred million entries", elapsed)
	printMemUsage()
}

// same function as testloop but using go routines for efficiency
func hundredMillionSlice() {
	printMemUsage()
	start := time.Now()
	s := make([]int, 100000000)

	for i := range s {
		s[i] = i
	}

	elapsed := time.Since(start)

	fmt.Println("Slice initialized with one hundred million entries", elapsed)
	printMemUsage()
}

func main() {
	fmt.Println("Hello, World!")

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		hundredMillionMap()
		wg.Done()
	}()

	go func() {
		hundredMillionSlice()
		wg.Done()
	}()

	wg.Wait()
}
