package maps

import (
	"errors"
	"strconv"
)

func Merge(maps ...*map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for _, t := range maps {
		for k, v := range *t {
			if _, me := m[k]; me {
				if m1, ok := m[k].(map[string]interface{}); ok {
					if m2, is := v.(map[string]interface{}); is {
						v = Merge(&m1, &m2)
					}
				}
			}
			m[k] = v
		}
	}
	return m
}

func Bool(search *map[string]interface{}, fallback bool, keys ...string) (bool, error) {
	if len(keys) == 0 {
		return fallback, errors.New("No keys supplied")
	}
	if data, exists := (*search)[keys[0]]; exists {
		if m, ok := data.(map[string]interface{}); ok {
			return Bool(&m, fallback, keys[1:]...)
		}
		if b, ok := data.(bool); ok {
			return b, nil
		}
		return true, nil
	}
	return fallback, errors.New("Unable to find data by supplied key(s)")
}

func String(search *map[string]interface{}, fallback string, keys ...string) (string, error) {
	if len(keys) == 0 {
		return fallback, errors.New("No keys supplied")
	}
	if data, exists := (*search)[keys[0]]; exists {
		if m, ok := data.(map[string]interface{}); ok {
			return String(&m, fallback, keys[1:]...)
		}
		if s, ok := data.(string); ok {
			return s, nil
		}
		switch t := data.(type) {
		default:
			return fallback, errors.New("Unable to cast data to string")
		case bool:
			return strconv.FormatBool(t), nil
		case int:
			return strconv.FormatInt(int64(t), 10), nil
		case int8:
			return strconv.FormatInt(int64(t), 10), nil
		case int32:
			return strconv.FormatInt(int64(t), 10), nil
		case int64:
			return strconv.FormatInt(t, 10), nil
		case float32:
			return strconv.FormatFloat(float64(t), 'f', -1, 64), nil
		case float64:
			return strconv.FormatFloat(t, 'f', -1, 64), nil
		}
	}
	return fallback, errors.New("Unable to find data by supplied key(s)")
}

func Int(search *map[string]interface{}, fallback int64, keys ...string) (int64, error) {
	if len(keys) == 0 {
		return fallback, errors.New("No keys supplied")
	}
	if data, exists := (*search)[keys[0]]; exists {
		if m, ok := data.(map[string]interface{}); ok {
			return Int(&m, fallback, keys[1:]...)
		}
		if i, ok := data.(int64); ok {
			return i, nil
		} else if i, ok := data.(int); ok {
			return int64(i), nil
		}
		if s, ok := data.(string); ok {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return fallback, err
			}
			return i, nil
		}
	}
	return fallback, errors.New("Unable to find data by supplied key(s)")
}

func Float(search *map[string]interface{}, fallback float64, keys ...string) (float64, error) {
	if len(keys) == 0 {
		return fallback, errors.New("No keys supplied")
	}
	if data, exists := (*search)[keys[0]]; exists {
		if m, ok := data.(map[string]interface{}); ok {
			return Float(&m, fallback, keys[1:]...)
		}
		if f, ok := data.(float32); ok {
			return float64(f), nil
		} else if f, ok := data.(float64); ok {
			return f, nil
		}
		if s, ok := data.(string); ok {
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return fallback, err
			}
			return f, nil
		}
	}
	return fallback, errors.New("Unable to find data by supplied key(s)")
}
