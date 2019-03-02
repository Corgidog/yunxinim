package params

// TeamAdd 拉人入群
type TeamAdd struct {
	Tid     string
	Owner   string
	Members string
	Magree  string
	Msg     string
	Attach  string
}

func (p TeamAdd) Format() map[string]string {
	return format(p)
}

func (p TeamAdd) GetPath() string {
	return "/team/add.action"
}
