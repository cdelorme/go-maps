package maps

import (
	"encoding/json"
	"errors"
	"strconv"
)

var errorNoKeys = errors.New("no key(s) supplied")
var errorNotFound = errors.New("no index found at supplied key(s)")
var errorCastFailed = errors.New("failed to cast data to expected type")

func Get(search map[string]interface{}, keys ...string) (interface{}, error) {
	if len(keys) == 0 {
		return nil, errorNoKeys
	} else if _, exists := search[keys[0]]; !exists {
		return nil, errorNotFound
	}

	if len(keys) > 1 {
		if m, ok := search[keys[0]].(map[string]interface{}); ok {
			return Get(m, keys[1:]...)
		}
	}
	return search[keys[0]], nil
}

func Bool(search map[string]interface{}, fallback bool, keys ...string) (bool, error) {
	d, err := Get(search, keys...)
	if err != nil {
		return fallback, err
	}

	switch b := d.(type) {
	default:
		return true, nil
	case bool:
		return b, nil
	case int:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case int8:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case int32:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case int64:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case uint:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case uint8:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case uint32:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case uint64:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case float32:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case float64:
		if b == 0 {
			return false, nil
		}
		return true, nil
	case string:
		if len(b) == 0 || b == "false" || b == "0" || b == "null" || b == "undefined" || b == "nil" {
			return false, nil
		}
		return true, nil
	}
}

func String(search map[string]interface{}, fallback string, keys ...string) (string, error) {
	d, err := Get(search, keys...)
	if err != nil {
		return fallback, err
	}

	switch s := d.(type) {
	default:
		return fallback, errorCastFailed
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case int:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case int32:
		return strconv.FormatInt(int64(s), 10), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(s, 10), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	}
}

func Int(search map[string]interface{}, fallback int64, keys ...string) (int64, error) {
	d, err := Get(search, keys...)
	if err != nil {
		return fallback, err
	}

	switch i := d.(type) {
	default:
		return fallback, errorCastFailed
	case int:
		return int64(i), nil
	case int8:
		return int64(i), nil
	case int32:
		return int64(i), nil
	case int64:
		return i, nil
	case uint:
		return int64(i), nil
	case uint8:
		return int64(i), nil
	case uint32:
		return int64(i), nil
	case uint64:
		return int64(i), nil
	case float32:
		return int64(i), nil
	case float64:
		return int64(i), nil
	case string:
		s, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return fallback, errorCastFailed
		}
		return int64(s), nil
	}
}

func Float(search map[string]interface{}, fallback float64, keys ...string) (float64, error) {
	d, err := Get(search, keys...)
	if err != nil {
		return fallback, err
	}

	switch f := d.(type) {
	default:
		return fallback, errorCastFailed
	case int:
		return float64(f), nil
	case int8:
		return float64(f), nil
	case int32:
		return float64(f), nil
	case int64:
		return float64(f), nil
	case uint:
		return float64(f), nil
	case uint8:
		return float64(f), nil
	case uint32:
		return float64(f), nil
	case uint64:
		return float64(f), nil
	case float32:
		return float64(f), nil
	case float64:
		return f, nil
	case string:
		s, err := strconv.ParseFloat(f, 64)
		if err != nil {
			return fallback, errorCastFailed
		}
		return s, nil
	}
}

func Merge(maps ...map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for _, t := range maps {
		for k, v := range t {
			if _, me := m[k]; me {
				if m1, ok := m[k].(map[string]interface{}); ok {
					if m2, is := v.(map[string]interface{}); is {
						v = Merge(m1, m2)
					}
				}
			}
			m[k] = v
		}
	}
	return m
}

func To(obj interface{}, data ...map[string]interface{}) {
	j, _ := json.Marshal(Merge(data...))
	json.Unmarshal(j, &obj)
}
