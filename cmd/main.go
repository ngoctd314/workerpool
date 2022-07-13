package main

import (
	"fmt"
	"time"

	"github.com/ngoctd314/workerpool"
)

type person struct {
	name string
}

func (p person) ExecJob() error {
	fmt.Println(p)
	return nil
}

func (p person) JobID() string {
	return "person"
}

func main() {
	workerpool.Init(workerpool.Config{
		NumWorkers:  5,
		NumJobQueue: 10,
	})

	workerpool.SendJob(person{"tdn"})
	time.Sleep(time.Second * 5)
}
