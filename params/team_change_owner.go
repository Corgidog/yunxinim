package params

// TeamChangeOwner 移交群主
type TeamChangeOwner struct {
	Tid      string
	Owner    string
	Newowner string
	Leave    string
}

func (p TeamChangeOwner) Format() map[string]string {
	return format(p)
}

func (p TeamChangeOwner) GetPath() string {
	return "/team/changeOwner.action"
}
