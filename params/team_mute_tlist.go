package params

// TeamMuteMember 禁言群成员
type TeamMuteMember struct {
	Tid   string
	Owner string
	Accid string
	Mute  string
}

func (p TeamMuteMember) Format() map[string]string {
	return format(p)
}

func (p TeamMuteMember) GetPath() string {
	return "/team/muteTlist.action"
}
