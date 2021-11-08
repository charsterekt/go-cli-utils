package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Host Info CLI"
	app.Usage = "Lets you look up a site's host details"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "The host to lookup",
		},
	}

	app.Commands = []cli.Command{

		// NAME SERVER LOOKUP

		{
			Name:  "ns",
			Usage: "Looks up the NameServers for the given host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}

				fmt.Println("Looking up Nameservers...")
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},

		// IP ADDRESS LOOKUP

		{
			Name:  "ip",
			Usage: "Looks up the IP Addresses for the given host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}

				fmt.Println("Looking up IP addresses...")
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},

		// CNAME RECORD LOOKUP

		{
			Name:  "cname",
			Usage: "Looks up the CName for the given host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println("Looking up CNAME records...")
				fmt.Println(cname)
				return nil
			},
		},

		// MX RECORD LOOKUP

		{
			Name:  "mx",
			Usage: "Looks up the MX Records for the given host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}

				fmt.Println("Looking up MX records...")
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
