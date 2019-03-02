package params

// TeamJoinTeams 获取某用户所加入的群信息
type TeamJoinTeams struct {
	Accid string
}

func (p TeamJoinTeams) Format() map[string]string {
	return format(p)
}

func (p TeamJoinTeams) GetPath() string {
	return "/team/joinTeams.action"
}
