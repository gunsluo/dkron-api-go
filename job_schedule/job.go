package job_schedule

import "time"

type JobTags struct {
	Role   string `json:"role"`
	Server string `json:"server"`
}

type Job struct {
	Name       string  `json:"name"`
	Schedule   string  `json:"schedule"`
	Command    string  `json:"command"`
	Owner      string  `json:"owner"`
	OwnerEmail string  `json:"owner_email"`
	RunAsUser  string  `json:"run_as_user"`
	Disabled   bool    `json:"disabled"`
	Tags       JobTags `json:"tags"`
}

type JobResp struct {
	SuccessCount int       `json:"success_count"`
	ErrorCount   int       `json:"error_count"`
	LastSuccess  time.Time `json:"last_success"`
	LastError    time.Time `json:"last_error"`
	Job
}
