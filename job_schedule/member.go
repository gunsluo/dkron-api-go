package job_schedule

type Member struct {
	Name        string  `json:"name"`
	Addr        string  `json:"Addr"`
	Port        int     `json:"Port"`
	Status      int     `json:"Status"`
	ProtocolMin int     `json:"ProtocolMin"`
	ProtocolMax int     `json:"ProtocolMax"`
	ProtocolCur int     `json:"ProtocolCur"`
	DelegateMin int     `json:"DelegateMin"`
	DelegateMax int     `json:"DelegateMax"`
	DelegateCur int     `json:"DelegateCur"`
	Tags        JobTags `json:"Tags"`
}
