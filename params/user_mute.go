package params

// UserMute 账号全局禁言
type UserMute struct {
	Accid string
	Mute  string
}

func (p UserMute) Format() map[string]string {
	return format(p)
}

func (p UserMute) GetPath() string {
	return "/user/mute.action"
}
