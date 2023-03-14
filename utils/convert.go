package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

//IsASCIIUpper 是否是大写字母
func IsASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

//IsASCIIDigit 是否是数字
func IsASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

//CamelToLine 驼峰命名转下划线命名
func CamelToLine(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		t = append(t, 'X')
		i++
	}
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && IsASCIIUpper(s[i+1]) {
			continue
		}
		if IsASCIIDigit(c) {
			t = append(t, c)
			continue
		}

		if IsASCIIUpper(c) {
			c ^= ' '
		}
		t = append(t, c)

		for i+1 < len(s) && IsASCIIUpper(s[i+1]) {
			i++
			t = append(t, '_')
			t = append(t, bytes.ToLower([]byte{s[i]})[0])
		}
	}
	return string(t)
}

//StrToInt64 字符串转64位整型
func StrToInt64(s string) (i int64, err error) {
	i, err = strconv.ParseInt(s, 10, 64)
	return i, err
}

//StrToInt 字符串转整型
func StrToInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

//StrToFloat64 字符串转64位浮点型
func StrToFloat64(s string) (f float64, err error) {
	f, err = strconv.ParseFloat(s, 64)
	return f, err
}

//StrToByte 字符串转字节流
func StrToByte(s string) []byte {
	return []byte(s)
}

//IntToStr 整型转字符串
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

//IntToInt32 整型转32位整型
func IntToInt32(i int) int32 {
	return int32(i)
}

//IntToInt64 整型转64位整型
func IntToInt64(i int) int64 {
	return int64(i)
}

//Int32ToInt 32位整型转整型
func Int32ToInt(i int32) int {
	return int(i)
}

//Int32ToInt64 32位整型转64位整型
func Int32ToInt64(i int32) int64 {
	return int64(i)
}

//Int64ToInt 64位整型转整型
func Int64ToInt(i int64) int {
	return int(i)
}

//Int64ToInt32 64位整型转32位整型
func Int64ToInt32(i int64) int32 {
	return int32(i)
}

//Int64ToStr 64位整型转字符串
func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

//ToStr 字节流、数字、布尔转成字符串
func ToStr(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv := reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
}

//StructToMap 结构体转成集合
func StructToMap(s interface{}) map[string]interface{} {
	jsn, _ := json.Marshal(&s)
	maper := make(map[string]interface{})
	json.Unmarshal([]byte(jsn), &maper)
	return maper
}

//ToStruct 结构体转结构体
func ToStruct(source interface{}, target interface{}) {
	jsn, _ := json.Marshal(source)
	json.Unmarshal([]byte(jsn), target)
}

//MapToStr 集合转字符串
func MapToStr(s interface{}) string {
	jsn, _ := json.Marshal(&s)
	return string(jsn)
}

//StructToStr 结构体转成字符串
func StructToStr(s interface{}) string {
	jsn, _ := json.Marshal(&s)
	return string(jsn)
}
