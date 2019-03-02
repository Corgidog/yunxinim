package params

type BroadcastMsg struct {
	Body      string
	From      string
	IsOffline string
	Ttl       string
	TargetOs  string
}

func (p BroadcastMsg) Format() map[string]string {
	return format(p)
}

func (p BroadcastMsg) GetPath() string {
	return "/msg/broadcastMsg.action"
}
