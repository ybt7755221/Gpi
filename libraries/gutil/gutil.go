package gutil

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

//结构体赋值结构体
func BeanUtil(out interface{}, in interface{}) {
	outType := reflect.TypeOf(out).Elem()
	outValue := reflect.ValueOf(out).Elem()
	inType := reflect.TypeOf(in).Elem()
	inValue := reflect.ValueOf(in).Elem()
	outNum := outType.NumField()
	for i := 0; i < outNum; i++ {
		outFieldInfo := outType.Field(i)
		inTypeInfo, ok := inType.FieldByName(outFieldInfo.Name)
		if ok {
			outType := outFieldInfo.Type.String()
			inType := inTypeInfo.Type.String()
			inVal := inValue.FieldByName(outFieldInfo.Name)
			if outType == inType {
				outValue.FieldByName(outFieldInfo.Name).Set(reflect.Value(inVal))
			} else {
				var val interface{}
				switch outType {
				case "int":
					if inType == "string" {
						val, _ = strconv.Atoi(inVal.String())
					} else {
						val = int(inVal.Int())
					}
				case "int32":
					if inType == "string" {
						val, _ = strconv.ParseInt(inVal.String(), 10, 32)
					} else {
						val = int32(inVal.Int())
					}
				case "int64":
					if inType == "string" {
						val, _ = strconv.ParseInt(inVal.String(), 10, 64)
					} else {
						val = int64(inVal.Int())
					}
				case "string":
					if inType == "time.Time" {
						val = inVal.Interface().(time.Time).Format("2006-01-02 15:04:05")
					} else {
						val = inVal.String()
					}
				case "float32":
					val, _ = strconv.ParseFloat(inVal.String(), 32)
				case "float64":
					val, _ = strconv.ParseFloat(inVal.String(), 32)
				case "time.Time":
					tmpValue := inVal.String()
					if len(tmpValue) == 10 && strings.Index(tmpValue, "-") == -1 {
						intTm, _ := strconv.ParseInt(tmpValue, 10, 64)
						tm := time.Unix(intTm, 0)
						tmpValue = tm.Format("2006-01-02 15:04:05")
					}
					val, _ = time.Parse("2006-01-02 15:04:05", tmpValue)
				default:
					val = nil
				}
				outValue.FieldByName(outFieldInfo.Name).Set(reflect.ValueOf(val))
			}
		}
	}
}

//首字母小写
func FirstToLower(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

//首字母大写
func FirstToUpper(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
