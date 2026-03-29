package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Path   string `mapstructure:"path"`
		Driver string `mapstructure:"driver"`
	} `mapstructure:"database"`
	Server struct {
		Port       int      `mapstructure:"port"`
		Cors       struct {
			AllowedOrigins []string `mapstructure:"allowedOrigins"`
		} `mapstructure:"cors"`
	} `mapstructure:"server"`
	Crawler struct {
		Interval time.Duration `mapstructure:"interval"`
		Sources  []struct {
			Name string `mapstructure:"name"`
			URL  string `mapstructure:"url"`
		} `mapstructure:"sources"`
	} `mapstructure:"crawler"`
	OpenAI struct {
		APIKey string `mapstructure:"api_key"`
		Model  string `mapstructure:"model"`
	} `mapstructure:"openai"`
	Logging struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"logging"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// čŽžç˝ŽçŻĺ˘ĺé
	viper.SetEnvPrefix("OPFLOW")
	viper.AutomaticEnv()

	// ćŁćĽĺżčŚéç˝?
	if config.OpenAI.APIKey == "" {
		config.OpenAI.APIKey = os.Getenv("OPENAI_API_KEY")
		if config.OpenAI.APIKey == "" {
			return nil, fmt.Errorf("OpenAI API key is required")
		}
	}

	return &config, nil
}
