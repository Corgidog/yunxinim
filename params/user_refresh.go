package params

// UserRefreshTokenParams 用户刷新token参数
type UserRefreshTokenParams struct {
	Accid string //账号id
}

func (p UserRefreshTokenParams) Format() map[string]string {
	return format(p)
}

func (p UserRefreshTokenParams) GetPath() string {
	return "/user/refreshToken.action"
}
