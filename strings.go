package convert

import "errors"

func Strings(i interface{}) ([]string, error) {
	switch i.(type) {
	case [][]byte:
		result := []string{}
		s := i.([][]byte)
		for _, v := range s {
			result = append(result, string(v[:]))
		}
		return result, nil
	default:
		return nil, errors.New("invalid input value")
	}
}
