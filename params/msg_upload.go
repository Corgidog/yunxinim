package params

// 文件上传
type MsgUpload struct {
	Content   string
	Type      string
	Ishttps   string
	ExpireSec string
	Tag       string
}

func (p MsgUpload) Format() map[string]string {
	return format(p)
}

func (p MsgUpload) GetPath() string {
	return "/msg/upload.action"
}
