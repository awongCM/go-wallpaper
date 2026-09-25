package rotate

import (
	"sync"
	"testing"
	"time"
)

type stubBackend struct {
	mu      sync.Mutex
	setCalls []string
}

func (s *stubBackend) Name() string { return "stub" }

func (s *stubBackend) Get() (string, error) { return "", nil }

func (s *stubBackend) DefaultWallpaperDir() string { return "" }

func (s *stubBackend) Set(imagePath string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.setCalls = append(s.setCalls, imagePath)
	return nil
}

func (s *stubBackend) sets() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]string, len(s.setCalls))
	copy(out, s.setCalls)
	return out
}

func TestRunUntilStop_restoresOriginalOnInterrupt(t *testing.T) {
	backend := &stubBackend{}
	stop := make(chan struct{})
	done := make(chan error, 1)

	go func() {
		done <- runUntilStop(backend, []string{"/tmp/a.jpg", "/tmp/b.jpg"}, time.Hour, "/original.jpg", stop)
	}()

	time.Sleep(20 * time.Millisecond)
	close(stop)

	if err := <-done; err != nil {
		t.Fatalf("runUntilStop: %v", err)
	}

	sets := backend.sets()
	if len(sets) < 1 {
		t.Fatalf("expected at least one set call, got %v", sets)
	}
	last := sets[len(sets)-1]
	if last != "/original.jpg" {
		t.Fatalf("expected restore to %q, last set was %q (all: %v)", "/original.jpg", last, sets)
	}
}

func TestRunUntilStop_skipsRestoreWhenOriginalEmpty(t *testing.T) {
	backend := &stubBackend{}
	stop := make(chan struct{})

	go func() {
		_ = runUntilStop(backend, []string{"/tmp/a.jpg"}, time.Hour, "", stop)
	}()

	time.Sleep(20 * time.Millisecond)
	close(stop)
	time.Sleep(20 * time.Millisecond)

	sets := backend.sets()
	for _, path := range sets {
		if path == "" {
			t.Fatal("unexpected empty restore set")
		}
	}
	if len(sets) != 1 || sets[0] != "/tmp/a.jpg" {
		t.Fatalf("expected single rotation set, got %v", sets)
	}
}

func TestRun_rejectsEmptyPaths(t *testing.T) {
	err := Run(&stubBackend{}, nil, time.Second, "/x")
	if err == nil {
		t.Fatal("expected error for empty paths")
	}
}
