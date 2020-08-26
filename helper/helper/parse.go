package helper

import (
	"encoding/json"
	"fmt"

	"github.com/mozillazg/go-pinyin"
)

func GetPinYin(str string) (py []string) {
	return pinyin.LazyConvert(str, nil)
}

func SliceCapitalize(str []string) string {
	res := ""
	for _, v := range str {
		res += Capitalize(v)
	}
	return res
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		}
		// 需要拼接后面的取消屏蔽
		// else {
		//	upperStr += string(vv[i])
		//}
	}
	return upperStr
}

func StructToMapViaJson(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(s)
	json.Unmarshal(j, &m)

	return m
}

func Translate(s interface{}) map[string]interface{} {
	data := StructToMapViaJson(s)
	if m, ok := data["contract_start"]; ok {
		contractStart, _ := UTCString2Time(m.(string))
		data["contract_start"] = Time2String(contractStart)
	}
	if m, ok := data["contract_end"]; ok {
		contractEnd, _ := UTCString2Time(m.(string))
		data["contract_end"] = Time2String(contractEnd)
	}

	createdAt, _ := UTCString2Time(data["created_at"].(string))
	updatedAt, _ := UTCString2Time(data["created_at"].(string))
	data["created_at"] = Time2String(createdAt)
	data["updated_at"] = Time2String(updatedAt)
	return data
}
