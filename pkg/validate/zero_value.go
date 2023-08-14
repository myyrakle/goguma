package validate

import "reflect"

// value 값과 pointer 값에 대해서 모두 Zero 값을 체크함
func IsZero[T any](x T) bool {
	typeInfo := reflect.TypeOf(x)
	valueInfo := reflect.ValueOf(x)
	isPointer := typeInfo.Kind() == reflect.Ptr

	method, hasIsZero := typeInfo.MethodByName("IsZero")

	// IsZero 메소드가 있는 경우에는 해당 메소드를 호출하여 결과를 반환
	if hasIsZero {
		result := method.Func.Call([]reflect.Value{valueInfo})

		if len(result) > 0 {
			return result[0].Bool()
		} else {
			panic("IsZero method must return bool")
		}
	}

	// 포인터인 경우에는 nil이면 zero value로 간주
	// 포인터가 아니면 기본값을 zero value로 간주
	if isPointer {
		if valueInfo.IsNil() {
			return true
		} else {
			return false
		}
	} else {
		return valueInfo.IsZero()
	}
}
