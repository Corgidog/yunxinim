package params

// UserCreateParams 用户创建参数
type UserCreateParams struct {
	Accid  string //账号id
	Name   string //昵称
	Props  string //json属性,自定义
	Icon   string //头像url
	Token  string //token
	Sign   string //签名
	Email  string
	Birth  string
	Mobile string
	Gender string //0未知,1男,2女
	Ex     string //扩展字段
}

func (p UserCreateParams) Format() map[string]string {
	return format(p)
}

func (p UserCreateParams) GetPath() string {
	return "/user/create.action"
}
