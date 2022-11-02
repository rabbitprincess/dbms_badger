package util

import (
	"sync"
	"time"
)

type SnowflakeID struct {
	mu            sync.Mutex
	lastTimestamp int64
	sequence      uint32
	NodeID        uint32
}

func NewSnowflake(NodeID uint32) *SnowflakeID {
	return &SnowflakeID{
		lastTimestamp: 0,
		sequence:      0,
		NodeID:        NodeID,
	}
}

func (t *SnowflakeID) Next() uint64 {
	t.mu.Lock()
	timestamp := time.Now().UnixMilli()
	if timestamp > t.lastTimestamp {
		t.lastTimestamp = timestamp
		t.sequence = 0
	} else {
		t.sequence++
	}
	s := uint64(timestamp)<<22 | uint64(t.NodeID&(1<<10-1))<<12 | uint64(t.sequence&(1<<12-1))
	t.mu.Unlock()
	return s
}
