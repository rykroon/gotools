package sqlx

import "reflect"

// StructToScanArgs takes a pointer to a struct and returns a list of pointers to the fields of the struct.
func StructToScanArgs(s any) []any {
	v := reflect.ValueOf(s).Elem()
	t := v.Type()
	n := t.NumField()
	result := make([]any, n)
	for i := range n {
		field := v.Field(i)
		result[i] = field.Addr().Interface()
	}
	return result
}

// FlattenArgs takes a list of arguments and flattens them into a single list of arguments.
func FlattenArgs(args ...any) []any {
	var r []any
	for _, arg := range args {
		switch typedArg := arg.(type) {
		case []any:
			r = append(r, typedArg...)
		case []int:
			for i := range typedArg {
				r = append(r, typedArg[i])
			}
		case []string:
			for i := range typedArg {
				r = append(r, typedArg[i])
			}
		case []float64:
			for i := range typedArg {
				r = append(r, typedArg[i])
			}
		case []bool:
			for i := range typedArg {
				r = append(r, typedArg[i])
			}

		case int, string, float64, bool:
			r = append(r, arg)
		default:
			panic("unsupported type")
		}
	}
	return r

}
