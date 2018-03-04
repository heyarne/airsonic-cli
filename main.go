package main

import (
  "log"
  "os"

  "airsonic-cli/config"
  "airsonic-cli/api/ping"
  "airsonic-cli/api/scan"

  "gopkg.in/urfave/cli.v1"
)

func main() {
  conf := &config.Config{}
  app := cli.NewApp()
  app.Name = "airsonic-cli"
  app.Version = "0.0.1b1"
  app.Usage = "management tool for Subsonic/Airsonic API"
  app.UsageText = "Usage needs to be written 0_o"
  cli.VersionFlag = cli.BoolFlag{
    Name: "version, V",
    Usage: "Print the version",
  }
  cli.HelpFlag = cli.BoolFlag{
    Name: "help, h",
    Usage: "Show help",
  }
  app.Flags = []cli.Flag{
    cli.BoolFlag{
      Name:  "verbose, v",
      Usage: "Enable verbose mode",
    },
    cli.StringFlag{
      Name:  "config, c",
      Value: config.CONFIG_PATH,
      Usage: "Load configuration from `<file>`",
    },
    cli.StringFlag{
      Name:  "server, s",
      Usage: "Specify server location `<server>[:<port>]`",
    },
    cli.StringFlag{
      Name:  "username, u",
      Usage: "Specify user name `<username>`",
    },
    cli.StringFlag{
      Name:  "password, p",
      Usage: "Specify user password `<password>`",
    },
    cli.StringFlag{
      Name:  "token, T",
      Usage: "Specify user generated token `<token>`",
    },
    cli.StringFlag{
      Name:  "salt, S",
      Usage: "Specify user generated salt `<salt>`",
    },
    cli.StringFlag{
      Name:  "appname, a",
      Usage: "Specify application name`<appname>`",
    },
  }
  app.Before = func(ctx *cli.Context) error {
    if ctx.String("config") != "" {
      config.ReadConfig(conf)
      config.LoadConfig(conf, ctx)
    }
    return nil
  }
  app.After = func(ctx *cli.Context) error {
    if ctx.String("config") != "" {
      config.WriteConfig(conf)
    }
    return nil
  }
  app.Commands = []cli.Command{
    {
      Name:    "ping",
      Usage:   "Ping your server",
      Action:  func(ctx *cli.Context) error  {
        ping.PingAction(conf)
        return nil
      },
    },
    {
      Name:    "scan",
      Usage:   "Start a rescan of your server's library",
      Action:  func(ctx *cli.Context) error  {
        scan.StartScanAction(conf)
        return nil
      },
      Subcommands: []cli.Command{
        {
          Name:  "status",
          Usage: "Get the scanning status",
          Action: func(c *cli.Context) error {
            scan.ScanStatusAction(conf)
            return nil
          },
        },
      },
    },
  }
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
