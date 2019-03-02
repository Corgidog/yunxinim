package params

//TeamUpdate 群信息更新
type TeamUpdate struct {
	Tid             string
	Tname           string
	Owner           string
	Announcement    string
	Intro           string
	Joinmode        string
	Custom          string
	Icon            string
	Beinvitemode    string
	Invitemode      string
	Uptinfomode     string
	Upcustommo      string
	TeamMemberLimit string
}

func (p TeamUpdate) Format() map[string]string {
	return format(p)
}

func (p TeamUpdate) GetPath() string {
	return "/team/update.action"
}
