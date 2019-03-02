package params

// SendBatchAttachMsg 批量发送点对点自定义消息
type SendBatchAttachMsg struct {
	FromAccid   string
	ToAccids    string
	Attach      string
	PushContent string
	Payload     string
	Sound       string
	Save        string
	Option      string
}

func (p SendBatchAttachMsg) Format() map[string]string {
	return format(p)
}

func (p SendBatchAttachMsg) GetPath() string {
	return "/msg/sendBatchAttachMsg.action"
}
