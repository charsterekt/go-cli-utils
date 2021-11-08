package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/urfave/cli"
)

func main() {

	// cli client
	app := cli.NewApp()
	app.Name = "File Downloader"
	app.Usage = "Download a file from a given link"

	// flags
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "link, l",
			Usage: "Link to the file",
		},
	}

	// app commands
	app.Commands = []cli.Command{

		{
			Name:  "dl",
			Usage: "Downloads a file present at the given link",
			Flags: myFlags,
			Action: func(c *cli.Context) {
				download(c.String("link"))
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func download(link string) {

	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", link)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
}
