package params

// TeamRemoveManager 移除管理员
type TeamRemoveManager struct {
	Tid     string
	Owner   string
	Members string
}

func (p TeamRemoveManager) Format() map[string]string {
	return format(p)
}

func (p TeamRemoveManager) GetPath() string {
	return "/team/removeManager.action"
}
