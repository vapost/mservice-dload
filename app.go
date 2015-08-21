package main

import (
	"github.com/vapost/mservice-dload/service"
	"errors"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

func getConfig(c *cli.Context) (service.Config, error) {

	yamlPath := c.GlobalString("config")

	config := service.Config{}

	if _, err := os.Stat(yamlPath); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)
	return config, err
}

type Bla struct {
	yo string
}

func (b *Bla)goNuts() string {
	return "Going crazey"
}

func main() {

	app := cli.NewApp()
	app.Name = "TTDL"
	app.Usage = "Traveltainment (TT) data download microservice"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			"config, c",
			"config.yaml",
			"config file to use",
			"APP_CONFIG",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the http server",
			Action: func(c *cli.Context) {
				cfg, err := getConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

				svc := service.StatService{}

				bla := Bla{}
				fmt.Println(bla.goNuts(), svc.Run(cfg))

				/*if err = svc.Run(cfg); err != nil {
					log.Fatal(err)
				}*/
			},
		},

		
	}

	app.Run(os.Args)
}