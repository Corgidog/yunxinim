package params

// TeamMute 修改消息提醒开关
type TeamMute struct {
	Tid   string
	Accid string
	Ope   string
}

func (p TeamMute) Format() map[string]string {
	return format(p)
}

func (p TeamMute) GetPath() string {
	return "/team/muteTeam.action"
}
