package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Description = "Encode/decode quadkeys to/from lat,long"
	app.Name = "quadkey"
	app.Usage = ""
	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:  "encode",
			Usage: "encode lat,long to quadkey",
			Action: func(c *cli.Context) error {
				return encode(c)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "coords, c",
					Usage: "comma-separated latitude and longitude, eg 51.051509,3.739270",
				},
				cli.IntFlag{
					Name:  "level, l",
					Usage: "Level of quadkey to encode, eg 6",
					Value: 12,
				},
			},
		},
		// {
		// 	Name:  "decode",
		// 	Usage: "convert quadkey to lat,long",
		// 	Action: func(c *cli.Context) error {
		// 		return executeCommand(c, importShow)
		// 	},
		// },
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
