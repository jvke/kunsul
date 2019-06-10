package main

import (
	"k8s.io/client-go/rest"
	"os"

	"github.com/flaccid/kunsul"
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
)

var (
	VERSION = "v0.0.0-dev"
)

func beforeApp(c *cli.Context) error {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "kunsul"
	app.Version = VERSION
	app.Usage = "kubernetes global ingress ui dashboard"
	app.Action = start
	app.Before = beforeApp
	app.Flags = []cli.Flag{
		// TODO: implement static file support for overlay
		// cli.StringFlag{
		// 	Name:  "directory,d",
		// 	Usage: "directory of static files to host",
		// 	Value: "./",
		// },
		cli.BoolFlag{
			Name:  "listings,l",
			Usage: "enable directory listings",
		},
		cli.IntFlag{
			Name:  "port,p",
			Usage: "port the listen on",
			Value: 8080,
		},
		cli.StringFlag{
			Name:  "config-dir,c",
			Usage: "configuration directory",
			Value: (func() string {
				if rt, err := os.Getwd() ; err == nil {
					return rt
				}
				return ""
			} )(),
		},
		cli.StringFlag{
			Name:  "template,t",
			Usage: "template file",
			Value: "template.html",
		},
		// todo: KUBERNETES_SERVICE_HOST and KUBERNETES_SERVICE_PORT
		// todo: always on, need to implement
		cli.BoolFlag{
			Name:  "access-log,a",
			Usage: "enable access logging of requests",
		},
		cli.BoolFlag{
			Name:  "debug,D",
			Usage: "run in debug mode",
		},
	}
	app.Run(os.Args)

}

func start(c *cli.Context) error {
	var (
		config *rest.Config
		err error
	)
	if config, err = kunsul.GetConfig(); err != nil {
		cli.ShowAppHelp(c)
		return err
	}
	kunsul.Serve(config,
		c.String("config-dir"),
		c.String("template"),
		c.Int("port"),
		c.Bool("listings"),
		c.Bool("access-log"))

	return nil
}
