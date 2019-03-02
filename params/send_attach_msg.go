package params

// SendAttachMsg 发送自定义系统通知
type SendAttachMsg struct {
	From        string
	MsgType     string
	To          string
	Attach      string
	Pushcontent string
	Payload     string
	Sound       string
	Save        string
	Option      string
}

func (p SendAttachMsg) Format() map[string]string {
	return format(p)
}

func (p SendAttachMsg) GetPath() string {
	return "/msg/sendAttachMsg.action"
}
