package schedule

// Job represents a piece of work that should be carried out when a
// certain criteria, determined by the `ShouldFire` method, is met.
type Job interface {
	// Should fire is a check that determines if a Job should run.
	ShouldFire() bool
	// Handle is called if a Job’s ShouldFire check is `true`.
	Handle() error
}

// Runner contains a number of Jobs. Running the `Run` method on
// a Runner checks if each job can be ran and runs it.
type Runner struct {
	jobs []Job
}

// Run iterates through the Jobs container in the Runner and
// determines if they should be ran. If so, it calls the Handle
// method on the Job.
func (s *Runner) Run() {
	for _, job := range s.jobs {
		if job.ShouldFire() {
			job.Handle()
		}
	}
}

// Add adds another Job to the Runner.
func (s *Runner) Add(jobs ...Job) {
	s.jobs = append(s.jobs, jobs...)
}
