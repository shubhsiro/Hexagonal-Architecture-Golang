// consensus.go
package consensus

import (
	"strings"
	"sync"
	"time"

	"Sarva_Assignment/internal/core/logger"
)

// Consensus represents the core logic for consensus interactions.
type Consensus struct {
	raftMessages []string
	mu           sync.Mutex
	logger       logger.Logger // Use the logger interface here
}

// NewConsensus creates a new instance of Consensus.
func NewConsensus(logger logger.Logger) *Consensus {
	return &Consensus{
		raftMessages: make([]string, 0),
		logger:       logger,
	}
}

// SendMessage sends a message for consensus.
func (c *Consensus) SendMessage(message string) error {
	c.logger.Log("Sending message for consensus: " + message)

	c.mu.Lock()
	c.raftMessages = append(c.raftMessages, message)
	c.mu.Unlock()

	c.logger.Log("Message sent successfully.")

	return nil
}

// RetrieveMessages retrieves consensus messages.
func (c *Consensus) RetrieveMessages() ([]string, error) {
	c.logger.Log("Retrieving consensus messages")
	c.mu.Lock()

	defer c.mu.Unlock()

	messages := make([]string, len(c.raftMessages))
	copy(messages, c.raftMessages)

	c.logger.Log("Retrieved messages: " + strings.Join(messages, ", "))

	return messages, nil
}

// StartCluster starts the consensus cluster.
func (c *Consensus) StartCluster() {
	c.logger.Log("Starting the consensus cluster")

	time.Sleep(1 * time.Second)

	c.logger.Log("Consensus cluster started successfully.")
}

// Shutdown shuts down the consensus cluster.
func (c *Consensus) Shutdown() {
	c.logger.Log("Shutting down the consensus cluster")

	time.Sleep(1 * time.Second)

	c.logger.Log("Consensus cluster shut down successfully.")
}
