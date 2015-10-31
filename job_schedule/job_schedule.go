package job_schedule

import "github.com/pquerna/ffjson/ffjson"

func Status(url string) (resp StatusResp, err error) {

	body, err := send(url+"v1/", "")
	if err != nil {
		return
	}
	buf := []byte(body)
	err = ffjson.Unmarshal(buf, &resp)

	return
}

func CreateOrUpdateJob(url string, job Job) (resp JobResp, err error) {

	buf, err := ffjson.Marshal(&job)
	if err != nil {
		return
	}

	body, err := sendPost(url+"v1/jobs", string(buf))
	if err != nil {
		return
	}

	data := []byte(body)
	err = ffjson.Unmarshal(data, &resp)

	return
}

func DeleteJob(url string, jobName string) (resp JobResp, err error) {

	body, err := sendDelete(url+"v1/jobs/"+jobName, "")
	if err != nil {
		return
	}

	data := []byte(body)
	err = ffjson.Unmarshal(data, &resp)

	return
}

func ShowJob(url string, jobName string) (resp JobResp, err error) {

	body, err := send(url+"v1/jobs/"+jobName, "")
	if err != nil {
		return
	}

	data := []byte(body)
	err = ffjson.Unmarshal(data, &resp)

	return
}

func ShowJobs(url string) (resp []JobResp, err error) {

	body, err := send(url+"v1/jobs", "")
	if err != nil {
		return
	}

	data := []byte(body)
	err = ffjson.Unmarshal(data, &resp)

	return
}

func RunJob(url string, jobName string) (resp JobResp, err error) {

	body, err := sendPost(url+"v1/jobs/"+jobName, "")
	if err != nil {
		return
	}

	data := []byte(body)
	err = ffjson.Unmarshal(data, &resp)

	return
}

func ExecutionJobList(url string, jobName string) (resp []ExecutionJobResp, err error) {

	body, err := send(url+"v1/executions/"+jobName, "")
	if err != nil {
		return
	}

	data := []byte(body)
	err = ffjson.Unmarshal(data, &resp)

	return
}

func MemberList(url string) (resp []Member, err error) {

	body, err := send(url+"v1/members", "")
	if err != nil {
		return
	}

	data := []byte(body)
	err = ffjson.Unmarshal(data, &resp)

	return
}

func MemberLeader(url string) (resp Member, err error) {

	body, err := send(url+"v1/leader", "")
	if err != nil {
		return
	}

	data := []byte(body)
	err = ffjson.Unmarshal(data, &resp)

	return
}
