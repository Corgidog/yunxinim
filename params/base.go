package params

import "reflect"

// Param 参数接口
type Param interface {
	Format() map[string]string
	GetPath() string
}

func format(p interface{}) map[string]string {
	m := make(map[string]string)

	re := reflect.ValueOf(p)
	reType := re.Type()
	for i := 0; i < re.NumField(); i++ {
		if val := re.Field(i).String(); val != "" {
			m[firstToLower(reType.Field(i).Name)] = val
		}
	}
	return m
}

func firstToLower(str string) string {
	var lowerStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 96 {
				vv[i] += 32
				lowerStr += string(vv[i])
			} else {
				return str
			}
		} else {
			lowerStr += string(vv[i])
		}
	}
	return lowerStr
}
