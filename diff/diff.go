package diff

import(
	"reflect"
)

func Diff(first map[string]interface{}, second map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k := range first {
		if _, ok := second[k]; !ok {
			result[k] = "undefined"
		} else if reflect.TypeOf(first[k]) != reflect.TypeOf(second[k]) {
			result[k] = second[k]
		} else {
			switch first[k].(type){
			default:
				if(first[k] != second[k]) {
					result[k] = second[k]
				}
			case map[string]interface{}:
				subResult := Diff(first[k].(map[string]interface{}), second[k].(map[string]interface{}))
				if len(subResult) != 0 {
					result[k] = subResult
				}
			case []interface{}:
				if !reflect.DeepEqual(first[k], second[k]) {
					result[k] = second[k]
				}
			}
		}
	}
	for k := range second {
		if _, ok := first[k]; !ok {
			result[k] = second[k]
		}
	}
	
	return result
}