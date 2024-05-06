package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

		// Change to your Api
	apiEnd := "http://localhost:8090/view/TestPage"
	totalRequest := 100

	var wg sync.WaitGroup
	// A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

	wg.Add(totalRequest)
	done := make(chan struct{})

	start := time.Now()

	for i := 0; i < totalRequest; i++ {
		go func() {
			defer wg.Done()

			// Sending Requests
			resp, err := http.Get(apiEnd)
			if err != nil {
				fmt.Println("Error :", err)
				return
			}
			defer resp.Body.Close()

			fmt.Println("Response code : ", resp.StatusCode)

			done <- struct{}{}

		}()
	}
	// Wait for all requests to complete
	go func() {
		wg.Wait()
		close(done)
	}()

	timeout := time.Second
	select {
	case <-done:
		fmt.Println("All request completed")
	case <-time.After(timeout):
		fmt.Println("Timeout: Some requests might still Pending")
	}

	elapsed := time.Since(start)
	fmt.Printf("Total time Taken %s ", elapsed)
}
