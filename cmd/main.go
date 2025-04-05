package main

import (
	"flag"
	"log"
	"microUrl/pkg/url"
	"strconv"
)

const DefaultLength string = "5"
const UsePaddingFlag = "p"
const UseLengthFlag = "l"

func main() {
	length := flag.String(UseLengthFlag, DefaultLength, "Length of the shortened URL")
	padding := flag.Bool(UsePaddingFlag, false, "Use padding")
	flag.Parse()
	args := flag.Args()

	shortUrl, err := url.Shrink(parseArgs(args, length, padding))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Params:", args)
	log.Printf("Shortened URL: %s Length: %d", shortUrl, len(shortUrl))
}

func parseArgs(args []string, length *string, padding *bool) (string, uint, bool) {
	var l uint
	if len(args) < 1 {
		log.Fatal("url is required")
	}

	le, err := strconv.Atoi(*length)
	if err != nil {
		log.Fatal("Unrecognized length")
	}
	l = uint(le)

	return args[0], l, *padding

}
