package params

// TeamGetMarkReadInfo 获取群组已读消息的已读详情信息
type TeamGetMarkReadInfo struct {
	Tid       string
	Msgid     string
	FromAccid string
	Snapshot  string
}

func (p TeamGetMarkReadInfo) Format() map[string]string {
	return format(p)
}

func (p TeamGetMarkReadInfo) GetPath() string {
	return "/team/getMarkReadInfo.action"
}
