package maps

import (
	"errors"
	"strconv"
)

func Merge(maps ...*map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for _, t := range maps {
		for k, v := range *t {
			m[k] = v
		}
	}
	return m
}

func Bool(search *map[string]interface{}, fallback bool, key string) (bool, error) {
	if data, exists := (*search)[key]; exists {
		if b, ok := data.(bool); ok {
			return b, nil
		}
		if s, ok := data.(string); ok {
			b, err := strconv.ParseBool(s)
			if err != nil {
				return fallback, err
			}
			return b, nil
		}
	}
	return fallback, errors.New("Unable to find data by supplied key(s)")
}

func String(search *map[string]interface{}, fallback string, key string) (string, error) {
	if data, exists := (*search)[key]; exists {
		if s, ok := data.(string); ok {
			return s, nil
		} else {
			return fallback, errors.New("Unable to assert string type")
		}
	}
	return fallback, errors.New("Unable to find data by supplied key(s)")
}

func Int(search *map[string]interface{}, fallback int64, key string) (int64, error) {
	if data, exists := (*search)[key]; exists {
		if i, ok := data.(int64); ok {
			return i, nil
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

func Float(search *map[string]interface{}, fallback float64, key string) (float64, error) {
	if data, exists := (*search)[key]; exists {
		if f, ok := data.(float64); ok {
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
