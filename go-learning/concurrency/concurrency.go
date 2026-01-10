package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// RunConcurrency demonstrates Go's concurrency features
func RunConcurrency() {
	demonstrateGoroutines()
	demonstrateChannels()
	demonstrateBufferedChannels()
	demonstrateWaitGroup()
	demonstrateSelect()
}

func demonstrateGoroutines() {
	fmt.Println("Goroutines:")

	// Sequential execution
	fmt.Println("  Sequential:")
	printMessage("Hello", 2)
	printMessage("World", 2)

	// Concurrent execution with goroutines
	fmt.Println("  Concurrent:")
	go printMessage("Goroutine 1", 2)
	go printMessage("Goroutine 2", 2)

	// Wait for goroutines to finish
	time.Sleep(time.Millisecond * 500)
	fmt.Println()
}

func printMessage(msg string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("  %s (%d)\n", msg, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func demonstrateChannels() {
	fmt.Println("Channels:")

	// Create a channel
	messages := make(chan string)

	// Send value in a goroutine
	go func() {
		messages <- "Hello from channel!"
	}()

	// Receive value (blocking)
	msg := <-messages
	fmt.Printf("  Received: %s\n", msg)
}

func demonstrateBufferedChannels() {
	fmt.Println("Buffered Channels:")

	// Buffered channel (can hold 2 values)
	numbers := make(chan int, 2)

	// Send without blocking
	numbers <- 1
	numbers <- 2

	// Receive values
	fmt.Printf("  Received: %d\n", <-numbers)
	fmt.Printf("  Received: %d\n", <-numbers)
}

func demonstrateWaitGroup() {
	fmt.Println("WaitGroup (sync multiple goroutines):")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("  All workers done!")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when done

	fmt.Printf("  Worker %d starting\n", id)
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("  Worker %d done\n", id)
}

func demonstrateSelect() {
	fmt.Println("Select (multiplexing channels):")

	c1 := make(chan string)
	c2 := make(chan string)

	// Two goroutines sending to different channels
	go func() {
		time.Sleep(time.Millisecond * 100)
		c1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(time.Millisecond * 200)
		c2 <- "Message from channel 2"
	}()

	// Select receives from whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("  %s\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("  %s\n", msg2)
		}
	}
}
