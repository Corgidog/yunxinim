package params

// TeamMuteTlistAll 将群组整体禁言
type TeamMuteTlistAll struct {
	Tid      string
	Owner    string
	Mute     string
	MuteType string
}

func (p TeamMuteTlistAll) Format() map[string]string {
	return format(p)
}

func (p TeamMuteTlistAll) GetPath() string {
	return "/team/muteTlistAll.action"
}
