package config

import (
  "os"
  "io/ioutil"

  "airsonic-cli/utils"

  "gopkg.in/yaml.v2"
  "gopkg.in/urfave/cli.v1"
)

const APPNAME = "Airsonic CLI"
const APIFORMAT = "json"
const APIVERSION = "1.15.0"
const CONFIG_PATH = "config.yml"

type Config struct {
  APIVersion string
  APIFormat string
  AppName string
  Server string
  Username string
  Token string
  Salt string
  Verbose bool
}

type savedConfig struct {
  AppName string    `yaml:"appname"`
  Server string     `yaml:"server"`
  Username string   `yaml:"username"`
  Token string      `yaml:"token"`
  Salt string       `yaml:"salt"`
}

func LoadConfig(conf *Config, ctx *cli.Context)  {
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

func ReadConfig(conf *Config) error {
  file, _ := os.Open(CONFIG_PATH)
  decoder := yaml.NewDecoder(file)
  err := decoder.Decode(conf)
  if err != nil {
    return err
  }
  return nil
}

func WriteConfig(conf *Config) error {
  sconf := &savedConfig{
    AppName: GetAppName(conf),
    Server: GetServer(conf),
    Username: GetUsername(conf),
    Token: GetToken(conf),
    Salt: GetSalt(conf),
  }
  bytes, err := yaml.Marshal(&sconf)
  if err != nil {
    return err
  }
  return ioutil.WriteFile(CONFIG_PATH, bytes, 0640)
}

func GetAppName(conf *Config) string { return conf.AppName }
func GetAPIVersion(conf *Config) string { return conf.APIVersion }
func GetAPIFormat(conf *Config) string { return conf.APIFormat }
func GetServer(conf *Config) string { return conf.Server }
func GetUsername(conf *Config) string { return conf.Username }
func GetSalt(conf *Config) string { return conf.Salt }
func GetToken(conf *Config) string { return conf.Token }
func IsVerbose(conf *Config) bool { return conf.Verbose }

func SetAppName(conf *Config, data string) { conf.AppName = data }
func SetAPIVersion(conf *Config, data string) { conf.APIVersion = data }
func SetAPIFormat(conf *Config, data string) { conf.APIFormat = data }
func SetServer(conf *Config, data string) { conf.Server = data }
func SetUsername(conf *Config, data string) { conf.Username = data }
func SetSalt(conf *Config, data string) { conf.Salt = data }
func SetToken(conf *Config, data string) { conf.Token = data }
func SetVerbose(conf *Config, data bool) { conf.Verbose = data }
