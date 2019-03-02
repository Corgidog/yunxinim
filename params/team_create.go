package params

// TeamCreate 创建群聊
type TeamCreate struct {
	Tname           string
	Owner           string
	Members         string
	Announcement    string
	Intro           string
	Msg             string
	Magree          string
	Joinmode        string
	Custom          string
	Icon            string
	Beinvitemode    string
	Invitemode      string
	Uptinfomode     string
	Upcustommode    string
	TeamMemberLimit string
}

func (p TeamCreate) Format() map[string]string {
	return format(p)
}

func (p TeamCreate) GetPath() string {
	return "/team/create.action"
}
