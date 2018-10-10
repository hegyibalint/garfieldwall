package main

import (
	"flag"
	"github.com/hegyibalint/garfieldwall/comic"
	"log"
	"math"
	"os"
)

func main() {
	//var help bool
	var width, height int
	var scale float64
	var help bool

	width = 1920
	height = 1080
	scale = 4.0

	flag.IntVar(&width, "width", math.MaxInt64, "The width of the comic wall")
	flag.IntVar(&height, "height", math.MaxInt64, "The height of the comic wall")
	flag.Float64Var(&scale, "scale", 3.0, "Downscale factor of comics")
	flag.BoolVar(&help, "help", false, "Showing the help")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	errorLogger := log.New(os.Stderr, "", 0)
	if width == math.MaxInt64 {
		errorLogger.Println("Parameter width have to defined")
		flag.Usage()
		os.Exit(-10)
	} else if width <= 0 {
		errorLogger.Println("Width cannot be a nagative number")
		flag.Usage()
		os.Exit(-11)
	}

	if height == math.MaxInt64 {
		errorLogger.Println("Parameter height have to defined")
		flag.Usage()
		os.Exit(-20)
	} else if height <= 0 {
		errorLogger.Println("Height cannot be a nagative number")
		flag.Usage()
		os.Exit(-21)
	}

	if flag.NArg() == 0 {
		errorLogger.Println("The path must be given")
		flag.Usage()
		os.Exit(-30)
	}
	path := flag.Arg(0)

	timeRange := comic.DefaultRange()
	wall, err := comic.CreateWall(width, height, scale, timeRange)
	if err != nil {
		errorLogger.Println(err)
		os.Exit(-100)
	}

	err = comic.SaveWall(wall, path)
	if err != nil {
		errorLogger.Println(err)
		os.Exit(-101)
	}
}