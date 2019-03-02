package params

// TeamQuery 群信息与成员列表查询
type TeamQuery struct {
	Tids string
	Ope  string
}

func (p TeamQuery) Format() map[string]string {
	return format(p)
}

func (p TeamQuery) GetPath() string {
	return "/team/query.action"
}
