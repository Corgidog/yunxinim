package yunxinim

import (
	"crypto/sha1"
)

var hexDigits = []rune{'0', '1', '2', '3', '4', '5',
	'6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

// GetCheckSum 取签名
func getCheckSum(appSecret, nonce, curTime string) string {
	return encode("sha1", appSecret+nonce+curTime)
}

func encode(encryptType, body string) string {
	var encryResult []byte
	switch encryptType {
	case "sha1":
		e := sha1.New()
		e.Write([]byte(body))
		encryResult = e.Sum(nil)
	}

	return getFormattedText(encryResult)
}

func getFormattedText(bodyBytes []byte) string {
	len := len(bodyBytes)
	bytes := make([]byte, len*2)

	for i := 0; i < len; i++ {
		bytes[i*2] = byte(hexDigits[(bodyBytes[i]>>4)&0x0f])
		bytes[(i*2)+1] = byte(hexDigits[bodyBytes[i]&0x0f])
	}

	return string(bytes)
}
