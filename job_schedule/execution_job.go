package job_schedule

import "time"

type ExecutionJobResp struct {
	JobName    string    `josn:"job_name"`
	StartedAt  time.Time `josn:"started_at"`
	FinishedAt time.Time `josn:"finished_at"`
	Success    bool      `josn:"success"`
	Output     string    `josn:"output"`
	NodeName   string    `josn:"node_name"`
}
