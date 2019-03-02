package params

// ReCallMsgParams 撤回消息
type ReCallMsgParams struct {
	DeleteMsgid string
	Timetag     string
	Type        string
	From        string
	To          string
	Msg         string
	IgnoreTime  string
	PushContent string
	Payload     string
}

func (p ReCallMsgParams) Format() map[string]string {
	return format(p)
}

func (p ReCallMsgParams) GetPath() string {
	return "/msg/recall.action"
}
