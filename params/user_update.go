package params

// UserUpdateParams 用户更新参数
type UserUpdateParams struct {
	Accid string //账号id
	Props string //json属性,自定义
	Token string //token
}

func (p UserUpdateParams) Format() map[string]string {
	return format(p)
}

func (p UserUpdateParams) GetPath() string {
	return "/user/update.action"
}
