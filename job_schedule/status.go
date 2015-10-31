package job_schedule

type StatusAgentResp struct {
	Bankend string `json:"backend"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type StatusSerfResp struct {
	Encrypted   string `json:"encrypted"`
	EventQueue  string `json:"event_queue"`
	EventTime   string `json:"event_time"`
	Failed      string `json:"failed"`
	IntentQueue string `json:"intent_queue"`
	Left        string `json:"left"`
	MemberTime  string `json:"member_time"`
	Members     string `json:"members"`
	QueryQueue  string `json:"query_queue"`
	QueryTime   string `json:"query_time"`
}

type StatusTagsResp struct {
	Datacenter string `json:"datacenter"`
	Key        string `json:"key"`
	Role       string `json:"role"`
	Server     string `json:"server"`
}

type StatusResp struct {
	Agent StatusAgentResp `json:"agent"`
	Serf  StatusSerfResp  `json:"serf"`
	Tags  StatusTagsResp  `json:"tags"`
}
