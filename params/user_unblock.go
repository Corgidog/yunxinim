package params

// UserUnBlockParams 解封用户参数
type UserUnBlockParams struct {
	Accid string //账号id
}

func (p UserUnBlockParams) Format() map[string]string {
	return format(p)
}

func (p UserUnBlockParams) GetPath() string {
	return "/user/unblock.action"
}
