package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Message represents the data structure
type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

// Create a buffered channel to act as the message queue
// It holds up to 100 messages before blocking the sender
var messageQueue = make(chan Message, 100)

// worker processes messages from the queue asynchronously
func worker(id int) {
	for msg := range messageQueue {
		// Simulate network latency or processing time
		time.Sleep(1 * time.Second)
		fmt.Printf("[Worker %d] Delivered: [%s] -> [%s]: %s\n", id, msg.Sender, msg.Recipient, msg.Content)
	}
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Push the message into the queue (non-blocking if channel is not full)
	select {
	case messageQueue <- msg:
		// Message accepted into queue successfully
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted) // 202 Accepted
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "queued",
			"message": "Message received and queued for " + msg.Recipient,
		})
	default:
		// Queue is full
		http.Error(w, "Server busy, queue full", http.StatusServiceUnavailable)
	}
}

func main() {
	// Start 3 background worker goroutines to process the queue
	for i := 1; i <= 3; i++ {
		go worker(i)
	}

	http.HandleFunc("/send", sendMessageHandler)

	fmt.Println("Server starting with Message Queue on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
