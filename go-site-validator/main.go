package main

import (
	"awesomeProject/Scanner"
	"fmt"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://httpbin.org/status/404",
		"https://invalid.website.that.does.not.exist",
	}
	fmt.Println("=== Последовательная проверка ===")
	nowt := time.Now()
	for _, url := range urls {
		result := Scanner.Checksite(url)
		fmt.Printf("%s: %v (%dms)\n", url, result.Alive, result.Duration.Milliseconds())
	}
	finish := time.Since(nowt)
	fmt.Println("\n=== Параллельная проверка (3 workers) ===")
	start := time.Now()
	results := Scanner.CheckConcurrently(urls, 3)
	concTime := time.Since(start)
	for _, result := range results {
		fmt.Printf("%s: %v (%dms)\n", result.URL, result.Alive, result.Duration.Milliseconds())
	}
	fmt.Printf("\n=== ИТОГИ ===\n")
	fmt.Printf("Последовательно: %v\n", finish)
	fmt.Printf("Параллельно (3 workers): %v\n", concTime)
	fmt.Printf("Ускорение: %.1fx\n", float64(finish)/float64(concTime))
}
