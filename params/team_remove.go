package params

// TeamRemove 解散群
type TeamRemove struct {
	Tid   string
	Owner string
}

func (p TeamRemove) Format() map[string]string {
	return format(p)
}

func (p TeamRemove) GetPath() string {
	return "/team/remove.action"
}
