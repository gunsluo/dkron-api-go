package job_schedule

type JobScheduleClient struct {
	Url string
}

func NewClient(url string) (client *JobScheduleClient) {

	client = new(JobScheduleClient)
	client.Url = url
	return client
}

func (this *JobScheduleClient) Status() (resp StatusResp, err error) {

	return Status(this.Url)
}

func (this *JobScheduleClient) CreateOrUpdateJob(job Job) (resp JobResp, err error) {

	return CreateOrUpdateJob(this.Url, job)
}

func (this *JobScheduleClient) DeleteJob(jobName string) (resp JobResp, err error) {

	return DeleteJob(this.Url, jobName)
}

func (this *JobScheduleClient) ShowJob(jobName string) (resp JobResp, err error) {

	return ShowJob(this.Url, jobName)
}

func (this *JobScheduleClient) ShowJobs() (resp []JobResp, err error) {

	return ShowJobs(this.Url)
}

func (this *JobScheduleClient) RunJob(jobName string) (resp JobResp, err error) {

	return RunJob(this.Url, jobName)
}

func (this *JobScheduleClient) ExecutionJobList(jobName string) (resp []ExecutionJobResp, err error) {

	return ExecutionJobList(this.Url, jobName)
}

func (this *JobScheduleClient) MemberList() (resp []Member, err error) {

	return MemberList(this.Url)
}

func (this *JobScheduleClient) MemberLeader() (resp Member, err error) {

	return MemberLeader(this.Url)
}
