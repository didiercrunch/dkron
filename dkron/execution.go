package dkron

import (
	"fmt"
	"time"
)

type Execution struct {
	// Name of the job this executions refers to.
	JobName string `json:"job_name,omitempty"`

	// Start time of the execution.
	StartedAt time.Time `json:"started_at,omitempty"`

	// When the execution finished running.
	FinishedAt time.Time `json:"finished_at,omitempty"`

	// If this execution executed succesfully.
	Success bool `json:"success,omitempty"`

	// Partial output of the execution.
	Output []byte `json:"output,omitempty"`

	// Node name of the node that run this execution.
	NodeName string `json:"node_name,omitempty"`

	// Execution group to what this execution belongs to.
	Group int64 `json:"group,omitempty"`

	// Retry attempt of this execution.
	Attempt uint `json:"attempt,omitempty"`
}


type ExecutionsByStartedAt []*Execution


func (a ExecutionsByStartedAt) Len() int{
	return len(a)
}


func (a ExecutionsByStartedAt) Swap(i, j int)  {
	a[i], a[j] = a[j], a[i]
}


func (a ExecutionsByStartedAt) Less(i, j int) bool {
	return a[i].StartedAt.UnixNano() < a[j].StartedAt.UnixNano()
}


// Init a new execution
func NewExecution(jobName string) *Execution {
	return &Execution{
		JobName: jobName,
		Group:   time.Now().UnixNano(),
		Attempt: 1,
	}
}

// Used to enerate the execution Id
func (e *Execution) Key() string {
	return fmt.Sprintf("%d-%s", e.StartedAt.UnixNano(), e.NodeName)
}
