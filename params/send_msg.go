package params

// SendMsgParams https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E5%8F%91%E9%80%81%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF
type SendMsgParams struct {
	From             string
	Ope              string // 0：点对点个人消息，1：群消息（高级群），其他返回414
	To               string // ope==0是表示accid即用户id，ope==1表示tid即群id
	Type             string
	Body             string
	Antispam         string
	AntispamCustom   string
	Option           string
	Pushcontent      string
	Payload          string
	Ext              string
	Forcepushlist    string
	Forcepushcontent string
	Forcepushall     string
	Bid              string
	UseYidun         string
	MarkRead         string
	CheckFriend      string
}

func (p SendMsgParams) Format() map[string]string {
	return format(p)
}

func (p SendMsgParams) GetPath() string {
	return "/msg/sendMsg.action"
}
