package consensus

import (
	"log"
	"os"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

// RaftAdapter represents the adapter for RAFT consensus.
type RaftAdapter struct {
	raft   *raft.Raft
	logger hclog.Logger
}

// NewRaftAdapter creates a new instance of RaftAdapter.
func NewRaftAdapter(nodeID string, logs hclog.Logger) (*RaftAdapter, error) {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(nodeID)
	config.SnapshotInterval = 1000

	// Create a Raft backend using a BoltDB store
	boltDB, err := raftboltdb.NewBoltStore("raft.db")
	if err != nil {
		log.Printf("Error creating BoltDB store: %v", err)

		return nil, err
	}

	// Create a snapshot store
	snapshotStore, err := raft.NewFileSnapshotStore(".", 1, os.Stdout)
	if err != nil {
		log.Printf("Error creating snapshot store: %v", err)
		return nil, err
	}

	// Create a transport layer
	_, transport := raft.NewInmemTransport(raft.ServerAddress(config.LocalID))

	// Create a Raft configuration
	raftConfig := &raft.Config{
		ProtocolVersion:  config.ProtocolVersion,
		LocalID:          config.LocalID,
		HeartbeatTimeout: config.HeartbeatTimeout,
		ElectionTimeout:  config.ElectionTimeout,
		CommitTimeout:    config.CommitTimeout,
		MaxAppendEntries: config.MaxAppendEntries,
		ShutdownOnRemove: config.ShutdownOnRemove,
		Logger:           config.Logger,
	}

	var (
		fsm    raft.FSM
		stable raft.StableStore
	)

	// Create a Raft instance
	raftInstance, err := raft.NewRaft(raftConfig, fsm, boltDB, stable, snapshotStore, transport)
	if err != nil {
		log.Printf("Error creating Raft: %v", err)

		return nil, err
	}

	return &RaftAdapter{raft: raftInstance, logger: logs}, nil
}

// SendMessage sends a message to the RAFT cluster.
func (r *RaftAdapter) SendMessage(message string) error {
	future := r.raft.Apply([]byte(message), 10*time.Second)
	if err := future.Error(); err != nil {
		log.Printf("Error applying Raft log: %v", err)

		return err
	}

	return nil
}
