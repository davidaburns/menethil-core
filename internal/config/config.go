package config

import (
	"fmt"
	"strconv"

	"github.com/davidaburns/menethil-core/internal/errors"
	"gopkg.in/ini.v1"
)


type Config struct {
	handle *ini.File
}

type ConfigSection string
type ConfigKey string
type ConfigSupportedValueTypes interface {
	~string | ~int | ~int64 | ~bool | ~float64
}

func LoadConfig(path string) (*Config, error) {
	handle, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	return &Config{handle: handle}, nil
}

func GetValue[T ConfigSupportedValueTypes](c *Config, section ConfigSection, key ConfigKey) (T, error) {
	var zero T
	sectionStr := string(section)
	keyStr := string(key)

	if !c.handle.HasSection(sectionStr) {
		return zero, fmt.Errorf("%v: %v", errors.ConfigSectionNotFound, sectionStr)
	}

	sec := c.handle.Section(sectionStr)
	if !sec.HasKey(keyStr) {
		return zero, fmt.Errorf("%v: %v", errors.ConfigKeyNotFound, keyStr)
	}

	str := sec.Key(keyStr).String()
	switch any(zero).(type) {
	case string:
		return any(str).(T), nil

	case int, int64:
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return zero, fmt.Errorf("error converting '%s' to integer: %w", str, err)
		}

		if _, ok := any(zero).(int); ok {
			return any(int(i)).(T), nil
		}

		return any(i).(T), nil

	case bool:
		b, err := strconv.ParseBool(str)
		if err != nil {
			return zero, fmt.Errorf("error converting '%s' to bool: %w", str, err)
		}
		return any(b).(T), nil

	case float64:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return zero, fmt.Errorf("error converting '%s' to float64: %w", str, err)
		}
		return any(f).(T), nil

	default:
		return zero, fmt.Errorf("unsupported type conversion for %T", zero)
	}
}
