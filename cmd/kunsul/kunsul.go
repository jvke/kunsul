package main

import (
	"k8s.io/client-go/rest"
	"os"

	"github.com/flaccid/kunsul"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	VERSION      = "v0.0.0-dev"
	templateFile = "/usr/share/kunsul/template.html"
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
		cli.IntFlag{
			Name:  "port,p",
			Usage: "port the listen on",
			Value: 8080,
		},
		// for for future use when a config file might be used
		cli.StringFlag{
			Name:  "config-dir,c",
			Usage: "configuration directory",
			Value: "/etc/kunsul",
		},
		cli.StringFlag{
			Name:  "template,t",
			Usage: "html template file",
			Value: (func() string {
				if _, err := os.Stat(templateFile); os.IsNotExist(err) {
					cwd, err := os.Getwd()
					if err != nil {
						log.Fatal(err)
					}
					return cwd + "/template.html"
				}
				return templateFile
			})(),
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
		err    error
	)
	if config, err = kunsul.GetConfig(); err != nil {
		cli.ShowAppHelp(c)
		return err
	}
	kunsul.Serve(config,
		c.String("config-dir"),
		c.String("template"),
		c.Int("port"),
		c.Bool("access-log"))

	return nil
}
