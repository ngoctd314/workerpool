package workerpool

type worker struct {
	pool chan chan Job
	// job run on this worker
	jobChannel chan Job
}

func newWorker(pool chan chan Job) worker {
	return worker{
		pool:       pool,
		jobChannel: make(chan Job),
	}
}

// start worker
func (w worker) start() {
	handle := func() {
		for {
			// send job address to pool
			// if pool buffer is full => block
			w.pool <- w.jobChannel

			select {
			case job := <-w.jobChannel:
				// worker receive new job, run it
				err := job.Exec()
				// if error happen, store job to file
				if err != nil {
					// store to file
				}
			}
		}
	}

	// run worker by fork new goroutine
	go handle()
}
