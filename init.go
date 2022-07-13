package workerpool

// Config of worker pool
//
// max number of job can solve without blocking: config.NumJobQueue + config.NumWorkers
type Config struct {
	// maximum number of workers can run job
	NumWorkers int32
	// maximum number of jobs can send to queue
	NumJobQueue int32
}

var workerPool chan chan Job
var jobQueue chan Job

// SendJob send job to job queue and then deliver to worker pool
func SendJob(job Job) {
	// reach limit
	// store job file
	if len(jobQueue) == cap(jobQueue) {

		return
	}

	jobQueue <- job
}

// Init worker pool
func Init(config Config) {
	// max number of job can solve without blocking: config.NumJobQueue + config.NumWorkers
	workerPool = make(chan chan Job, config.NumWorkers)
	jobQueue = make(chan Job, config.NumJobQueue)

	for i := 0; i < int(config.NumWorkers); i++ {
		newWorker(workerPool).start()
	}

	jobAssign := func() {
		for {
			select {
			// receive job from job queue
			case job := <-jobQueue:
				// get worker
				worker := <-workerPool
				// assign job for worker
				worker <- job
			}
		}
	}

	// run job assign by fork new goroutine
	go jobAssign()
}
