package schedule

import (
	"log"
	"testing"
)

func TestJobsCanBeAddedToRunner(t *testing.T) {
	r := NewRunner(&TestJob{})

	if len(r.jobs) != 1 {
		t.Errorf("expected 1 job, got %d", len(r.jobs))
	}
}

func TestRunnerRunsJobs(t *testing.T) {
	j := &TestJob{}

	NewRunner(j).Run()

	if !j.hasFired {
		t.Error("TestJob did not fire")
	}
}

// This test doesn't really test anything programatically. It simply
// runs the Handler for each of the TestJobs, which, when ran several
// times, would display the concurrency ability.
func TestJobsRunConcurrently(t *testing.T) {
	j1 := &TestJob{Input: "asd"}
	j2 := &TestJob{Input: "123"}
	j3 := &TestJob{Input: "zxc"}

	NewRunner(j1, j2, j3).Run()
}

type TestJob struct {
	Input    string
	hasFired bool
}

func (t *TestJob) Handle() error {
	log.Println(t.Input)
	t.hasFired = true
	return nil
}

func (TestJob) ShouldFire() bool {
	return true
}
