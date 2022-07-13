package workerpool

// Job each job instance must implement Job interface
type Job interface {
	// ExecJob method run job instance
	ExecJob() error
	// JobID method return identity for job (for storing leak job in file)
	JobID() string
}
