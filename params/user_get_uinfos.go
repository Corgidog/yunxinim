package params

// UserGetUinfosParams 获取用户信息
type UserGetUinfosParams struct {
	Accids string //账号id
}

func (p UserGetUinfosParams) Format() map[string]string {
	return format(p)
}

func (p UserGetUinfosParams) GetPath() string {
	return "/user/getUinfos.action"
}
