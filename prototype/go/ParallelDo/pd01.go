package main

// https://talks.golang.org/2013/distsys.slide#36
import (
	// "fmt"
)

type ParallelDo struct {
	n    int
	max  int
	done chan bool
}

func NewParallelDo(max int) ParallelDo {
	return ParallelDo{max: max, done: make(chan bool, max)}
}

func (pd *ParallelDo) Run(job func()) {
	pd.n++
	if pd.n > pd.max {
		<-pd.done
		pd.n--
	}

	job()
	pd.done <- true
}

func (pd *ParallelDo) Wait() {
	for ; pd.n > 0; pd.n-- {
		<-pd.done
	}
}
