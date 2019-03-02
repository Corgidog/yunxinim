package params

// TeamKick 踢人出群
type TeamKick struct {
	Tid     string
	Owner   string
	Member  string
	Members string
	Attach  string
}

func (p TeamKick) Format() map[string]string {
	return format(p)
}

func (p TeamKick) GetPath() string {
	return "/team/kick.action"
}
