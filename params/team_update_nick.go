package params

// TeamUpdateNick 修改群昵称
type TeamUpdateNick struct {
	Tid    string
	Owner  string
	Accid  string
	Nick   string
	Custom string
}

func (p TeamUpdateNick) Format() map[string]string {
	return format(p)
}

func (p TeamUpdateNick) GetPath() string {
	return "/team/updateTeamNick.action"
}
