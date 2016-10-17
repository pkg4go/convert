package convert

import "errors"

func Strings(i interface{}) ([]string, error) {
	switch i.(type) {
	case []interface{}:
		result := []string{}
		s := i.([]interface{})
		for _, v := range s {
			result = append(result, String(v))
		}
		return result, nil
	case [][]byte:
		result := []string{}
		s := i.([][]byte)
		for _, v := range s {
			result = append(result, string(v[:]))
		}
		return result, nil
	case []string:
		s := i.([]string)
		return s, nil
	default:
		return nil, errors.New("invalid input value")
	}
}
