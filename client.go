package yunxinim

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (self *Sdk) buildHeader() http.Header {
	var h = http.Header{}
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := TimeUUID().String()
	checkSum := getCheckSum(self.appSecret, nonce, curTime)
	//log.Fatal(checkSum)

	h.Set("Content-Type", "application/x-www-form-urlencoded")
	h.Set("AppKey", self.appKey)
	h.Set("Nonce", nonce)
	h.Set("CurTime", curTime)
	h.Set("CheckSum", checkSum)

	return h
}

func (self *Sdk) httpDo(path string, posts map[string]string) ([]byte, error) {
	client := &http.Client{}

	params := url.Values{}
	for key, val := range posts {
		params.Add(key, val)
	}

	reqURL := self.urlPrefix + path

	req, err := http.NewRequest("POST", reqURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header = self.buildHeader()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
