package params

// TeamListMute 获取群组禁言列表
type TeamListMute struct {
	Tid   string
	Owner string
}

func (p TeamListMute) Format() map[string]string {
	return format(p)
}

func (p TeamListMute) GetPath() string {
	return "/team/listTeamMute.action"
}
