package main

import (
	"log"
	"net/http"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-http-server"
	app.Usage = "Start web server and public your current directory."

	app.Version = "0.0.2"
	app.Authors = append(app.Authors, cli.Author{Name: "Kohei Kawasaki", Email: "mynameiskawasaq@gmail.com"})

	app.Commands = []cli.Command{
		{
			Name: "start",
			Aliases: []string{"s"},
			Usage: "Start the web server",
			Action: start,
			Flags: []cli.Flag {
				cli.StringFlag {
					Name: "port, p",
					Value: "8000",
					Usage: "Port number to listen to",
					EnvVar: "PORT",
				},
			},
		},
	}

	app.Action = start

	app.Run(os.Args)
}

func start(c *cli.Context) {
	port := c.String("port")
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir("."))))
}
