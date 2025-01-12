package constants

const (
	// TempTaskOutputPrefix is prefix for files to hold temporary output of task.
	// Temporariy files used to reduce load on DB in case of a lot of parts in task's output
	TempTaskOutputPrefix = "temp_task_output_"
)
