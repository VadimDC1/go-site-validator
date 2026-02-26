package main

import (
	"awesomeProject/Scanner"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Проверка доступности сайтов")
	fmt.Print("Введите URL (можно несколько через пробел): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	urls := strings.Fields(input)
	for i, url := range urls {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			urls[i] = "https://" + url
		}
	}
	fmt.Println("\n=== Последовательная проверка ===")
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
