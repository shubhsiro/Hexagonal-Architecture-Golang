package consensus

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Trace(message string) {
	//TODO implement me
	panic("implement me")
}

func (m *MockLogger) Debug(message string) {
	//TODO implement me
	panic("implement me")
}

func (m *MockLogger) Metrics(metrics map[string]interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

func TestSendMessage(t *testing.T) {
	mockLogger := &MockLogger{}
	consensus := NewConsensus(mockLogger)

	err := consensus.SendMessage("Test Message")
	assert.NoError(t, err)
	assert.Contains(t, mockLogger.Logs, "Sending message for consensus: Test Message")
	assert.Contains(t, mockLogger.Logs, "Message sent successfully.")
}

func TestStartCluster(t *testing.T) {
	mockLogger := &MockLogger{}
	consensus := NewConsensus(mockLogger)

	consensus.StartCluster()
	assert.Contains(t, mockLogger.Logs, "Starting the consensus cluster")
	time.Sleep(2 * time.Second) // Allow some time for StartCluster to complete
	assert.Contains(t, mockLogger.Logs, "Consensus cluster started successfully.")
}

func TestShutdown(t *testing.T) {
	mockLogger := &MockLogger{}
	consensus := NewConsensus(mockLogger)

	consensus.Shutdown()
	assert.Contains(t, mockLogger.Logs, "Shutting down the consensus cluster")
	time.Sleep(2 * time.Second) // Allow some time for Shutdown to complete
	assert.Contains(t, mockLogger.Logs, "Consensus cluster shut down successfully.")
}
