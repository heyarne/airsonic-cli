package config

import (
	"io/ioutil"
	"os"

	"airsonic-cli/utils"

	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v2"
)

// APPNAME define the default application name
const APPNAME = "Airsonic CLI"

// APIFORMAT define the default API format to use (should not be changed)
const APIFORMAT = "json"

// APIVERSION define the default API version to use (should not be changed)
const APIVERSION = "1.15.0"

// CONFIGPATH define the default path where to find the configuration file
const CONFIGPATH = "config.yml"

// Config is the application configuration structure that can be easily passed trough functions
type Config struct {
	APIVersion string
	APIFormat  string
	AppName    string
	Server     string
	Username   string
	Token      string
	Salt       string
	Verbose    bool
}

type savedConfig struct {
	AppName  string `yaml:"appname"`
	Server   string `yaml:"server"`
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
	Salt     string `yaml:"salt"`
}

// LoadConfig allows you to load your configuration from cli.Context and config.Config
func LoadConfig(conf *Config, ctx *cli.Context) {
	SetAPIVersion(conf, APIVERSION)
	SetAPIFormat(conf, APIFORMAT)

	if ctx.String("appname") != "" {
		SetAppName(conf, ctx.String("appname"))
	} else {
		if GetAppName(conf) == "" {
			SetAppName(conf, APPNAME)
		}
	}

	if ctx.String("server") != "" {
		SetServer(conf, ctx.String("server"))
	} else {
		if GetServer(conf) == "" {
			SetServer(conf, utils.Prompt("server"))
		}
	}

	if ctx.String("username") != "" {
		SetUsername(conf, ctx.String("username"))
	} else {
		if GetUsername(conf) == "" {
			SetUsername(conf, utils.Prompt("username"))
		}
	}

	if ctx.String("salt") != "" && ctx.String("token") != "" {
		SetSalt(conf, ctx.String("salt"))
		SetToken(conf, ctx.String("token"))
	} else {
		if ctx.String("password") != "" {
			SetSalt(conf, utils.GenerateSalt())
			SetToken(conf, utils.GenerateToken(ctx.String("password"), GetSalt(conf)))
		} else {
			if GetSalt(conf) == "" && GetToken(conf) == "" {
				SetSalt(conf, utils.GenerateSalt())
				SetToken(conf, utils.GenerateToken(utils.Prompt("password"), GetSalt(conf)))
			}
		}
	}
	SetVerbose(conf, ctx.Bool("verbose"))
}

// ReadConfig allows you to read the configuration from a configuration file
func ReadConfig(conf *Config) error {
	file, _ := os.Open(CONFIGPATH)
	decoder := yaml.NewDecoder(file)
	err := decoder.Decode(conf)
	if err != nil {
		return err
	}
	return nil
}

// WriteConfig allows you to write the configuration to a configuration file
func WriteConfig(conf *Config) error {
	sconf := &savedConfig{
		AppName:  GetAppName(conf),
		Server:   GetServer(conf),
		Username: GetUsername(conf),
		Token:    GetToken(conf),
		Salt:     GetSalt(conf),
	}
	bytes, err := yaml.Marshal(&sconf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(CONFIGPATH, bytes, 0640)
}

// GetAppName returns the application name
func GetAppName(conf *Config) string { return conf.AppName }

// GetAPIVersion returns the API version
func GetAPIVersion(conf *Config) string { return conf.APIVersion }

// GetAPIFormat returns the API format
func GetAPIFormat(conf *Config) string { return conf.APIFormat }

// GetServer returns the server URL
func GetServer(conf *Config) string { return conf.Server }

// GetUsername returns the user name
func GetUsername(conf *Config) string { return conf.Username }

// GetSalt returns the user generated salt
func GetSalt(conf *Config) string { return conf.Salt }

// GetToken returns the user generated token
func GetToken(conf *Config) string { return conf.Token }

// IsVerbose returns wether you enabled verbose mode
func IsVerbose(conf *Config) bool { return conf.Verbose }

// SetAppName allows you to write the application name to the configuration
func SetAppName(conf *Config, data string) { conf.AppName = data }

// SetAPIVersion allows you to write the API version to the configuration
func SetAPIVersion(conf *Config, data string) { conf.APIVersion = data }

// SetAPIFormat allows you to write the API format to the configuration
func SetAPIFormat(conf *Config, data string) { conf.APIFormat = data }

// SetServer allows you to write the server URL to the configuration
func SetServer(conf *Config, data string) { conf.Server = data }

// SetUsername allows you to write the user name to the configuration
func SetUsername(conf *Config, data string) { conf.Username = data }

// SetSalt allows you to write the user generated salt to the configuration
func SetSalt(conf *Config, data string) { conf.Salt = data }

// SetToken allows you to write the user generated token to the configuration
func SetToken(conf *Config, data string) { conf.Token = data }

// SetVerbose allows you to write the verbose state to the configuration
func SetVerbose(conf *Config, data bool) { conf.Verbose = data }
