package params

// TeamAddManager 任命管理员
type TeamAddManager struct {
	Tid     string
	Owner   string
	Members string
}

func (p TeamAddManager) Format() map[string]string {
	return format(p)
}

func (p TeamAddManager) GetPath() string {
	return "/team/addManager.action"
}
