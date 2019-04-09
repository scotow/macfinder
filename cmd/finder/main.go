package main

import (
	"flag"
	"fmt"
	"github.com/scotow/macfinder"
	"log"
	"os"
	"strconv"
)

var (
	storageFlag 	= flag.String("s", "", "invalid disk capacity")
	modelFlag 	 	= flag.String("m", "", "model string")
	yearFlag	 	= flag.Int("y", 0, "release year")
	colorFlag    	= flag.String("c", "", "product color")
	ramFlag      	= flag.Int("r", 0, "ram size")
	screenFlag 		= flag.Int("d", 0, "screen dimension")
)

func main() {
	flag.Parse()

	if *storageFlag == "" || *modelFlag == "" || *yearFlag == 0 || *colorFlag == "" || *ramFlag == 0 || *screenFlag == 0 {
		flag.Usage()
		os.Exit(1)
	}

	model, err := macfinder.FindModel(macfinder.Specs{
		Capacity: *storageFlag,
		Name:     *modelFlag,
		Year:     strconv.Itoa(*yearFlag),
		Color:    *colorFlag,
		Ram:      strconv.Itoa(*ramFlag) + "gb",
		Screen:   strconv.Itoa(*screenFlag) + "inch",
	})
	if err != nil {
		log.Fatalln(model)
	}

	if model == nil {
		fmt.Println("Model not available.")
		os.Exit(1)
	}

	fmt.Printf("Model available: %s\n", "https://apple.com" + model.Link)
}