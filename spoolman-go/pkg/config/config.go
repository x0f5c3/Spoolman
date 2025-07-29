package config

import (
	"github.com/spf13/pflag"
	"strings"
	"sync"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
)

// LogConfig holds logging configuration.
type LogConfig struct {
	Level      string `koanf:"level"`
	LogFile    string `koanf:"log_file"`
	MaxSizeMB  int    `koanf:"max_size_mb"`
	MaxBackups int    `koanf:"max_backups"`
	MaxAgeDays int    `koanf:"max_age_days"`
}

// DBConfig holds database configuration.
type DBConfig struct {
	URL      string `koanf:"url"`
	PoolSize int    `koanf:"pool_size"`
	Type     string `koanf:"type"`
}

// AppConfig is the root configuration.
type AppConfig struct {
	ListenAddr string    `koanf:"listen_addr"`
	Log        LogConfig `koanf:"log"`
	DB         DBConfig  `koanf:"db"`
	JWTKeyPath string    `koanf:"jwt_key_path"`
}

var (
	k        = koanf.New(".")
	cfg      *AppConfig
	cfgMutex sync.RWMutex
	onlyOnce sync.Once
)

func LoadOnce(cfgFile string, fs *pflag.FlagSet) (err error) {

	onlyOnce.Do(func() {
		err = Load(cfgFile, fs)
	})
	return
}

// Load loads config from TOML, ENV, and CLI flags.
// - cfgFile: path to TOML config
// - fs: parsed *pflag.FlagSet (from cobra or stdlib flag)
func Load(cfgFile string, fs *pflag.FlagSet) error {
	// 1. File (TOML)
	if err := k.Load(file.Provider(cfgFile), toml.Parser()); err != nil {
		return err
	}
	// 2. ENV (prefix "APP_", case-insensitive, . separator)
	err := k.Load(env.Provider("APP_", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, "APP_")), "_", ".")
	}), nil)
	if err != nil {
		return err
	}
	// 3. CLI flags (use stdlib flag or cobra/pflag)
	if fs != nil {
		err = k.Load(posflag.Provider(fs, ".", k), nil)
		if err != nil {
			return err
		}
	}

	// Unmarshal to struct
	cfgMutex.Lock()
	defer cfgMutex.Unlock()
	return k.Unmarshal("", cfg)
}

// Get returns a copy of the current config struct.
func Get() *AppConfig {
	cfgMutex.RLock()
	defer cfgMutex.RUnlock()
	return cfg
}
