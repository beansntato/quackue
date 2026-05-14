package quackue_test

import (
	"fmt"
	"testing"
)

type CrashPoint int

const (
	AfterWALBeforeEnqueue CrashPoint = iota
	AfterEnqueueBeforeDequeue
	AfterDequeueBeforeExec
	AfterExecBeforeWalAck
	AfterWalBeforeDelete
)

func runJobWithCrashAt(point CrashPoint) error {
	// start wal, producer, and worker
	// insert a single job
	// inject crash at a point of failure
	// test recovery: replay/dlq
	// test idempotency, the single job should only run successfully 1 time
	// return result
	return nil
}

func TestExactlyOnce(t *testing.T) {
	points := []CrashPoint{
		AfterWALBeforeEnqueue,
		AfterEnqueueBeforeDequeue,
		AfterDequeueBeforeExec,
		AfterExecBeforeWalAck,
		AfterWalBeforeDelete,
	}

	for _, point := range points {
		for i := 0; i < 100; i++ {
			err := runJobWithCrashAt(point)
			fmt.Println("error", err)
		}
	}
}
