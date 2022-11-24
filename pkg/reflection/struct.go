package reflection

import "reflect"

func ConvStructToMap(value any) map[string]any {
	refType := reflect.TypeOf(value)
	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
	}
	if refType.Kind() != reflect.Struct {
		return nil
	}
	fieldNum := refType.NumField()
	result := make(map[string]any)
	nameList := make([]string, 0)
	for i := 0; i < fieldNum; i++ {
		nameList = append(nameList, refType.Field(i).Name)
	}
	refValue := reflect.ValueOf(value)
	if refValue.Kind() == reflect.Ptr {
		refValue = refValue.Elem()
	}
	for i := 0; i < fieldNum; i++ {
		result[nameList[i]] = refValue.FieldByName(nameList[i]).Interface()
	}
	return result
}