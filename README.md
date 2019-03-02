# yunxinim
网易云信IM的golang版sdk

## 使用

```
appKey := "your appkey"
appSecret := "your appsecret"

sdk := yunxinim.New(appKey, appSecret)

// 取用户信息
getuinfos := params.UserGetUinfosParams{
    Accids: `["aaaaa"]`,
}

log.Println(string(sdk.Call(getuinfos)))

// 发送普通消息
sendmsg := params.SendMsgParams{
    From: "aaaaa",
    Ope:  "0",
    To:   "bbbbb",
    Type: "0",
    Body: `{"msg":"woaini"}`,
}
log.Println(string(sdk.Call(sendmsg)))
```
