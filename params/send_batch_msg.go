package params

// SendBatchMsg 批量发送点对点普通消息
type SendBatchMsg struct {
	FromAccid   string
	ToAccids    string
	Type        string
	Body        string
	Option      string
	PushContent string
	Payload     string
	Ext         string
	Bid         string
	UseYidun    string
	ReturnMsgid string
}

func (p SendBatchMsg) Format() map[string]string {
	return format(p)
}

func (p SendBatchMsg) GetPath() string {
	return "/msg/sendBatchMsg.action"
}
