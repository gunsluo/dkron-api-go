package main

import (
	"fmt"

	"github.com/gunsluo/dkron-api-go/job_schedule"
)

func main() {

	url := "http://127.0.0.1:8080/"
	client := job_schedule.NewClient(url)

	// status
	statusResp, err := client.Status()
	fmt.Println(statusResp, err)

	//create or update
	job := job_schedule.Job{
		Name:       "lj_job",
		Schedule:   "0 30 * * * *",
		Command:    "/bin/date",
		Owner:      "jerrylou",
		OwnerEmail: "gunsluo@gmail.com",
		RunAsUser:  "jerrylou",
		Disabled:   false,
		Tags: job_schedule.JobTags{
			Role: "admin",
		},
	}
	resp, err := client.CreateOrUpdateJob(job)
	fmt.Println(resp, err)

	//show job
	resp, err = client.ShowJob("lj_job")
	fmt.Println(resp, err)

	//show all job
	resps, err := client.ShowJobs()
	fmt.Println(resps, err)

	//run job
	resp, err = client.RunJob("lj_job")
	fmt.Println(resp, err)

	//run job result
	eResps, err := client.ExecutionJobList("lj_job")
	fmt.Println(eResps, err)

	//delete job
	resp, err = client.DeleteJob("lj_job")
	fmt.Println(resp, err)

	//members
	mResps, err := client.MemberList()
	fmt.Println(mResps, err)

	//member leader
	mResp, err := client.MemberLeader()
	fmt.Println(mResp, err)
}
