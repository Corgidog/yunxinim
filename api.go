package yunxinim

import (
	"log"

	"github.com/Corgidog/yunxinim/params"
)

type Sdk struct {
	urlPrefix string
	appKey    string
	appSecret string
}

func New(appKey, appSecret string) *Sdk {
	return &Sdk{
		urlPrefix: "https://api.netease.im/nimserver",
		appKey:    appKey,
		appSecret: appSecret,
	}
}

func (sdk *Sdk) request(path string, posts map[string]string) []byte {
	body, err := sdk.httpDo(path, posts)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (sdk *Sdk) Call(p params.Param) []byte {
	return sdk.request(p.GetPath(), p.Format())
}
