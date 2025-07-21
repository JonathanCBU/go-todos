package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"sync"
)

type CsvConfig struct {
	TimestampFormat string   `json:"timestamp_format"`
	Headers         []string `json:"headers"`
	Statuses        []string `json:"statuses"`
}

type Config struct {
	CSV CsvConfig `json:"csv"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &config, nil
}

var (
	instance *Config
	once     sync.Once
)

func Load(filename string) error {
	var err error
	once.Do(func() {
		data, readErr := os.ReadFile(filename)
		if readErr != nil {
			err = fmt.Errorf("failed to read config file: %w", readErr)
			return
		}

		var config Config
		parseErr := json.Unmarshal(data, &config)
		if parseErr != nil {
			err = fmt.Errorf("failed to parse JSON: %w", parseErr)
			return
		}

		instance = &config
	})
	return err
}

func Get() *Config {
	if instance == nil {
		log.Fatal("Config not loaded. Call config.Load() first")
	}
	return instance
}

func GetTimestampFormat() string {
	return Get().CSV.TimestampFormat
}

func GetHeaders() []string {
	return Get().CSV.Headers
}

func ValidateStatus(status string) error {
	if !slices.Contains(instance.CSV.Statuses, status) {
		return fmt.Errorf("%s is not a valid status", status)
	}
	return nil
}

func ValidateHeader(header string) error {
	if !slices.Contains(instance.CSV.Headers, header) {
		return fmt.Errorf("%s is not a valid header", header)
	}
	return nil
}
