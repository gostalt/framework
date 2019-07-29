package schedule

// Job represents a piece of work that should be carried out when a
// certain criteria, determined by the `ShouldFire` method, is met.
type Job interface {
	Handle() error
	ShouldFire() bool
}

// Runner runs the Jobs contained in it. This can be triggered
// manually, but in production you should set a cron job up to
// call this functionality automatically.
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
