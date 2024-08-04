package binance_connector

import "fmt"

func ToInt64(digit interface{}) (i int64, err error) {
	if intVal, ok := digit.(int); ok {
		return int64(intVal), nil
	}
	if floatVal, ok := digit.(float64); ok {
		return int64(floatVal), nil
	}
	return 0, fmt.Errorf("unexpected digit: %v", digit)
}
