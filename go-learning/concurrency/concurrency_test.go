package concurrency

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

// captureOutput captures stdout during function execution
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestDemonstrateGoroutines(t *testing.T) {
	done := make(chan bool)

	go func() {
		output := captureOutput(demonstrateGoroutines)

		// Verify output contains expected content
		if !strings.Contains(output, "Goroutines:") {
			t.Error("Expected output to contain 'Goroutines:'")
		}
		if !strings.Contains(output, "Sequential:") {
			t.Error("Expected output to contain 'Sequential:'")
		}
		if !strings.Contains(output, "Concurrent:") {
			t.Error("Expected output to contain 'Concurrent:'")
		}
		if !strings.Contains(output, "Hello") {
			t.Error("Expected output to contain 'Hello'")
		}
		if !strings.Contains(output, "World") {
			t.Error("Expected output to contain 'World'")
		}

		done <- true
	}()

	select {
	case <-done:
		// Test completed successfully
	case <-time.After(2 * time.Second):
		t.Fatal("Test timed out - demonstrateGoroutines took too long")
	}
}

func TestDemonstrateChannels(t *testing.T) {
	done := make(chan bool)

	go func() {
		output := captureOutput(demonstrateChannels)

		if !strings.Contains(output, "Channels:") {
			t.Error("Expected output to contain 'Channels:'")
		}
		if !strings.Contains(output, "Hello from channel!") {
			t.Error("Expected output to contain 'Hello from channel!'")
		}

		done <- true
	}()

	select {
	case <-done:
		// Test completed successfully
	case <-time.After(2 * time.Second):
		t.Fatal("Test timed out - demonstrateChannels took too long")
	}
}

func TestDemonstrateBufferedChannels(t *testing.T) {
	done := make(chan bool)

	go func() {
		output := captureOutput(demonstrateBufferedChannels)

		if !strings.Contains(output, "Buffered Channels:") {
			t.Error("Expected output to contain 'Buffered Channels:'")
		}
		// Should receive both values
		if strings.Count(output, "Received:") != 2 {
			t.Error("Expected output to contain two 'Received:' messages")
		}

		done <- true
	}()

	select {
	case <-done:
		// Test completed successfully
	case <-time.After(2 * time.Second):
		t.Fatal("Test timed out - demonstrateBufferedChannels took too long")
	}
}

func TestDemonstrateWaitGroup(t *testing.T) {
	done := make(chan bool)

	go func() {
		output := captureOutput(demonstrateWaitGroup)

		if !strings.Contains(output, "WaitGroup") {
			t.Error("Expected output to contain 'WaitGroup'")
		}
		if !strings.Contains(output, "All workers done!") {
			t.Error("Expected output to contain 'All workers done!'")
		}
		// Check that all 3 workers started and finished
		for i := 1; i <= 3; i++ {
			if !strings.Contains(output, "Worker "+string(rune('0'+i))+" starting") {
				t.Errorf("Expected output to contain 'Worker %d starting'", i)
			}
			if !strings.Contains(output, "Worker "+string(rune('0'+i))+" done") {
				t.Errorf("Expected output to contain 'Worker %d done'", i)
			}
		}

		done <- true
	}()

	select {
	case <-done:
		// Test completed successfully
	case <-time.After(2 * time.Second):
		t.Fatal("Test timed out - demonstrateWaitGroup took too long")
	}
}

func TestDemonstrateSelect(t *testing.T) {
	done := make(chan bool)

	go func() {
		output := captureOutput(demonstrateSelect)

		if !strings.Contains(output, "Select") {
			t.Error("Expected output to contain 'Select'")
		}
		if !strings.Contains(output, "Message from channel 1") {
			t.Error("Expected output to contain 'Message from channel 1'")
		}
		if !strings.Contains(output, "Message from channel 2") {
			t.Error("Expected output to contain 'Message from channel 2'")
		}

		done <- true
	}()

	select {
	case <-done:
		// Test completed successfully
	case <-time.After(2 * time.Second):
		t.Fatal("Test timed out - demonstrateSelect took too long")
	}
}

func TestRunConcurrency(t *testing.T) {
	done := make(chan bool)

	go func() {
		output := captureOutput(RunConcurrency)

		// Verify all sections are present
		sections := []string{
			"Goroutines:",
			"Channels:",
			"Buffered Channels:",
			"WaitGroup",
			"Select",
		}

		for _, section := range sections {
			if !strings.Contains(output, section) {
				t.Errorf("Expected output to contain '%s'", section)
			}
		}

		done <- true
	}()

	select {
	case <-done:
		// Test completed successfully
	case <-time.After(5 * time.Second):
		t.Fatal("Test timed out - RunConcurrency took too long")
	}
}

func TestPrintMessage(t *testing.T) {
	output := captureOutput(func() {
		printMessage("Test", 3)
	})

	// Should contain 3 iterations
	if strings.Count(output, "Test") != 3 {
		t.Errorf("Expected 3 occurrences of 'Test', got %d", strings.Count(output, "Test"))
	}
	if !strings.Contains(output, "(1)") {
		t.Error("Expected output to contain '(1)'")
	}
	if !strings.Contains(output, "(2)") {
		t.Error("Expected output to contain '(2)'")
	}
	if !strings.Contains(output, "(3)") {
		t.Error("Expected output to contain '(3)'")
	}
}

func TestWorker(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	output := captureOutput(func() {
		worker(1, &wg)
	})

	if !strings.Contains(output, "Worker 1 starting") {
		t.Error("Expected output to contain 'Worker 1 starting'")
	}
	if !strings.Contains(output, "Worker 1 done") {
		t.Error("Expected output to contain 'Worker 1 done'")
	}
}
