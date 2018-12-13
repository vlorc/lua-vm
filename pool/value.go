package pool

import (
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
	"reflect"
	"strconv"
	"unicode/utf8"
)

type __type struct {
	src reflect.Type
	dst lua.LValueType
}

var TypeTable = map[reflect.Type]func(interface{}, lua.LValue){
	reflect.TypeOf(""): func(dst interface{}, src lua.LValue) {
		*(dst.(*string)) = src.String()
	},
	reflect.TypeOf(true): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNil:
			*(dst.(*bool)) = false
		case lua.LTString:
			*(dst.(*bool)), _ = strconv.ParseBool(string(src.(lua.LString)))
		default:
			*(dst.(*bool)) = true
		}
	},
	reflect.TypeOf(float64(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*float64)) = float64(src.(lua.LNumber))
		case lua.LTString:
			*(dst.(*float64)), _ = strconv.ParseFloat(string(src.(lua.LString)), 64)
		default:
			*(dst.(*float64)) = 0
		}
	},
	reflect.TypeOf(float32(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*float32)) = float32(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseFloat(string(src.(lua.LString)), 32)
			*(dst.(*float32)) = float32(v)
		default:
			*(dst.(*float32)) = 0
		}
	},
	reflect.TypeOf(int(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*int)) = int(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseInt(string(src.(lua.LString)), 0, 63)
			*(dst.(*int)) = int(v)
		default:
			*(dst.(*int)) = 0
		}
	},
	reflect.TypeOf(int8(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*int8)) = int8(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseInt(string(src.(lua.LString)), 0, 7)
			*(dst.(*int8)) = int8(v)
		default:
			*(dst.(*int8)) = 0
		}
	},
	reflect.TypeOf(int16(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*int16)) = int16(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseInt(string(src.(lua.LString)), 0, 15)
			*(dst.(*int16)) = int16(v)
		default:
			*(dst.(*int16)) = 0
		}
	},
	reflect.TypeOf(int32(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*int32)) = int32(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseInt(string(src.(lua.LString)), 0, 31)
			*(dst.(*int32)) = int32(v)
		default:
			*(dst.(*int32)) = 0
		}
	},
	reflect.TypeOf(int64(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*int64)) = int64(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseInt(string(src.(lua.LString)), 0, 63)
			*(dst.(*int64)) = int64(v)
		default:
			*(dst.(*int64)) = 0
		}
	},
	reflect.TypeOf(rune(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*rune)) = rune(src.(lua.LNumber))
		case lua.LTString:
			if s := string(src.(lua.LString)); len(s) > 0 {
				*(dst.(*rune)), _ = utf8.DecodeRuneInString(s)
			}
		default:
			*(dst.(*rune)) = 0
		}
	},
	reflect.TypeOf(uint(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*uint)) = uint(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseUint(string(src.(lua.LString)), 0, 64)
			*(dst.(*uint)) = uint(v)
		default:
			*(dst.(*uint)) = 0
		}
	},
	reflect.TypeOf(byte(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*byte)) = byte(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseUint(string(src.(lua.LString)), 0, 8)
			*(dst.(*byte)) = byte(v)
		default:
			*(dst.(*byte)) = 0
		}
	},
	reflect.TypeOf(uint8(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*uint8)) = uint8(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseUint(string(src.(lua.LString)), 0, 8)
			*(dst.(*uint8)) = uint8(v)
		default:
			*(dst.(*uint8)) = 0
		}
	},
	reflect.TypeOf(uint16(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*uint16)) = uint16(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseUint(string(src.(lua.LString)), 0, 16)
			*(dst.(*uint16)) = uint16(v)
		default:
			*(dst.(*uint16)) = 0
		}
	},
	reflect.TypeOf(uint32(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*uint32)) = uint32(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseUint(string(src.(lua.LString)), 0, 32)
			*(dst.(*uint32)) = uint32(v)
		default:
			*(dst.(*uint32)) = 0
		}
	},
	reflect.TypeOf(uint64(0)): func(dst interface{}, src lua.LValue) {
		switch src.Type() {
		case lua.LTNumber:
			*(dst.(*uint64)) = uint64(src.(lua.LNumber))
		case lua.LTString:
			v, _ := strconv.ParseUint(string(src.(lua.LString)), 0, 64)
			*(dst.(*uint64)) = uint64(v)
		default:
			*(dst.(*uint64)) = 0
		}
	},
}

func __toValue(value lua.LValue, result interface{}) bool {
	val := reflect.Indirect(reflect.ValueOf(result))
	if !val.CanSet() {
		return false
	}
	convert := TypeTable[val.Type()]
	if nil != convert {
		convert(result, value)
		return true
	}

	if lua.LTTable == value.Type() {
		if reflect.Map == val.Kind() || reflect.Struct == val.Kind() || reflect.Slice == val.Kind() {
			gluamapper.Map(value.(*lua.LTable), result)
		}
		return true
	}
	src := gluamapper.ToGoValue(value, gluamapper.Option{})
	if val.Type().ConvertibleTo(reflect.TypeOf(src)) {
		val.Set(reflect.ValueOf(src).Convert(val.Type()))
		return true
	}
	return false
}
