package params

// TeamLeave 主动退群
type TeamLeave struct {
	Tid   string
	Accid string
}

func (p TeamLeave) Format() map[string]string {
	return format(p)
}

func (p TeamLeave) GetPath() string {
	return "/team/leave.action"
}
