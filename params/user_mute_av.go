package params

// UserMute 账号全局禁言AV
type UserMuteAv struct {
	Accid string
	Mute  string //true：全局禁言，false:取消全局禁言
}

func (p UserMuteAv) Format() map[string]string {
	return format(p)
}

func (p UserMuteAv) GetPath() string {
	return "/user/muteAv.action"
}
