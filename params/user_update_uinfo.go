package params

// UserUpdateUinfoParams 用户更新信息
type UserUpdateUinfoParams struct {
	Accid  string //账号id
	Name   string //昵称
	Icon   string //头像url
	Sign   string //签名
	Email  string
	Birth  string
	Mobile string
	Gender string //0未知,1男,2女
	Ex     string //扩展字段
}

func (p UserUpdateUinfoParams) Format() map[string]string {
	return format(p)
}

func (p UserUpdateUinfoParams) GetPath() string {
	return "/user/updateUinfo.action"
}
