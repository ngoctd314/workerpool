package workerpool

// Job each job instance must implement Job interface
type Job interface {
	// Exec method run job instance
	Exec() error
	// ID method return identity for job (for storing leak job in file)
	ID() string
}
