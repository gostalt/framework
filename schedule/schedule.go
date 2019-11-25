package schedule

import "sync"

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

// NewRunner creates a new runner from a list of Jobs.
func NewRunner(jobs ...Job) *Runner {
	return &Runner{
		jobs: jobs,
	}
}

// Run iterates through the Jobs container in the Runner and
// determines if they should be ran. If so, it calls the Handle
// method on the Job.
func (r *Runner) Run() {
	wg := sync.WaitGroup{}
	for _, job := range r.jobs {
		if job.ShouldFire() {
			wg.Add(1)

			// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
			go func(job Job) {
				job.Handle()
				wg.Done()
			}(job)
		}
	}

	wg.Wait()
}

// Add adds another Job to the Runner.
func (r *Runner) Add(jobs ...Job) {
	r.jobs = append(r.jobs, jobs...)
}
