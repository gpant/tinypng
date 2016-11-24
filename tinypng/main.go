/*

A TinyPNG client written in Go

Compress a PNG or JPG file with the help of the TinyPNG service.

Usage

Then you can run the command:

    tinypng <input.png> [output.png]

*/
package main

import (
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/gpant/tinypng"
	"github.com/mitchellh/go-homedir"
)

func main() {
	inputFilename, outputFilename := getFilenames(os.Args)

	// First check if the input file actually exist
	if !fileExists(inputFilename) {
		fatalError("Input file does not exist.")
	}

	// Verify that the input file is a PNG or JPEG file
	if !validFileType(inputFilename) {
		fatalError("Input file is not a valid PNG or JPEG file.")
	}

	// Then make sure that the output file doesn’t exist
	if fileExists(outputFilename) {
		fatalError("Output file already exist.")
	}

	// Get the API key
	apiKey, err := getAPIKeyFromEnv()
	check(err)

	res, err := tinypng.ShrinkFn(apiKey, inputFilename)

	if err != nil {
		fatalRed(res.Error+":", res.Message)
	}

	// Check if we are in debug mode
	verbose := os.Getenv("TINYPNG_VERBOSE")

	// Print if TINYPNG_VERBOSE is true
	if verbose == "true" {
		res.Print()
	}

	// Download the compressed PNG
	res.SaveAs(outputFilename)
}

//Configuration

// Info from config file
type Config struct {
	TINYPNG_API_KEY string
}

// Reads info from config file
func ReadConfig(configfile string) Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}

// Handle input

func getFilenames(args []string) (string, string) {
	// Make sure that we got one or two command line arguments
	if len(args) < 2 || len(args) > 3 {
		getAPIKeyFromEnv()
		fmt.Printf(green("TinyPNG") + " v0.0.4\n")
		fatalGreen("Usage:", "tinypng <input.png/jpg> [output.png/jpg]")
	}

	if len(args) == 2 {
		dir, file := path.Split(args[1])

		return args[1], path.Join(dir, "tiny-"+file)
	}

	return args[1], args[2]
}

func getAPIKeyFromEnv() (string, error) {
	var apiKey string
	var configPath string

	path, err := homedir.Dir()
	if err == nil {
		configPath = path + "/.tinypng/tinypng.config"
	}

	if _, err := os.Stat(configPath); err == nil {
		var config = ReadConfig(configPath)
		apiKey = config.TINYPNG_API_KEY
	} else {
		apiKey = os.Getenv("TINYPNG_API_KEY")

		if apiKey == "" {
			message := "No API key found. Please either create a config file or set an ENV variable " + green("TINYPNG_API_KEY") + " for it"
			return "", errors.New(message)
		}
	}
	return apiKey, nil
}

// IO

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Valid file type

func validFileType(fn string) bool {
	return validPNGFile(fn) || validJPEGFile(fn)
}

// PNG

func validPNGFile(fn string) bool {
	pngImage, err := os.Open(fn)

	check(err)

	defer pngImage.Close()

	_, err = png.DecodeConfig(pngImage)

	if err != nil {
		return false
	}

	return true
}

// JPEG

func validJPEGFile(fn string) bool {
	jpegImage, err := os.Open(fn)

	check(err)

	defer jpegImage.Close()

	_, err = jpeg.DecodeConfig(jpegImage)

	if err != nil {
		return false
	}

	return true
}

// Fatal

func check(err error) {
	if err != nil {
		fatal(red("Error:"), err)
	}
}

func fatal(v ...interface{}) {
	fmt.Println(v...)

	os.Exit(1)
}

func fatalError(message string) {
	fatalRed("Error:", message)
}

func fatalRed(title, message string) {
	fatal(red(title), message)
}

func fatalGreen(title, message string) {
	fatal(green(title), message)
}

// Color

func color(c, s string) string {
	// No ANSI escape sequences for Windows
	if runtime.GOOS == "windows" {
		return s
	}

	return "\033[" + c + "m" + s + "\033[0m"
}

func red(s string) string {
	return color("0;31", s)
}

func green(s string) string {
	return color("0;32", s)
}
