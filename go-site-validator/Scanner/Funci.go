package Scanner

import (
	"net/http"
	"time"
)

type Status struct {
	URL      string
	Alive    bool
	Code     int
	Duration time.Duration
	Error    string
}

func Checksite(url string) Status {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	start := time.Now()
	resp, err := client.Head(url)
	duration := time.Since(start)
	if err != nil {
		return Status{
			URL:      url,
			Alive:    false,
			Code:     0,
			Duration: duration,
			Error:    err.Error(),
		}
	}
	defer resp.Body.Close()
	isAlive := resp.StatusCode >= 200 && resp.StatusCode < 400
	return Status{
		URL:      url,
		Alive:    isAlive,
		Code:     resp.StatusCode,
		Duration: duration,
		Error:    "",
	}
}

func worker(go_vhod <-chan string, go_vivod chan<- Status) {
	for value := range go_vhod {
		result := Checksite(value)
		go_vivod <- result
	}
}

func CheckConcurrently(urls []string, numworkers int) []Status {
	jobs := make(chan string, len(urls))
	results := make(chan Status, len(urls))
	for i := 0; i < numworkers; i++ {
		go worker(jobs, results)
	}
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)
	var allresults []Status
	for a := 0; a < len(urls); a++ {
		result := <-results
		allresults = append(allresults, result)
	}
	return allresults
}
