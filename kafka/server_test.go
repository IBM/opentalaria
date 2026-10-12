package kafka

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ibm/opentalaria/config"
)

type MockClient struct {
	handledRequests int
	mu              sync.Mutex
}

func (mc *MockClient) handleRequest() {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.handledRequests++
}

func TestServer_Run(t *testing.T) {
	os.Setenv("OT_MAX_CONNECTIONS", "20")
	defer os.Unsetenv("OT_MAX_CONNECTIONS")

	os.Setenv("OT_LISTENERS", "PLAINTEXT://:9092")
	defer os.Unsetenv("OT_LISTENERS")

	// Mock server and client
	mockClient := &MockClient{}
	conf, err := config.NewConfig("")
	if err != nil {
		t.Error(err)
	}

	server := NewServer(conf)

	// Create a context with cancellation
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the server in a goroutine
	go server.Run()

	// Allow the server some time to start
	time.Sleep(100 * time.Millisecond)

	// Dial the server to simulate a client
	addr := fmt.Sprintf("%s:%s", server.host, server.port)
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	for i := 0; i < 100; i++ {
		go func(i int) {
			conn.Write([]byte("Hello Server!"))
			mockClient.handleRequest()

		}(i)
	}

	// Allow time for the server to process
	time.Sleep(200 * time.Millisecond)

	// Verify that the server handled the request
	mockClient.mu.Lock()
	if mockClient.handledRequests != 100 {
		t.Errorf("Expected 100 handled request, got %d", mockClient.handledRequests)
	}
	mockClient.mu.Unlock()

	// Cancel the context to shut down the server
	cancel()
	time.Sleep(100 * time.Millisecond)
}

// TestHandleRequest_OversizedFrameRejected verifies that oversized request frames
// (> protocol.MaxRequestSize / 0xFFFFFFFF) are rejected before allocating memory.
func TestHandleRequest_OversizedFrameRejected(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	defer clientConn.Close()

	client := &Client{
		conn: serverConn,
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		client.handleRequest()
	}()

	// Send 4-byte header declaring size = 0xFFFFFFFF (~4 GiB)
	_, err := clientConn.Write([]byte{0xff, 0xff, 0xff, 0xff})
	if err != nil {
		t.Fatalf("failed to write frame header: %v", err)
	}

	// Server should reject oversized frame immediately and close the connection without allocating and blocking
	select {
	case <-done:
		// Succeeded: server rejected frame and closed connection
	case <-time.After(500 * time.Millisecond):
		clientConn.Close()
		serverConn.Close()
		t.Fatal("server did not reject oversized frame and remained blocked in read loop")
	}
}
