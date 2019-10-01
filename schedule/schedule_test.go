package schedule

import "testing"

func TestJobsCanBeAddedToRunner(t *testing.T) {
	r := Runner{}

	r.Add(&TestJob{})

	if len(r.jobs) != 1 {
		t.Errorf("expected 1 job, got %d", len(r.jobs))
	}
}

func TestRunnerRunsJobs(t *testing.T) {
	r := Runner{}

	j := &TestJob{}

	r.Add(j)

	r.Run()

	if !j.hasFired {
		t.Error("TestJob did not fire")
	}
}

type TestJob struct {
	hasFired bool
}

func (t *TestJob) Handle() error {
	t.hasFired = true
	return nil
}

func (TestJob) ShouldFire() bool {
	return true
}
