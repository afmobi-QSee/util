/* 绝密 TOP SECRET, COPYRIGHT © AFMOBI GROUP */
package util

import (
	"reflect"
	"strconv"
	"github.com/kataras/go-errors"
	"regexp"
	"fmt"
)

/**
	校验参数方法ValidParams
	列子如下：
type TestStruct struct {
	A  string `NotNull:"true"`
	B  string `Min:"5"`
	B1 uint8 `Min:"5"`
	B2 float32 `Min:"5"`
	C  int `Min:"2" Max:"9"`
	D  string `SizeMin:"5"`
	E  string `SizeMin:"5" SizeMax:"10"`
	F  string `Pattern:"abc|def"`
	F1 string `Pattern:"[\u4e00-\u9fa5]"`
	H  *TestStruct1 `InStruct:"true"`
}

type TestStruct1 struct {
	I string `NotNull:"true"`
	J int `Min:"5"`
}

	该方法基于struct的tag属性，利用反射功能自动校验每个参数，一些常用tag设计如下
	NotNull   参数不为空,只支持string, true为校验,其他值或没有为不校验
	Min,Max   数字不大于或不小于某个值,支持int,uint,float,string类型, 后面的值为其限定范围
	SizeMin,SizeMax 字符长度不大于或不小于某个值,只支持string, 后面的值为其限定范围
	Pattern   用正则表达式匹配参数值,支持string,int类型,后面的值为校验的正则表达式
	InStruct  表示该参数为嵌套的struct,且需要校验里面的值,true为需要校验,其他值不用校验
 */

const (
	ValidOk string = "ok"
)

func ValidParams(p interface{}) string {
	v := ValidOk
	f := reflect.ValueOf(p)
	if f.MethodByName("Customization").Kind() == reflect.Invalid{
		fe := reflect.ValueOf(p).Elem()
		typeOfF := fe.Type()
		for i := 0; i < fe.NumField(); i++ {
			v := validSingleParam(typeOfF.Field(i), fe.Field(i).Interface())
			if v != ValidOk {
				return v
			}
		}
	}else {
		v1 := f.MethodByName("Customization").Call(nil)
		v = fmt.Sprint(v1[0])
	}
	return v
}

func validSingleParam(f reflect.StructField, v interface{}) string {
	t := f.Tag
	if t.Get("InStruct") == "true" {
		result := ValidParams(v)
		if result != "ok" {
			return result
		}
	}
	if t.Get("NotNull") == "true" {
		if EmptyStr(v.(string)) {
			return f.Name + " is null"
		}
	}
	if value := t.Get("Min"); value != "" {
		intV, intT, err := transferInt(value, v)
		if err != nil {
			return f.Name + err.Error()
		}
		if intV < intT {
			return f.Name + " is less than " + value
		}
	}
	if value := t.Get("Max"); value != "" {
		intV, intT, err := transferInt(value, v)
		if err != nil {
			return f.Name + err.Error()
		}
		if intV > intT {
			return f.Name + " is more than " + value
		}
	}
	if value := t.Get("SizeMin"); value != "" {
		intT, err := strconv.Atoi(value)
		if err != nil {
			return f.Name + " tag is not int"
		}

		lenV, err := lenStr(v)
		if err != nil {
			return f.Name + err.Error()
		}

		if lenV < intT {
			return f.Name + "'s length is less than " + value
		}
	}
	if value := t.Get("SizeMax"); value != "" {
		intT, err := strconv.Atoi(value)
		if err != nil {
			return f.Name + " tag is not int"
		}

		lenV, err := lenStr(v)
		if err != nil {
			return f.Name + err.Error()
		}

		if lenV > intT {
			return f.Name + "'s length is more than " + value
		}
	}
	if value := t.Get("Pattern"); value != "" {
		r, _ := regexp.Compile(value)
		if !r.Match([]byte(fmt.Sprint(v))) {
			return f.Name + " not match with pattern"
		}
	}
	return ValidOk
}

func transferInt(value string, v interface{}) (int, int, error) {
	intT, err := strconv.Atoi(value)
	if err != nil {
		return 0, 0, errors.New(" tag is not int")
	}

	var errv error
	var intV int
	switch s := v.(type) {
	case int:
		intV = s
	case int8:
		intV = int(s)
	case int16:
		intV = int(s)
	case int32:
		intV = int(s)
	case int64:
		intV = int(s)
	case byte:
		intV = int(s)
	case uint16:
		intV = int(s)
	case uint32:
		intV = int(s)
	case uint64:
		intV = int(s)
	case float32:
		intV = int(s)
	case float64:
		intV = int(s)
	case string:
		c, err := strconv.Atoi(s)
		if err != nil {
			errv = errors.New(" is not int")
		}
		intV = c
	default:
		errv = errors.New(" not support type")
	}
	return intV, intT, errv
}

func lenStr(v interface{}) (int, error) {
	switch s := v.(type) {
	case string:
		return len(s), nil
	default:
		return 0, errors.New(" not support type")
	}
}