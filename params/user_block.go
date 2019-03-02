package params

// UserBlockParams 封禁用户参数
type UserBlockParams struct {
	Accid    string //账号id
	Needkick string // 是否踢掉被禁用户，true或false，默认false
}

func (p UserBlockParams) Format() map[string]string {
	return format(p)
}

func (p UserBlockParams) GetPath() string {
	return "/user/block.action"
}
