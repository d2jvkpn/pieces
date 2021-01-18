package taskgroup

const (
	STATUS_Running    = "running"
	STATUS_Cancelled  = "cancelled"
	STATUS_Done       = "done"
	STATUS_Failed     = "failed"
	STATUS_Unexpected = "unexpected"

	STAGE_Starting  = "starting"  // created and add task
	STAGE_Running   = "running"   // waiting
	STAGE_Canceling = "canceling" // call taskgroup.Cancel
	STAGE_Exit      = "exit"      // one of tasks status is cancelled or failed
	STAGE_Done      = "done"      // all tasks are done
)
