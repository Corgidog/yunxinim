package params

// TeamDetail 获取群组详细信息
type TeamDetail struct {
	Tid string
}

func (p TeamDetail) Format() map[string]string {
	return format(p)
}

func (p TeamDetail) GetPath() string {
	return "/team/queryDetail.action"
}
